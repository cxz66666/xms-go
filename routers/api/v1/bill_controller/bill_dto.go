package bill_controller

import (
	"time"
	"xms/models"
)

type CreateBillReq struct {
	Title string `json:"title" form:"title" binding:"required,max=100"`
	Content string `json:"content" form:"content" binding:"required,max=1000"`
	Type int `json:"type" form:"type" binding:"oneof=0 1"`
	UnitPrice int `json:"unitPrice" form:"unitPrice" binding:"required,min=0"`
	Quantity int `json:"quantity" form:"quantity" binding:"required,min=0"`
	Trader string `json:"trader" form:"trader" binding:"required,max=50"`
}


type BillTransactionsViewModel struct {
	UserId int `json:"userId" form:"userId"`
	UserName string `json:"userName" form:"userName"`
	UpdateAt time.Time `json:"updateAt" form:"updateAt"`
}

func NewBillTransactionsViewModel(record models.BillTransactionRecord,user models.User) BillTransactionsViewModel  {
	return BillTransactionsViewModel{
		UserId: record.OperatorId,
		UpdateAt: record.UpdateTime,
		UserName: user.Name,
	}
}

type BillViewModel struct {
	Id           int                         `json:"id"`
	Title        string                      `json:"title"`
	Content      string                      `json:"content"`
	Type         models.BillType             `json:"type"`
	UnitPrice    int                         `json:"unitPrice"`
	Quantity     int                         `json:"quantity"`
	Trader       string                      `json:"trader"`
	Transactions []BillTransactionsViewModel `json:"transactions"`
}

// NewBillViewModel create a BillViewModel froom models.Bill, just simply copy them!
func NewBillViewModel(b models.Bill) *BillViewModel {
	return &BillViewModel{
		Id: b.ID,
		Title: b.Title,
		Content: b.Content,
		Type: b.Type,
		UnitPrice: b.UnitPrice,
		Quantity: b.Quantity,
		Trader: b.Trader,
		Transactions: make([]BillTransactionsViewModel,0,len(b.Transactions)),
	}
}