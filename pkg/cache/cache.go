package cache

import (
	"github.com/muesli/cache2go"
	"time"
	"xms/pkg/logging"
	"xms/pkg/setting"
)


var cache *cache2go.CacheTable


func init() {
	cache=cache2go.Cache("xms")
}

// GetOrCreate accept a key and a function to return data (interface), if there isn't key-value in cache, it will call the
// function to get data and store it
func GetOrCreate(key string, f func()interface{}) interface{} {
	res,err:=cache.Value(key)
	if err!=nil{
		newValue:=f()
		cache.Add(key,setting.ServerSetting.CacheExpireTime,newValue)
		return newValue
	}
	return res.Data()
}

// SetDefault set kv to cache with the default expiredTime
func SetDefault(key string,value interface{})  {
	cache.Add(key,setting.ServerSetting.CacheExpireTime,value)
}

// Set force update the cache[key] (if exists), and set the expired time to custom
func Set(key string, value interface{}, expiredTime time.Duration)  {
	cache.Add(key,expiredTime,value)
}

// Remove a key in cache, no matter whether it exists
func Remove(key string)  {
	cache.Delete(key)
}



// ForceFlush wipe the entire cache table
func ForceFlush()  {
	logging.Info("force flush cache")
	cache.Flush()
}