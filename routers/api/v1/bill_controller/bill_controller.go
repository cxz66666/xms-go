package bill_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"xms/models"
	"xms/pkg/cache"
	"xms/pkg/response"
	"xms/service/bill_service"
)

func CreateNewBill(c *gin.Context)  {
	g:=response.Gin{C: c}
	var data CreateBillReq
	if err:=c.ShouldBind(&data);err!=nil{
		g.Error(http.StatusOK,response.ERROR_BILL_INVALID_TYPE,err)
		return
	}
	bill:=models.Bill{
		Title: data.Title,
		Content: data.Content,
		Type: models.BillType(data.Type),
		UnitPrice: data.UnitPrice,
		Quantity: data.Quantity,
		Trader: data.Trader,
	}
	userId:=g.GetUserId()
	if userId<0 {
		g.Error(http.StatusOK,response.ERROR_NOT_LOGIN,nil)
		return
	}

	record:=models.BillTransactionRecord{
		OperatorId: userId,
		UpdateTime: time.Now(),
	}
	bill.Transactions = append(bill.Transactions, record)

	cache.Remove(cache.GetBillFirstPageKey())

	id,err:= bill_service.CreateBillAndUpdateCache(bill)
	if err!=nil{
		g.Error(http.StatusOK,response.ERROR_DATABASE_QUERY,nil)
		return
	}
	 g.Success(http.StatusOK,gin.H{
		"id":id,
	})
	return 
}

func GetBill(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	//create cache
	bill:=cache.GetOrCreate(cache.GetKey(cache.Bill,id), func() interface{} {
		bill,err:=models.GetBillById(id)
		if err!=nil{
			return models.Bill{}
		}
		return bill
	}).(models.Bill)
	if bill.ID<=0	{
		 g.Error(http.StatusOK,response.ERROR_BILL_NOT_FOUND,nil)
		return
	}
	vm:=NewBillViewModel(bill)
	for _,item:=range bill.Transactions{
		user:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,item.OperatorId), func() interface{} {
			user,_:= models.GetUserById(item.OperatorId)
			return user
		}).(*models.User)
		vm.Transactions = append(vm.Transactions, NewBillTransactionsViewModel(item,*user))
	}
	g.Success(http.StatusOK,vm)
	return
}

func DeleteBill(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	bill,err:=models.GetBillById(id)
	if err!=nil	{
		g.Error(http.StatusOK,response.ERROR_BILL_NOT_FOUND,err)
		return
	}
	err=models.DeleteBillById(id)
	if err!=nil	{
		g.Error(http.StatusOK,response.ERROR_DATABASE_QUERY,err)
		return
	}

	bill_service.UpdateBasicAmount(bill.Type==models.Income,-bill.UnitPrice*bill.Quantity)
	bill_service.UpdateBasicCount(bill.Type==models.Income,false)
	cache.Remove(cache.GetBillFirstPageKey())
	cache.Remove(cache.GetKey(cache.Bill,id))


	g.Success(http.StatusOK,nil)
	return
}

func UpdateBill(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	var data CreateBillReq
	if err:=c.ShouldBind(&data);err!=nil{
		g.Error(http.StatusOK,response.ERROR_BILL_INVALID_TYPE,err)
		return
	}
	bill,err:=models.GetBillById(id)
	if err!=nil{
		g.Error(http.StatusOK,response.ERROR_BILL_NOT_FOUND,err)
		return
	}

	cache.Remove(cache.GetBillFirstPageKey())

	bill_service.UpdateBasicAmount(bill.Type==models.Income,-bill.UnitPrice*bill.Quantity)
	bill_service.UpdateBasicCount(bill.Type==models.Income,false)


	bill.Title=data.Title
	bill.Content=data.Content
	bill.Type=models.BillType(data.Type)
	bill.UnitPrice=data.UnitPrice
	bill.Quantity=data.Quantity
	bill.Trader=data.Trader

	userId:=g.GetUserId()
	bill.Transactions = append(bill.Transactions, models.BillTransactionRecord{
		OperatorId: userId,
		UpdateTime: time.Now(),
	})

	//delete cache first
	cache.Remove(cache.GetBillFirstPageKey())
	//remove this cache
	cache.Remove(cache.GetKey(cache.Bill,id))

	//update bill and ignore the error
	_=models.UpdateBill(*bill)

	bill_service.UpdateBasicAmount(data.Type==int(models.Income),data.UnitPrice*data.Quantity)
	bill_service.UpdateBasicCount(data.Type==int(models.Income),true)

	g.Success(http.StatusOK,nil)

}

func GetBillList(c *gin.Context)  {
	g:=response.Gin{C: c}
	typeStr:=c.Query("type")
	pageSizeStr:=c.Query("pageSize")
	pageIdStr:=c.Query("pageId")
	invalid:=false

	var  pageSize int
	var  pageId int
	var billType int

	//check type   -1 0 1 or error
	if len(typeStr)==0 {
		billType=-1
	} else {
		num,ok:=strconv.Atoi(typeStr)
		if ok!=nil||(models.BillType(num)!=models.Income&&models.BillType(num)!=models.Outgoing){
			invalid=true
		} else {
			billType=num
		}
	}

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
	if len(pageSizeStr)==0 {
		pageSize=30
	} else {
		num,ok:=strconv.Atoi(pageSizeStr)
		if ok!=nil||num<0||num>50 {
			invalid=true
		} else {
			pageSize=num
		}
	}

	if invalid{
		g.Error(http.StatusOK,response.ERROR_BILL_INVALID_QUERY,nil)
		return
	}
	var bills []models.Bill
	//only first page can use cache
	//if pageId==1 &&pageSize==30&&billType==-1{
	//	bills=cache.GetOrCreate(cache.GetBillFirstPageKey(), func() interface{} {
	//		bills,_:=models.GetBillsPaginated(billType,pageId,pageSize)
	//		return bills
	//	}).([]models.Bill)
	//} else {
	//	bills,_=models.GetBillsPaginated(billType,pageId,pageSize)
	//}


	bills,pageCount:=models.GetBillsPaginated(billType,pageId,pageSize)

	billVms:=make([]BillViewModel,0,len(bills))

	for _,item:=range bills{
		vm:=NewBillViewModel(item)
		for _,t:=range item.Transactions {
			user:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,t.OperatorId), func() interface{} {
				user,_:= models.GetUserById(t.OperatorId)
				return user
			}).(*models.User)
			vm.Transactions = append(vm.Transactions, NewBillTransactionsViewModel(t,*user))
		}
		billVms = append(billVms, *vm)
	}

	g.Success(http.StatusOK, gin.H{
		"pageId":pageId,
		"pageSize":pageSize,
		"data":billVms,
		"pageCount":pageCount,
	})

	return
}

func GetStatsAmount(c *gin.Context)  {
	g:=response.Gin{C: c}
	cfg:=models.GetDbConfigValue(models.BasicAmountKey)
	balance, income, outgoing:=models.GetBasicAmountFromString(cfg)
	g.Success(http.StatusOK,gin.H{
		"balance":balance,
		"income":income,
		"outgoing":outgoing,
	})
	return
}



func GetStatsCount(c *gin.Context)  {
	g:=response.Gin{C: c}
	cfg:=models.GetDbConfigValue(models.BasicCountKey)
	income, outgoing:=models.GetBasicCountFromString(cfg)
	g.Success(http.StatusOK,gin.H{
		"income":income,
		"outgoing":outgoing,
	})
	return
}