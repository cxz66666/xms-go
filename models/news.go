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
