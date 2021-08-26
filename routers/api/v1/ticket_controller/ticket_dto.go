package ticket_controller

import (
	"time"
	"xms/models"
)

type QueryParam struct {
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageCount int  `json:"pageCount" form:"pageCount"`
	ApOnly bool `json:"apOnly" form:"apOnly"`
}

func (q *QueryParam) IsEmpty() bool {
	return q.PageIndex==0&&q.PageCount==0&&q.ApOnly==false
}

type NoteVM struct {
	Id int  `json:"id"`
	Op string `json:"op"`
	Type models.NoteType `json:"type"`
	Content string `json:"content"`
	CreatedTime time.Time `json:"createdTime"`
	Avatar string `json:"avatar"`
}

type TicketInfoRes struct {
	Id int `json:"id" form:"id"`

	Type models.TicketType `json:"type" form:"type"`

	Device string `json:"device" form:"device" binding:"required,max=50"`

	DeviceModel string `json:"deviceModel" form:"deviceModel" binding:"required,max=50"`

	Owner string `json:"owner" form:"owner" binding:"required,max=20"`

	Phone string `json:"phone" form:"phone" binding:"required,max=20"`

	Description string `json:"description" form:"description" binding:"required,max=30"`

	Workers []string `json:"workers"`

	CreatedTime time.Time `json:"createdTime"`

	Status models.TicketStatus `json:"status"`

	IsConfirmed bool `json:"isConfirmed"`

	Notes []NoteVM `json:"notes"`

	Accessories []int `json:"accessories"`

	Picked bool `json:"picked"`


}

type MigrateTicketWorkerInfo struct {
	TicketID int `json:"ticketId" form:"ticketId"`
	WorkerID int `json:"workerId" form:"workerId"`
}

func (m *MigrateTicketWorkerInfo) IsFull() bool {
	return m.TicketID!=0&&m.WorkerID!=0
}

type TicketInfoReq struct {
	Type models.TicketType `json:"type" form:"type" `

	Device string  `json:"device" form:"device" binding:"max=50"`

	DeviceModel string `json:"deviceModel" form:"deviceModel" binding:"max=50"`

	Owner string `json:"owner" form:"owner" binding:"max=20"`

	Phone string `json:"phone" form:"phone" binding:"max=20"`

	Description  string `json:"description" form:"description" binding:"max=30"`

	Accessories []int `json:"accessories" form:"accessories"`

}

type TicketBriefInfo struct {
	Id int `json:"id" form:"id"`

	Type models.TicketType  `json:"type" form:"type"`

	Status models.TicketStatus `json:"status" form:"status"`

	Device string `json:"device" form:"device"`

	DeviceModel string `json:"deviceModel" form:"deviceModel"`

	Owner string `json:"owner" form:"owner"`

	Phone string `json:"phone" form:"phone"`

	CreatedTime time.Time `json:"createdTime" form:"createdTime"`

	IsConfirmed bool `json:"isConfirmed" form:"isConfirmed"`
}

type TicketListRes struct {
	PageIndex int `json:"pageIndex"`

	PageCount int `json:"pageCount"`

	Size int `json:"size"`

	QueryType int `json:"queryType"`

	Data []TicketBriefInfo `json:"data"`
}


type ChangeStatusBodyReq struct {
	Status models.TicketStatus `json:"status" form:"status" binding:"min=0,max=8"`
}

type NewCommentReq struct {

	Content string `json:"content" form:"content" binding:"required,max=500"`
}