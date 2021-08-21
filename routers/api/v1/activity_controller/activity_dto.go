package activity_controller

import "time"

type ActivityPiece struct {

	Id int `json:"id"`
	//content need to be set to "{userName} {noteString}"
	Content string `json:"content"`

	//these three fields is not used for response
	UserId int `json:"-"`
	UserName string `json:"-"`
	NoteString string `json:"-"`
	
	UpdateAt time.Time `json:"updateAt"`
}


// ActivityRes is the response struct for activity controller
type ActivityRes struct {
	Count int `json:"count"`
	PageId int `json:"pageId"`
	Content []ActivityPiece `json:"content"`
}