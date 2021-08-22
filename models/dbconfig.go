package models

//DbConfig is used to record some database info
type DbConfig struct {
	ID int `gorm:"primaryKey;column:Id" json:"id" form:"id"`

	//key
	Key string `gorm:"column:Key"`

	//value
	Vale string `gorm:"column:Value"`
}

func (DbConfig)TableName() string {
	return "dbconfigs"
}

// GetDbConfigValue return the table dbconfig and the key
func GetDbConfigValue(key string) string {
	var dbconfig DbConfig
	db.Where("`Key` = ?",key).First(&dbconfig)
	return dbconfig.Vale
}

func SetDbConfigValue(key string,value string)  {
	db.Model(&DbConfig{}).Where("`Key` = ?",key).Update("`Value`",value)
}