package ipblocker

import (
	"sync"
	"time"
)

//rate is a ConcurrentDictionary
var rate *RateLimit


func init()  {
	rate =NewRateLimit()
}

//Success set the ips struct to refresh
func Success(ip string)  {
	rate.Lock()
	defer rate.Unlock()
	result:=rate.GetorAdd(ip,NewLockInfo())
	result.Success()
}

//Fail add one try count to this ips struct
func Fail(ip string)  {
	rate.Lock()
	defer rate.Unlock()
	result:=rate.GetorAdd(ip,NewLockInfo())
	result.Fail()
}

//IsLoginable check the ips status and return it can log in  or not
func IsLoginable(ip string) bool {
	rate.Lock()
	defer rate.Unlock()
	result:=rate.GetorAdd(ip,NewLockInfo())
	return result.IsLoginable()
}

//limits is the interface for real ips struct
type limits interface {
	Fail()
	Success()
	IsLoginable()bool
}


//RateLimit is just a ConcurrentDictionary
type RateLimit struct {
	RateMap map[string]limits
	sync.Mutex
}

func NewRateLimit() *RateLimit {
	return &RateLimit{
		RateMap: make(map[string]limits),
	}
}
//GetorAdd is a useful function in c#, if it doesn't have one, then add one
func (rate *RateLimit)GetorAdd(ip string,limit limits) limits {

	if l,ok:=rate.RateMap[ip];ok {
		return l
	} else {
		rate.RateMap[ip]=limit
		return limit
	}
}



//LockInfo is the struct for block ip
type LockInfo struct {
	//尝试次数
	TryCount int
	//封禁时间
	BlockedTime time.Time

}

func NewLockInfo() *LockInfo {
	return &LockInfo{
		TryCount: 0,
		BlockedTime: time.Time{},
	}
}
func (lock *LockInfo)Fail(){
	lock.TryCount++
	if lock.TryCount>5{
		lock.TryCount=0
		lock.BlockedTime=time.Now()
	}
}

func (lock *LockInfo) Success() {
	lock.TryCount=0
}

func (lock *LockInfo)IsLoginable() bool {
	return lock.BlockedTime.Add(5*time.Minute).Before(time.Now())&&lock.TryCount<=5
}