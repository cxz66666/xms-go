package models

import (
	"gorm.io/gorm/clause"
	"time"
)

// TicketType means the 类型
type TicketType uint8
const (
	//电器
	Appliance TicketType =iota
	//电脑
	Computer
)

// TicketAccessory means the 附件
type TicketAccessory uint8
const (
	//无附件
	None TicketAccessory =iota
	//U盘
	UsbDisk
	//鼠标/接收器
	Mouse
	//电源适配器
	Power
	//其他（于评论中详述）
	Others
)

func (ta TicketAccessory) GetIndex() int {
	return int(ta)
}

func (ta TicketAccessory) ToDisplayName() string {
	switch ta {
	case None:
		return "无附件"
	case UsbDisk:
		return "U盘"
	case Mouse:
		return "鼠标/接收器"
	case Power:
		return "电源适配器"
	case Others:
		return "其他（于评论中详述）"
	default:
		return "未知"
	}
}

// Accessory is the table accessories
type Accessory struct {
	ID int `gorm:"primaryKey;column:Id" json:"id" form:"id"`

	//表单外键ID
	TicketID int `gorm:"column:TicketId" json:"ticket_id" form:"ticket_id"`

	//附件类型
	Type TicketAccessory `gorm:"column:Type" json:"type" form:"type"`
}


func (Accessory) TableName()string {
	return "accessories"
}


//TicketStatus means the 工单状态
type TicketStatus uint8
const (
	//已创建/交接中
	Created TicketStatus =iota
	//维修中
	Fixing
	//劝退待取回
	NonFixedWaitingg
	//劝退已取回
	NonFixedDone
	//维修成国公待取回
	SuccessWaiting
	//维修成功已取回
	SuccessDone
	//维修翻车待取回
	FailedWaiting
	//维修翻车已取回
	FailedDone
	//工单已作废
	Deleted

)
func (ts TicketStatus) GetIndex() int {
	return int(ts)
}

func (ts TicketStatus) ToDisplayName() string {
	switch ts {
	case Created:
		return "已创建/交接中"
	case Fixing:
		return "维修中"
	case NonFixedWaitingg:
		return "劝退待取回"
	case NonFixedDone:
		return "劝退已取回"
	case SuccessWaiting:
		return "维修成功待取回"
	case SuccessDone:
		return "维修成功已取回"
	case FailedWaiting:
		return "维修翻车待取回"
	case FailedDone:
		return "维修翻车已取回"
	case Deleted:
		return "工单已作废"

	default:
		return "未知"
	}
}

func CheckIfLockable(status TicketStatus) bool {
	return status==FailedDone||status==SuccessDone||status==NonFixedDone||status==Deleted
}


type NoteType uint8
const (

	Create NoteType =iota
	Join
	ChangeState
	Comment
	Edit
)

func (nt NoteType) GetIndex() int {
	return int(nt)
}

func (nt NoteType) ToDisplayName() string {
	switch nt {
	case Create:
		return "创建"
	case Join:
		return "认领"
	case ChangeState:
		return "改变状态"
	case Comment:
		return "评论"
	case Edit:
		return "修改"

	default:
		return "未知"
	}
}

// Note is the table notes
type Note struct {
	ID int  `gorm:"primaryKey;column:Id" json:"id" form:"id"`

	//操作者ID
	OperatorId int `gorm:"column:OperatorId" json:"operator_id" form:"operator_id"`

	//Note类型
	Type NoteType `gorm:"column:Type" json:"type" form:"type"`

	//内容
	Content string `gorm:"column:Content" json:"content" form:"content"`

	//创建时间
	CreatedTime time.Time `gorm:"column:CreatedTime" json:"created_time" form:"created_time"`

	// foreign key to Ticket
	TicketID int `gorm:"column:TicketId"`
}

//NewNote is used to create a note with time
func NewNote() *Note {
	return &Note{
		CreatedTime: time.Now(),
	}
}
//SetCreatedTime set Note.CreatedTime to time.Now
func (note *Note) SetCreatedTime() {
	note.CreatedTime=time.Now()
}

func (Note) TableName() string {
	return "notes"
}

// GetNoteById return the note with specially id
func GetNoteById(id int) (*Note,error)  {
	note:=Note{}
	if err:=db.First(&note,id).Error;err!=nil{
		return &note,err
	}
	return &note,nil
}

// GetNotesDesc will return at most num notes by descending order of Note.CreatedTime, also there type is not be comment
func GetNotesDesc(num int) ([]Note,error)  {
	var notes []Note
	if err:=db.Where("Type <> ?",Comment).Order("CreatedTime desc").Limit(num).Find(&notes).Error;err!=nil{
		return nil,err
	}
	return notes,nil
}









// TicketWorker 新开一个表去存这个ticket的Worker
type TicketWorker struct {

	ID int `gorm:"primaryKey;column:Id" json:"id" form:"id"`

	//表单ID外键
	TicketID int `gorm:"column:TicketId" json:"ticket_id" form:"ticket_id"`

	//操作者ID
	WorkerID int `gorm:"column:WorkerId" json:"worker_id" form:"worker_id"`
}

func (TicketWorker) TableName() string {
	return "ticketworkers"
}


type Ticket struct {
	ID int `gorm:"primaryKey;column:Id" json:"id" form:"id"`

	//设备种类
	Type TicketType `gorm:"column:Type" json:"type" form:"type" binding:"required"`

	//设备品牌
	Device string `gorm:"column:Device;size:50" json:"device" form:"device" binding:"required,max=50"`

	//设备型号
	DeviceModel string `gorm:"column:DeviceModel;size:50" json:"device_model" form:"device" binding:"required,max=50"`

	//机主姓名
	Owner string `gorm:"column:Owner;size:20" json:"owner" form:"owner" binding:"required,max=20"`

	//电话
	Phone string `gorm:"column:Phone;size:20" json:"phone" form:"phone" binding:"required,max=20"`

	//问题描述
	Description string `gorm:"column:Description;size:30" json:"description" form:"description" binding:"required,max=30"`

	//维修人员
	Workers []TicketWorker `gorm:"foreignKey:TicketID"`

	//备注
	Notes []Note `gorm:"foreignKey:TicketID"`

	//附件
	Accessories []Accessory `gorm:"foreignKey:TicketID"`

	//创建时间
	CreatedTime time.Time `gorm:"column:CreatedTime" json:"created_time" form:"created_time" `

	//状态
	Status TicketStatus `gorm:"column:Status" json:"status" form:"status"`

	//是否确认
	IsConfirmed bool `gorm:"column:IsConfirmed" json:"is_confirmed" form:"is_confirmed"`

	//wx id
	WeChatConfigId int `gorm:"column:	WeChatConfigId" json:"we_chat_config_id"`
}

func (Ticket) TableName() string {
	return "tickets"
}

//NewTicket simple init a ticket, fill it some fields to default value
func NewTicket() *Ticket {
	return &Ticket{
		IsConfirmed: false,
		WeChatConfigId: -1,
	}
}


// AddNewTicket will add the ticket to database, and return id if success
// I use ptr here, because we need a lot of id in the ticket
func AddNewTicket(ticket *Ticket) (int,error)  {
	if err:=db.Create(ticket).Error;err!=nil{
		return 0,err
	}
	return ticket.ID,nil
}


// GetTicketCount return the total tickets num
func GetTicketCount() (int,error) {
	var count int64
	if err:=db.Model(&Ticket{}).Count(&count).Error;err!=nil{
		return 0,err
	}
	return int(count),nil
}

// SearchTicket search the ticket which phone contain value or Owner contains value, then order by IsConfirmed and CretedTime
func SearchTicket(value string,pageId int,pageSize int, queryType int) ([]Ticket,int) {
	var tickets []Ticket
	var count int64
	pageId-=1
	tx:=db.Where("Phone LIKE ?","%"+value+"%").Or("Owner LIKE ?","%"+value+"%").
		Order("IsConfirmed , CreatedTime desc")
	switch queryType {
	case 1:
		tx=tx.Where("Type = ?",Appliance)
	case 2:
		tx=tx.Where("Type = ?",Computer)
	}
	tx.Count(&count)

	if err:=tx.Offset(pageId*pageSize).Limit(pageSize).Find(&tickets).Error;err!=nil{
			return nil,0
	}

	pageCount:=int(count)/pageSize
	if int(count)%pageSize!=0{
		pageCount++
	}


	return tickets,pageCount
}

func GetTicketsList(pageId int,pageSize int,queryType int) ([]Ticket,int)  {
	var tickets []Ticket
	var count int64
	pageId-=1
	tx:=db.Model(&Ticket{}).Order("IsConfirmed,CreatedTime desc")
	switch queryType {
	case 1:
		tx=tx.Where("Type = ?",Appliance)
	case 2:
		tx=tx.Where("Type = ?",Computer)
	}
	tx.Count(&count)

	if err:=tx.Offset(pageId*pageSize).Limit(pageSize).Find(&tickets).Error;err!=nil{
		return nil,0
	}

	pageCount:=int(count)/pageSize
	if int(count)%pageSize!=0{
		pageCount++
	}
	return tickets,pageCount
}


// GetTicketById find the id in the database with Transactions, and return error if not exists
func GetTicketById(id int) (*Ticket,error) {
	var ticket Ticket
	if err:=db.Preload("Notes").Preload("Workers").Preload("Accessories").First(&ticket,id).Error;err!=nil{
		return nil,err
	}
	return &ticket,nil
}

// UpdateTicket is used to update all the fields for ticket
func UpdateTicket(ticket *Ticket) error  {
	if err:=db.Select("*").Updates(ticket).Error;err!=nil{
		return err
	}
	return nil
}

func DeleteTicketById(id int) error  {
	if err:=db.Select(clause.Associations).Delete(&Ticket{ID: id}).Error;err!=nil{
		return err
	}
	return nil
}

// TicketsAndCount only used for cache to store the index page tickets
type TicketsAndCount struct {
	Tickets 	[]Ticket
	PageCount 	int
}