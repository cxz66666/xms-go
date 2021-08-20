package models

import (
	"fmt"
	"math/rand"
	"time"
)

type WechatConfig struct {
	ID int `gorm:"column:Id"`

	Phone string `gorm:"column:Phone"`

	Code int `gorm:"column:Code"`

	TicketID int `gorm:"column:TicketId"`

	Version int `gorm:"column:Version"`

	ExpiresAt time.Time `gorm:"column:ExpiresAt"`

	CommentCount int `gorm:"column:CommentCount"`
}

func (WechatConfig) TableName() string {
	return "wechatconfigs"
}


//GetCodeString returns the code with format code:0000
func (config *WechatConfig) GetCodeString() string {
	return fmt.Sprintf("%d:0000",config.Code)
}

//NewWechatConfig
func NewWechatConfig() *WechatConfig {
	return &WechatConfig{
		Code: 	rand.Int()%10000,
		Version: 2,
		ExpiresAt: time.Now().AddDate(0,6,0),
		CommentCount: 0,
	}
}
