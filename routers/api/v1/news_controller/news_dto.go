package news_controller

import "time"

type NewNewsReq struct {
	Title string `json:"title" form:"title" binding:"max=50"`
	Content string `json:"content" form:"content" binding:"max=1000"`
}

type NewsPiece struct {
	Id int 			`json:"id"`
	Title string 	`json:"title"`
	Content string `json:"content"`
	UpdateAt time.Time `json:"updateAt"`
}

type NewsRes struct {
	Count int  `json:"count"`
	PageId int  `json:"pageId"`
	Content []NewsPiece `json:"content"`
}

type NewsListRes struct {
	PageId int `json:"pageId"`
	Size int `json:"size"`
	PageCount int `json:"pageCount"`
	Content []NewsPiece `json:"content"`
}



