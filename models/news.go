package models

import "time"

type News struct {
	ID int `gorm:"primaryKey;column:Id" json:"id" form:"id"`

	//标题
	Title string `gorm:"column:Title;size:50" json:"title" form:"title" binding:"required,max=50"`

	//内容
	Content string `gorm:"column:Content;size:1000" json:"content" form:"content" binding:"required,max=1000"`

	//创建的时间
	CreatedTime time.Time  `gorm:"column:CreatedTime" json:"created_time"`
}

func (News)TableName() string {
	return "newses"
}


//SetCreatedTime set news.CreatedTime to time.Now
func (news *News) SetCreatedTime() {
	news.CreatedTime=time.Now()
}


//GetNewsPaginated get the news according to pageId and pageSize, but important pageId range from 1 to inf, not 0, so don't forget to minus 1
// second return is pageCount
func GetNewsPaginated(pageId int, pageSize int) ([]News,int) {
	var news []News
	var count int64
	pageId-=1
	db.Model(&news).Count(&count)

	db.Order("CreatedTime desc").Offset(pageSize*pageId).Limit(pageSize).Find(&news)

	pageCount:=int(count)/pageSize
	if int(count)%pageSize!=0 {
		pageCount++
	}

	return news,pageCount
}

// GetNewsById get the news from database by id
func GetNewsById(id int) (News,error) {
	var news News
	if err:=db.First(&news,id).Error;err!=nil{
		return news,err
	}
	return news,nil
}


// AddNewNews will add the bill to database, and return id if success
func AddNewNews(news News) (int,error)  {
	if err:=db.Create(&news).Error;err!=nil{
		return 0,err
	}
	return news.ID,nil
}


//UpdateNews will update all the column of bill
func UpdateNews(news News) error  {
	if err:=db.Select("*").Updates(&news).Error;err!=nil{
		return err
	}
	return nil
}

func DeleteNews(id int) error {
	if err:=db.Delete(&News{ID: id}).Error;err!=nil{
		return err
	}
	return nil
}