package models

import "time"

//WechatUser is the table wechatusers
type WechatUser struct {
	ID int `gorm:"column:Id"`

	OpenId string `gorm:"column:OpenId"`

	NickName string `gorm:"column:NickName"`

	Gender int `gorm:"column:Gender"`

	Language string `gorm:"column:Language"`

	City string `gorm:"column:City"`

	Province string `gorm:"column:Province"`

	Country string `gorm:"column:Country"`

	AvatarUrl string `gorm:"column:AvatarUrl"`

	CreatedTime time.Time `gorm:"column:CreatedTime"`
}

func (WechatUser)TableName() string  {
	return "wechatusers"
}

//NewWechatUser return a default WechatUser ptr
func NewWechatUser() *WechatUser {
	return &WechatUser{
		CreatedTime: time.Now(),
	}
}

type GenderType uint8
const (
	Unknown GenderType =iota
	Male 
	Female
)
//GetGender returns the user's gender
func (user *WechatUser)GetGender() GenderType {
	switch user.Gender {
	case 1:
		return Male
	case 2:
		return Female
	default:
		return Unknown

	}
}
