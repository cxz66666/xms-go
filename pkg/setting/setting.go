

package setting
/**
Package setting is used to parse the app/app.ini and store the variable in the single struct
*/
import (
	"github.com/go-ini/ini"
	"log"
	"path/filepath"
	"reflect"
	"time"
	"xms/pkg/utils"
)
const (
	DebugMode = "debug"
	ReleaseMode = "release"
	TestMode = "test"
)
// we provide AppSetting ServerSetting DatabaseSetting AdminSetting SecretSetting
// to record the setting, and you need use Setup function to init them


type App struct {
	RuntimeRootPath  string
	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}
var AppSetting= &App{}


type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	JwtExpireTime time.Duration
}
var ServerSetting = &Server{}



type Database struct {
	Type string
	User string
	Password string
	Port string
	DatabaseName string
	TablePrefix string
}
var DatabaseSetting = &Database{}


type Admin struct {
	Username string
	Password string
}
var AdminSetting = &Admin{}


type Secret struct {
	JwtKey string
	JwtIssuer string
	AesKey string
	AesIv string
	WeChatAppId string
	WeChatAppSec string
	SaltA string
	SaltB string
}
var SecretSetting = &Secret{}

// Setup init the five setting struct, so before you use them, please
// use setting.Setup() to init them (only need once in the runtime)
func Setup()  {
	Cfg,err:=ini.Load("conf/app.ini")
	if err!=nil{
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	//---------------- app config ----------------------
	err=Cfg.Section("app").MapTo(AppSetting)
	if err!=nil{
		log.Fatalf("Fail to parse 'AppSetting': %v", err)
	}
	//change the '/' to '\' in Windows env, do nothing in Unix env
	AppSetting.RuntimeRootPath=filepath.FromSlash(AppSetting.RuntimeRootPath)
	AppSetting.LogSavePath=filepath.FromSlash(AppSetting.LogSavePath)


	//---------------- server config ----------------------
	err=Cfg.Section("server").MapTo(ServerSetting)
	if err!=nil{
		log.Fatalf("Fail to parse 'ServerSetting': %v", err)
	}
	// change int to second (simply plus time.Second)
	ServerSetting.ReadTimeout=ServerSetting.ReadTimeout*time.Second
	ServerSetting.WriteTimeout=ServerSetting.WriteTimeout*time.Second
	ServerSetting.JwtExpireTime=ServerSetting.JwtExpireTime*time.Minute
	// change the port to 80 if release, you can comment it if you don't want to use docker.
	if ServerSetting.RunMode==ReleaseMode {
		ServerSetting.HttpPort = 80
	}

	//---------------- database config ----------------------
	err=Cfg.Section("database").MapTo(DatabaseSetting)
	if err!=nil{
		log.Fatalf("Fail to parse 'DatabaseSetting': %v", err)
	}
	// set env by docker-compose.yml
	utils.SetStructByReflect(reflect.ValueOf(&DatabaseSetting),"DB_NAME","DatabaseName")
	utils.SetStructByReflect(reflect.ValueOf(&DatabaseSetting),"DB_PORT","Port")
	utils.SetStructByReflect(reflect.ValueOf(&DatabaseSetting),"DB_USER","User")
	utils.SetStructByReflect(reflect.ValueOf(&DatabaseSetting),"DB_PASSWORD","Password")


	//---------------- admin config ----------------------
	err=Cfg.Section("admin").MapTo(AdminSetting)
	if err!=nil{
		log.Fatalf("Fail to parse 'AdminSetting': %v", err)
	}
	utils.SetStructByReflect(reflect.ValueOf(&AdminSetting),"ADMIN_USERNAME","Username")
	utils.SetStructByReflect(reflect.ValueOf(&AdminSetting),"ADMIN_PASSWORD","Password")


	//---------------- secret config ----------------------
	err=Cfg.Section("secret").MapTo(SecretSetting)
	if err!=nil	{
		log.Fatalf("Fail to parse 'SecretSetting': %v", err)
	}
	utils.SetStructByReflect(reflect.ValueOf(&SecretSetting),"SALT_A","SaltA")
	utils.SetStructByReflect(reflect.ValueOf(&SecretSetting),"SALT_B","SaltB")
	utils.SetStructByReflect(reflect.ValueOf(&SecretSetting),"JWT_KEY","JwtKey")
	utils.SetStructByReflect(reflect.ValueOf(&SecretSetting),"JWT_ISSUER","JwtIssuer")
	utils.SetStructByReflect(reflect.ValueOf(&SecretSetting),"AES_KEY","AesKey")
	utils.SetStructByReflect(reflect.ValueOf(&SecretSetting),"AES_IV","AesIv")
	utils.SetStructByReflect(reflect.ValueOf(&SecretSetting),"WECHAT_APPID","WeChatAppId")
	utils.SetStructByReflect(reflect.ValueOf(&SecretSetting),"WECHAT_APPSEC","WeChatAppSec")


}

