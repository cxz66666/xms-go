package models

import (
	"fmt"
	"gorm.io/gorm/clause"
	"strconv"
	"strings"
	"time"
)

const(
	BasicAmountKey = "BILL_STATS_BASIC_AMOUNT"
	BasicCountKey = "BILL_STATS_BASIC_COUNT"
)

func GenerateBasicAmountString(balance int, incomes int, outgoings int) string {
	return fmt.Sprintf("%d|%d|%d",balance,incomes,outgoings)
}

func GenerateBasicCountString(incomes int, outgoings int) string {
	return fmt.Sprintf("%d|%d",incomes,outgoings)
}

func GetBasicAmountFromString(str string) (int,int,int) {
	str=strings.Trim(str," ")
	s:=strings.SplitN(str,"|",3)
	a,_:= strconv.Atoi(s[0])
	b,_:=strconv.Atoi(s[1])
	c,_:=strconv.Atoi(s[2])
	return a,b,c
}

func GetBasicCountFromString(str string) (int,int) {
	str=strings.Trim(str," ")
	s:=strings.SplitN(str,"|",2)
	a,_:= strconv.Atoi(s[0])
	b,_:=strconv.Atoi(s[1])
	return a,b
}


type BillType int

const (
	Income BillType =iota
	Outgoing
)
// BillTransactionRecord is the table billtranscationrecords
type BillTransactionRecord struct {
	ID int `gorm:"primaryKey;column:Id" json:"id" form:"id"`

	//操作人员ID
	OperatorId int `gorm:"column:OperatorId"`

	//修改的时间
	UpdateTime time.Time `gorm:"column:UpdateTime"`

	//Bill的外键
	BillID int `gorm:"column:BillId"`
}

func (BillTransactionRecord) TableName() string {
	return "billtranscationrecords"
}


// Bill is the table bills
type Bill struct {
	ID int `gorm:"primaryKey;column:Id" json:"id" form:"id"`

	//标题
	Title string `gorm:"column:Title;size:100" form:"title" json:"title" binding:"required,max=100,min=0"`

	//内容
	Content string `gorm:"column:Content;size:1000"  form:"content;" json:"content" binding:"required,max=1000"`

	//收入/支出
	Type BillType `gorm:"column:Type" form:"type" json:"type" binding:"required,oneof=Income Outgoing"`

	//单价
	UnitPrice int `gorm:"column:UnitPrice" form:"unit_price" json:"unit_price" binding:"required,min=0"`

	//数量
	Quantity int `gorm:"column:Quantity" form:"quantity" json:"quantity" binding:"required,min=0"`

	//参与交易人员
	Trader string `gorm:"column:Trader;size:50" form:"trader" json:"trader" binding:"required,max=50"`

	Transactions []BillTransactionRecord `gorm:"foreignKey:BillID"`
}

func (Bill) TableName() string {
	return "bills"
}

// AddNewBill will add the bill to database, and return id if success
func AddNewBill(bill Bill) (int,error)  {
	if err:=db.Create(&bill).Error;err!=nil{
		return 0,err
	}
	return bill.ID,nil
}

// GetBillById find the id in the database with Transactions, and return error if not exists
func GetBillById(id int) (*Bill,error) {
	var bill Bill
	if err:=db.Preload("Transactions").First(&bill,id).Error;err!=nil{
		return nil,err
	}
	return &bill,nil

}

//DeleteBillById delete the bill model in database
func DeleteBillById(id int) error {
	if err:=db.Select("Transactions").Delete(&Bill{ID: id}).Error;err!=nil{
		return err
	}
	return nil
}

//UpdateBill will update all the column of bill
func UpdateBill(bill Bill) error  {
	db.Model(&bill).Association("Transactions").Replace(bill.Transactions)
	if err:=db.Omit(clause.Associations).Updates(&bill).Error;err!=nil{
		return err
	}
	return nil
}

//GetBillsPaginated get the bills according to the type id size, but important Id range from 1 to inf, not 0, so don't forget to minus 1
func GetBillsPaginated(billType int,pageId int, pageSize int) ([]Bill,int) {
	var bills []Bill
	var count int64
	pageId	-=1
	tx:=db
	switch billType {
	case 0:
		tx.Where("Type = ?",Income)
	case 1:
		tx.Where("Type = ?",Outgoing)
	}
	tx.Count(&count)
	tx.Order("Id desc").Preload("Transactions").Offset(pageSize*pageId).Limit(pageSize).Find(&bills)
	pageCount:=int(count)/pageSize

	if int(count)%pageSize!=0	{
		pageCount++
	}
	return bills,pageCount
}