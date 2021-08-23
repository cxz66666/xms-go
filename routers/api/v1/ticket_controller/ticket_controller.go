package ticket_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
	"xms/models"
	"xms/pkg/cache"
	"xms/pkg/response"
)

func SearchTicketList(c *gin.Context)  {
	g:=response.Gin{C: c}
	countStr:=c.Query("count")
	pageIdStr:=c.Query("pageId")
	valuesStr:=c.Query("values")
	queryTypeStr:=c.Query("queryType")


	var  count int
	var  pageId int
	var  queryType int

	invalid:=false


	// check pageId
	if len(pageIdStr)==0 {
		pageId=1
	} else {
		num,ok:=strconv.Atoi(pageIdStr)
		if ok!=nil||num<=0 {
			invalid=true
		} else {
			pageId=num
		}
	}

	//check pageSize
	if len(countStr)==0 {
		count=30
	} else {
		num,ok:=strconv.Atoi(countStr)
		if ok!=nil||num<0||num>30 {
			invalid=true
		} else {
			count=num
		}
	}


	//check queryType
	if len(queryTypeStr)==0 {
		queryType=0
	} else {
		num,ok:=strconv.Atoi(queryTypeStr)
		if ok!=nil||num<0||num>2 {
			invalid=true
		} else {
			queryType=num
		}
	}

	if invalid{
		g.Error(http.StatusOK,response.ERROR_TICKET_SEARCH_INVALID_QUERY,nil)
		return
	}

	var res TicketListRes
	res.QueryType=queryType

	tickets,pageCount:=models.SearchTicket(valuesStr,pageId,count,queryType)
	res.PageCount=pageCount
	res.Size=len(tickets)
	res.Data=make([]TicketBriefInfo,0,len(tickets))
	for _,item:=range tickets{
		res.Data = append(res.Data, TicketBriefInfo{
			Id: item.ID,
			Type: item.Type,
			Status: item.Status,
			Device: item.Device,
			DeviceModel: item.DeviceModel,
			Owner: item.Owner,
			Phone: item.Phone,
			CreatedTime:item.CreatedTime,
			IsConfirmed: item.IsConfirmed,
		})
	}
	g.Success(http.StatusOK,res)
	return
}

func GetTicketList(c *gin.Context)  {
	g:=response.Gin{C: c}
	countStr:=c.Query("count")
	pageIdStr:=c.Query("pageId")
	queryTypeStr:=c.Query("queryType")


	var  queryCount int
	var  pageId int
	var  queryType int

	invalid:=false


	// check pageId
	if len(pageIdStr)==0 {
		pageId=1
	} else {
		num,ok:=strconv.Atoi(pageIdStr)
		if ok!=nil||num<=0 {
			invalid=true
		} else {
			pageId=num
		}
	}

	//check pageSize
	if len(countStr)==0 {
		queryCount =30
	} else {
		num,ok:=strconv.Atoi(countStr)
		if ok!=nil||num<0||num>30 {
			invalid=true
		} else {
			queryCount =num
		}
	}


	//check queryType
	if len(queryTypeStr)==0 {
		queryType=0
	} else {
		num,ok:=strconv.Atoi(queryTypeStr)
		if ok!=nil||num<0||num>2 {
			invalid=true
		} else {
			queryType=num
		}
	}
	if invalid{
		g.Error(http.StatusOK,response.ERROR_TICKET_LIST_INVALID_QUERY,nil)
		return
	}

	res:=TicketListRes{}
	res.QueryType=queryType
	res.PageIndex=pageId

	var tickets []models.Ticket
	var pageCount int
	if pageId==1&&queryType==0&&queryCount==30{
		ticketAndCount:=cache.GetOrCreate(cache.GetTicketFirstPageKey(), func() interface{} {
				t,count:=models.GetTicketsList(1,30,0)
				return models.TicketsAndCount{
					Tickets: t,
					PageCount: count,
				}
		}).(models.TicketsAndCount)

		pageCount=ticketAndCount.PageCount
		tickets=ticketAndCount.Tickets
	} else {
		tickets,pageCount=models.GetTicketsList(pageId,queryCount,queryType)
	}

	res.PageCount=pageCount
	res.Size=len(tickets)
	res.Data=make([]TicketBriefInfo,0,res.Size)

	for _,item:=range tickets {
		res.Data = append(res.Data, TicketBriefInfo{
			Id: item.ID,
			Type: item.Type,
			Status: item.Status,
			Device: item.Device,
			DeviceModel: item.DeviceModel,
			Owner: item.Owner,
			Phone: item.Phone,
			CreatedTime:item.CreatedTime,
			IsConfirmed: item.IsConfirmed,
		})
	}

	g.Success(http.StatusOK,res)
	return
}

func GetTicketInfo(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	item:=cache.GetOrCreate(cache.GetKey(cache.Ticket,id), func() interface{} {
		t,err:=models.GetTicketById(id)
		if err!=nil{
			return &models.Ticket{}
		}
		return t
	}).(*models.Ticket)
	if item.ID<=0 {
		g.Error(http.StatusOK,response.ERROR_TICKET_NOT_FOUND,nil)
		return
	}

	res:=TicketInfoRes{
		Id: item.ID,
		Type: item.Type,
		Status: item.Status,
		Device: item.Device,
		DeviceModel: item.DeviceModel,
		Owner: item.Owner,
		Phone: item.Phone,
		CreatedTime:item.CreatedTime,
		IsConfirmed: item.IsConfirmed,
		Description: item.Description,
		Accessories: make([]int,0),
		Workers: make([]string,0),

	}
	for _,r:=range item.Workers {
		if r.ID==id&&r.WorkerID==g.GetUserId() {
			res.Picked=true
			break
		}
	}

	for _,r:=range item.Workers {
		user:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,r.WorkerID), func() interface{} {
			user,err:= models.GetUserById(r.WorkerID)
			if err!=nil	{
				return &models.User{
					ID: -1,
				}
			}
			return user
		}).(*models.User)
		if user.ID>0 {
			res.Workers = append(res.Workers, user.Name)
		}
	}

	for _,r:=range item.Accessories {
		res.Accessories = append(res.Accessories, int(r.Type))
	}

	for _,note:=range item.Notes {
		user:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,note.OperatorId), func() interface{} {
			user,err:= models.GetUserById(note.OperatorId)
			if err!=nil	{
				return &models.User{
					ID: -1,
				}
			}
			return user
		}).(*models.User)
		if user.ID>0 {
			url:="https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
			if len(user.AvatarURL)>0{
				url=user.AvatarURL
			}
			res.Notes = append(res.Notes, NoteVM{
				Id: note.ID,
				Content: note.Content,
				CreatedTime: note.CreatedTime,
				Op: user.Name,
				Type: note.Type,
				Avatar: url,
			})
		}
	}
	g.Success(http.StatusOK,res)
}

func UpdateTicketInfo(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	var info TicketInfoReq
	if err:=c.ShouldBind(&info);err!=nil{
		g.Error(http.StatusOK,response.ERROR_TICKET_INVALID_INFO,err)
		return
	}

	ticketToUpdate,err:=models.GetTicketById(id)
	if err!=nil	{
		g.Error(http.StatusOK,response.ERROR_TICKET_NOT_FOUND,nil)
		return
	}
	if ticketToUpdate.IsConfirmed {
		g.Error(http.StatusOK,response.ERROR_TICKET_IS_CONFIRMED,nil)
	}
	var sb strings.Builder
	if len(info.Owner)>0&&info.Owner!=ticketToUpdate.Owner {
		sb.WriteString(fmt.Sprintf("[Owner Changed] %s => %s\n",ticketToUpdate.Owner,info.Owner))
		ticketToUpdate.Owner=info.Owner
	}


	if len(info.Device)>0&&info.Device!=ticketToUpdate.Device {
		sb.WriteString(fmt.Sprintf("[Device Changed] %s => %s\n",ticketToUpdate.Device,info.Device))
		ticketToUpdate.Device=info.Device
	}

	if len(info.DeviceModel)>0&&info.DeviceModel!=ticketToUpdate.DeviceModel {
		sb.WriteString(fmt.Sprintf("[DeviceModel Changed] %s => %s\n",ticketToUpdate.DeviceModel,info.DeviceModel))
		ticketToUpdate.DeviceModel=info.DeviceModel
	}

	if info.Type!=ticketToUpdate.Type {
		sb.WriteString(fmt.Sprintf("[Type Changed] %d => %d\n",ticketToUpdate.Type,info.Type))
		ticketToUpdate.Type=info.Type
	}

	if len(info.Phone)>0&&info.Phone!=ticketToUpdate.Phone {
		sb.WriteString(fmt.Sprintf("[Phone Changed] %s => %s\n",ticketToUpdate.Phone,info.Phone))
		ticketToUpdate.Phone=info.Phone
	}

	if len(info.Description)>0&&info.Description!=ticketToUpdate.Description {
		sb.WriteString(fmt.Sprintf("[Description Changed] %s => %s\n",ticketToUpdate.Description,info.Description))
		ticketToUpdate.Description=info.Description
	}
	isChangeAc:=false
	if len(ticketToUpdate.Accessories)!=len(info.Accessories) {
		isChangeAc=true
	} else {
		for _,i:=range ticketToUpdate.Accessories {
			found:=false
			for _,j:=range info.Accessories {
				if int(i.Type)==j{
					found=true
					break
				}
			}
			if found==false {
				isChangeAc=true
				break
			}
		}
	}
	if isChangeAc {
		sb.WriteString("[Accessories Changed]  something=> otherthing")
	}


	if len(sb.String())==0 {
		g.Success(http.StatusOK,nil)
		return
	}


	note:=models.Note{
		Type: models.Edit,
		Content: sb.String(),
		OperatorId: g.GetUserId(),
		CreatedTime: time.Now(),
	}
	ticketToUpdate.Notes = append(ticketToUpdate.Notes, note)

	if isChangeAc {
		ticketToUpdate.Accessories=make([]models.Accessory,0,len(info.Accessories))
		for _,ac:=range info.Accessories {
			ticketToUpdate.Accessories = append(ticketToUpdate.Accessories, models.Accessory{
				TicketID: ticketToUpdate.ID,
				Type: models.TicketAccessory(ac),
			})
		}
	}

	err=models.UpdateTicket(ticketToUpdate)
	if err!=nil{
		g.Error(http.StatusOK,response.ERROR_DATABASE_QUERY,err)
		return
	}

	cache.Remove(cache.GetKey(cache.Ticket,ticketToUpdate.ID))
	cache.Remove(cache.GetTicketFirstPageKey())

	g.Success(http.StatusOK,nil)
	return
}

func CreateNewTicket(c *gin.Context) {
	g:=response.Gin{C: c}
	var info TicketInfoReq
	if err:=c.ShouldBind(&info);err!=nil{
		g.Error(http.StatusOK,response.ERROR_TICKET_INVALID_INFO,err)
		return
	}
	ticket:=models.Ticket{
		Type: info.Type,
		Device: info.Device,
		DeviceModel: info.DeviceModel,
		Owner: info.Owner,
		Phone:info.Phone,
		Description: info.Description,
		CreatedTime: time.Now(),
		Status: models.Created,

	}
	note:=models.Note{
		OperatorId: g.GetUserId(),
		Type: models.Create,
		Content: "",
		CreatedTime: time.Now(),
	}
	ticket.Notes = append(ticket.Notes, note)

	id,_:=models.AddNewTicket(&ticket)


	cache.Remove(cache.GetTicketFirstPageKey())

	for _,ac:=range info.Accessories {
		ticket.Accessories = append(ticket.Accessories, models.Accessory{
			TicketID: id,
			Type: models.TicketAccessory(ac),
		})
	}

	_=models.UpdateTicket(&ticket)


	wechatConfig:= models.WechatConfig{
		TicketID: id,
		Phone: strings.Replace(ticket.Phone," ","",-1),
	}

	configId,_:= models.AddNewWechatConfig(&wechatConfig)
	ticket.WeChatConfigId=configId

	_=models.UpdateTicket(&ticket)

	cache.Set(cache.GetKey(cache.Ticket,id),&ticket,24*30*time.Hour)
	cache.Set(cache.GetKey(cache.Note,note.ID),&note,24*30*time.Hour)
	cache.Set(cache.GetKey(cache.WechatConfig,configId),&wechatConfig,24*30*time.Hour)

	g.Success(http.StatusOK,gin.H{
		"id":id,
	})
	return
}


func PickupTicket(c *gin.Context){
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	ticket:=cache.GetOrCreate(cache.GetKey(cache.Ticket,id), func() interface{} {
		t,err:=models.GetTicketById(id)
		if err!=nil{
			return &models.Ticket{}
		}
		return t
	}).(*models.Ticket)

	if ticket.ID<=0 {
		g.Error(http.StatusOK,response.ERROR_TICKET_NOT_FOUND,nil)
		return
	}
	if ticket.IsConfirmed {
		g.Error(http.StatusOK,response.ERROR_TICKET_IS_CONFIRMED,nil)
		return
	}
	uid:=g.GetUserId()
	for _,r:=range ticket.Workers {
		if r.TicketID==id&&r.WorkerID==uid{
			g.Error(http.StatusOK,response.ERROR_TICKET_ALREADY_PICKED,nil)
			return
		}
	}

	newWorker:=models.TicketWorker{
		TicketID: id,
		WorkerID: uid,
	}
	ticket.Workers = append(ticket.Workers, newWorker)


	note:=models.Note{
		OperatorId: uid,
		Type: models.Join,
		Content: "",
		CreatedTime: time.Now(),
	}
	ticket.Notes = append(ticket.Notes, note)

	err:=models.UpdateTicket(ticket)
	if err!=nil{
		g.Error(http.StatusOK,response.ERROR_DATABASE_QUERY,err)
		return
	}
	cache.Remove(cache.GetKey(cache.Ticket,ticket.ID))

	g.Success(http.StatusOK,nil)
	return

}

func LockTicket(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	ticket:=cache.GetOrCreate(cache.GetKey(cache.Ticket,id), func() interface{} {
		t,err:=models.GetTicketById(id)
		if err!=nil{
			return &models.Ticket{}
		}
		return t
	}).(*models.Ticket)

	if ticket.ID<=0{
		g.Error(http.StatusOK,response.ERROR_TICKET_NOT_FOUND,nil)
		return
	}
	if ticket.IsConfirmed {
		g.Error(http.StatusOK,response.ERROR_TICKET_IS_CONFIRMED,nil)
		return
	}
	if !models.CheckIfLockable(ticket.Status) {
		g.Error(http.StatusOK,response.ERROR_NOT_LOCKABLE,nil)
		return
	}

	ticket.IsConfirmed=true
	if ticket.Status==models.SuccessDone {
		for _,ticketworker:=range ticket.Workers{
			worker,err:=models.GetUserById(ticketworker.WorkerID)
			if err!=nil	{
				if ticket.Type==models.Appliance {
					worker.ApplianceFixedCount++
				} else {
					worker.ComputerFixedCount++
				}
				models.UpdateUser(worker)
				cache.Remove(cache.GetKey(cache.UserInfo,ticketworker.WorkerID))
			}
		}
	}
	cache.Remove(cache.GetKey(cache.Ticket,ticket.ID))
	cache.Remove(cache.GetTicketFirstPageKey())

	g.Success(http.StatusOK,nil)
	return
}


func UnlockTicket(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	ticket:=cache.GetOrCreate(cache.GetKey(cache.Ticket,id), func() interface{} {
		t,err:=models.GetTicketById(id)
		if err!=nil{
			return &models.Ticket{}
		}
		return t
	}).(*models.Ticket)
	if ticket.ID<=0{
		g.Error(http.StatusOK,response.ERROR_TICKET_NOT_FOUND,nil)
		return
	}
	if !ticket.IsConfirmed {
		g.Error(http.StatusOK,response.ERROR_TICKET_IS_CONFIRMED,nil)
		return
	}

	ticket.IsConfirmed=false
	if ticket.Status==models.SuccessDone {
		for _,ticketworker:=range ticket.Workers{
			worker,err:=models.GetUserById(ticketworker.WorkerID)
			if err!=nil	{
				if ticket.Type==models.Appliance {
					worker.ApplianceFixedCount--
				} else {
					worker.ComputerFixedCount--
				}
				models.UpdateUser(worker)
				cache.Remove(cache.GetKey(cache.UserInfo,ticketworker.WorkerID))
			}
		}
	}
	cache.Remove(cache.GetKey(cache.Ticket,ticket.ID))
	cache.Remove(cache.GetTicketFirstPageKey())
	g.Success(http.StatusOK,nil)
	return

}

func ChangeTicketStatus(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	ticket:=cache.GetOrCreate(cache.GetKey(cache.Ticket,id), func() interface{} {
		t,err:=models.GetTicketById(id)
		if err!=nil{
			return &models.Ticket{}
		}
		return t
	}).(*models.Ticket)
	if ticket.ID<=0{
		g.Error(http.StatusOK,response.ERROR_TICKET_NOT_FOUND,nil)
		return
	}

	var data ChangeStatusBodyReq

	if err:=c.ShouldBind(&data);err!=nil{
		g.Error(http.StatusOK,response.ERROR_TICKET_INVALID_INFO,err)
		return
	}
	if ticket.IsConfirmed {
		g.Error(http.StatusOK,response.ERROR_TICKET_IS_CONFIRMED,nil)
		return
	}

	ticket.Status=data.Status

	note:=models.Note{
		Type: models.ChangeState,
		Content: data.Status.ToDisplayName(),
		OperatorId: g.GetUserId(),
		CreatedTime: time.Now(),
	}
	ticket.Notes = append(ticket.Notes, note)

	err:=models.UpdateTicket(ticket)
	if err!=nil{
		g.Error(http.StatusOK,response.ERROR_DATABASE_QUERY,nil)
		return
	}
	cache.Remove(cache.GetKey(cache.Ticket,id))
	cache.Remove(cache.GetTicketFirstPageKey())

	g.Success(http.StatusOK,nil)
	return

}

func DeleteTicket(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	ticket:=cache.GetOrCreate(cache.GetKey(cache.Ticket,id), func() interface{} {
		t,err:=models.GetTicketById(id)
		if err!=nil{
			return &models.Ticket{}
		}
		return t
	}).(*models.Ticket)
	if ticket.ID<=0{
		g.Error(http.StatusOK,response.ERROR_TICKET_NOT_FOUND,nil)
		return
	}

	if ticket.IsConfirmed {
		g.Error(http.StatusOK,response.ERROR_TICKET_IS_CONFIRMED,nil)
		return
	}

	var notesCache []string
	for _,note:=range ticket.Notes{
		notesCache = append(notesCache, cache.GetKey(cache.Note,note.ID))
	}

	if ticket.WeChatConfigId!=0 {
		_=models.DeleteWechatConfigById(ticket.WeChatConfigId)
	}

	err:=models.DeleteTicketById(id)
	if err!=nil{
		g.Error(http.StatusOK,response.ERROR_DATABASE_QUERY,err)
		return
	}
	cache.Remove(cache.GetKey(cache.Ticket,ticket.ID))
	cache.Remove(cache.GetKey(cache.WechatConfig,ticket.WeChatConfigId))
	cache.Remove(cache.GetTicketFirstPageKey())
	cache.RemoveRange(notesCache)
	g.Success(http.StatusOK,nil)
	return
}

func PostNewComment(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	ticket:=cache.GetOrCreate(cache.GetKey(cache.Ticket,id), func() interface{} {
		t,err:=models.GetTicketById(id)
		if err!=nil{
			return &models.Ticket{}
		}
		return t
	}).(*models.Ticket)
	if ticket.ID<=0{
		g.Error(http.StatusOK,response.ERROR_TICKET_NOT_FOUND,nil)
		return
	}

	if ticket.IsConfirmed {
		g.Error(http.StatusOK,response.ERROR_TICKET_IS_CONFIRMED,nil)
		return
	}

	var data NewCommentReq
	if err:=c.ShouldBind(data);err!=nil{
		g.Error(http.StatusOK,response.ERROR_TICKET_INVALID_INFO,nil)
		return
	}

	ticket.Notes = append(ticket.Notes, models.Note{
		OperatorId: g.GetUserId(),
		Type: models.Comment,
		Content: data.Content,
		CreatedTime: time.Now(),
	})

	err:=models.UpdateTicket(ticket)
	if err!=nil{
		g.Error(http.StatusOK,response.ERROR_DATABASE_QUERY,err)
		return
	}
	cache.Remove(cache.GetKey(cache.Ticket,ticket.ID))

	g.Success(http.StatusOK,nil)
}


