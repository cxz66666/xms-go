package models

import (
	"fmt"
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
	ID int `gorm:"primaryKey" json:"id" form:"id"`

	//操作人员ID
	OperatorId int

	//修改的时间
	UpdateTime time.Time

	// foreign key to Bill
	BillID int
}

func (BillTransactionRecord) TableName() string {
	return "billtranscationrecords"
}


// Bill is the table bills
type Bill struct {
	ID int `gorm:"primaryKey" json:"id" form:"id"`

	//标题
	Title string `form:"title" json:"title" binding:"required,max=100,min=0"`

	//内容
	Content string `form:"content" json:"content" binding:"required,max=1000"`

	//收入/支出
	Type BillType `form:"type" json:"type" binding:"required,oneof=Income Outgoing"`

	//单价
	UnitPrice int `form:"unit_price" json:"unit_price" binding:"required,min=0"`

	//数量
	Quantity int `form:"quantity" json:"quantity" binding:"required,min=0"`

	//参与交易人员
	Trader string `form:"trader" json:"trader" binding:"required,max=50"`

	Transactions []BillTransactionRecord `gorm:"foreignKey:BillID"`
}

func (Bill) TableName() string {
	return "bills"
}