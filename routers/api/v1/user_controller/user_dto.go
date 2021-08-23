package user_controller

import "xms/models"

type UserInfoRes struct {
	Name string `json:"name"`

	Avatar string `json:"avatar"`

	Userid int `json:"userid"`

	Email string `json:"email"`

	Title models.Role `json:"title"`

	Group models.Department `json:"group"`

	NotifyCount int `json:"notifyCount"`

	UnreadCount int `json:"unreadCount"`

}

type UserDetailsRes struct {
	Id int `json:"id"`

	StudentId  int `json:"studentId"`

	Name string `json:"name"`

	Avatar string `json:"avatar"`

	Email string `json:"email"`

	Title models.Role `json:"title"`

	Group models.Department `json:"group"`

	ComputerCount int `json:"computerCount"`

	ApplicanceCount int `json:"applicanceCount"`

}

type UserListRes struct {
	PageIndex int `json:"pageIndex"`
	
	PageCount int `json:"pageCount"`
	
	Size int `json:"size"`
	
	Role int `json:"role"`
	
	Department int `json:"department"`
	
	Data []UserDetailsRes `json:"data"`
}