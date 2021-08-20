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

