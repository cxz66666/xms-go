package news_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"xms/models"
	"xms/pkg/cache"
	"xms/pkg/response"
)

// GetLatestNews is used to get the latest newses ( at most 6)
func GetLatestNews(c *gin.Context) {
	g:=response.Gin{C: c}
	newsList:= cache.GetOrCreate(cache.GetNewsListKey(), func() interface{} {
		//take latest 6
		list,_:=models.GetNewsPaginated(1,6)
		return list
	}).([]models.News)
	var newsPiece []NewsPiece
	for _,item:=range newsList {
		newsPiece = append(newsPiece, NewsPiece{
			Id: item.ID,
			Content: item.Content,
			Title: item.Title,
			UpdateAt: item.CreatedTime,
		})
	}
	g.Success(http.StatusOK,NewsRes{
		PageId: 1,
		Count: len(newsPiece),
		Content: newsPiece,
	})
	return
}

// GetNews is used to get special news by id in param
func GetNews(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}

	news:=cache.GetOrCreate(cache.GetKey(cache.News,id), func() interface{} {
		news,err:=models.GetNewsById(id)
		if err!=nil{
			return models.News{}
		}
		return news
	}).(models.News)

	if news.ID<=0 {
		g.Error(http.StatusOK,response.ERROR_NEWS_NOT_FOUND,nil)
		return
	}
	g.Success(http.StatusOK,NewsPiece{
		Id: news.ID,
		Title: news.Title,
		Content: news.Content,
		UpdateAt: news.CreatedTime,
	})
	return
}

func PostNeNews(c *gin.Context)  {
	g:=response.Gin{C: c}
	var data NewNewsReq
	if err:=c.ShouldBind(&data);err!=nil{
		g.Error(http.StatusOK,response.ERROR_NEWS_INVALID_INFO,err)
		return
	}

	news:=models.News{
		Title: data.Title,
		Content: data.Content,
		CreatedTime: time.Now(),
	}

	cache.Remove(cache.GetNewsListKey())

	if id,err:=models.AddNewNews(news);err!=nil{
		g.Error(http.StatusOK,response.ERROR_DATABASE_QUERY,err)
	} else {
		g.Success(http.StatusOK,gin.H{
			"id":id,
		})
	}
}

func UpdateNews(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	var data NewNewsReq
	if err:=c.ShouldBind(&data);err!=nil{
		g.Error(http.StatusOK,response.ERROR_NEWS_INVALID_INFO,err)
		return
	}

	cache.Remove(cache.GetKey(cache.News,id))

	oldNews,err:=models.GetNewsById(id)
	if err!=nil{
		g.Error(http.StatusOK,response.ERROR_NEWS_NOT_FOUND,err)
		return
	}
	if len(data.Title)>0 {
		oldNews.Title=data.Title
	}
	if len(data.Content)>0{
		oldNews.Content=data.Content
	}

	_=models.UpdateNews(oldNews)

	cache.Remove(cache.GetNewsListKey())

	g.Success(http.StatusOK,nil)
	return
}

func DeleteNews(c *gin.Context)  {
	g:=response.Gin{C: c}
	idStr:=c.Param("id")
	id,ok:=strconv.Atoi(idStr)
	if !(ok==nil&&id>0){
		g.Error(http.StatusOK,response.ERROR_PARAM_NOT_VAILD,nil)
		return
	}
	_,err:=models.GetNewsById(id)
	if err!=nil	{
		g.Error(http.StatusOK,response.ERROR_NEWS_NOT_FOUND,err)
		return
	}

	err=models.DeleteNews(id)
	if err!=nil{
		g.Error(http.StatusOK,response.ERROR_DATABASE_QUERY,err)
		return
	}

	cache.Remove(cache.GetNewsListKey())
	cache.Remove(cache.GetKey(cache.News,id))

	g.Success(http.StatusOK,nil)
}

func GetNewsList(c *gin.Context)  {
	g:=response.Gin{C: c}
	pageSizeStr:=c.Query("size")
	pageIdStr:=c.Query("pageId")


	var  pageSize int
	var  pageId int
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
	if len(pageSizeStr)==0 {
		pageSize=30
	} else {
		num,ok:=strconv.Atoi(pageSizeStr)
		if ok!=nil||num<0||num>30 {
			invalid=true
		} else {
			pageSize=num
		}
	}
	if invalid{
		g.Error(http.StatusOK,response.ERROR_NEWS_INVALID_QUERY,nil)
		return
	}

	list,pageCount:=models.GetNewsPaginated(pageId,pageSize)

	res:=NewsListRes{
		PageCount: pageCount,
		Size: pageSize,
		PageId: pageId,
		Content: make([]NewsPiece,0,len(list)),
	}
	for _,item:=range list{
		res.Content = append(res.Content, NewsPiece{
			Id: item.ID,
			Title: item.Title,
			Content: item.Content,
			UpdateAt: item.CreatedTime,
		})
	}
	g.Success(http.StatusOK,res)
	return

}