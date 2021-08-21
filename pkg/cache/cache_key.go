package cache

import "fmt"

type CacheType uint8

const (
	UserInfo CacheType=iota
	StuIdToUId
	Ticket
	Note
	News
	WechatConfig
	Stats
	TicketWorker
	Bill
)


type StatsCacheId uint8
const (
	TicketCount StatsCacheId =iota
	UserCount
)

//
func hash(t CacheType, id int ) string {
	return fmt.Sprintf("%d#!%d",t,id)
}

func GetKey(t CacheType, id int) string  {
	return hash(t,id)
}

func GetNewsListKey() string {
	return hash(News,-1)
}

func GetActivityListKey() string {
	return hash(Note,-1)
}

func GetTicketFirstPageKey() string {
	return hash(Ticket,-2)
}

func GetBillFirstPageKey() string {
	return hash(Bill,-1)
}