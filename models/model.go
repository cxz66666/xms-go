package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"xms/pkg/setting"
)

//db is the single database instance
var db *gorm.DB

//Setup init the db, so you need use it before you use db
func Setup()  {
	var (
		dbType, dbName, user, password, host, tablePrefix string
		err error
	)
	dbType=setting.DatabaseSetting.Type
	dbName=setting.DatabaseSetting.DbName
	user=setting.DatabaseSetting.User
	password=setting.DatabaseSetting.Password
	host=setting.DatabaseSetting.Host
	tablePrefix=setting.DatabaseSetting.TablePrefix

	dsn:=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user,password,host,dbName)
	switch dbType {
	case "mysql":
		db,err=gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgresql":
		db, err=gorm.Open(postgres.Open(dsn), &gorm.Config{})

	default:
		log.Fatalf("sorry we can't support type of database: %s",dbType)
	}
	if err!=nil{
		log.Fatalf("connect to database error, connect string is %s, details is : %v",dsn,err)
	}
	if len(tablePrefix)>0{
		fmt.Printf("[warning] tablePrefix '%s' will be nothing to do in current version",tablePrefix)
	}
	//auto migrate  it can't handle the dependency relations, so you need handle it by yourself
	db.AutoMigrate(&DbConfig{},&News{},&Ticket{},&TicketWorker{},&User{},&WechatConfig{},&WechatUser{})
	db.AutoMigrate(&Bill{})
	db.AutoMigrate(&Accessory{},&Note{},&BillTransactionRecord{})
}



//Enum is a simple enum type, only have two methods
type Enum interface {
	//GetIndex return the index by int
	GetIndex() int
	//ToDisplayName return the human read name
	ToDisplayName()string
}