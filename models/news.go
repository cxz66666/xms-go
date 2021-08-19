package models

import "time"

type News struct {
	ID int `gorm:"primaryKey" json:"id" form:"id"`

	//标题
	Title string `json:"title" form:"title" binding:"required,max=50"`

	//内容
	Content string `json:"content" form:"content" binding:"required,max=1000"`

	//创建的时间
	CreatedTime time.Time  `json:"created_time"`
}

func (News)TableName() string {
	return "newses"
}


//SetCreatedTime set news.CreatedTime to time.Now
func (news *News) SetCreatedTime() {
	news.CreatedTime=time.Now()
}
