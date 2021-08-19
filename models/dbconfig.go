package models

//DbConfig is used to record some database info
type DbConfig struct {
	ID int `gorm:"primaryKey" json:"id" form:"id"`

	//key
	Key string

	//value
	Vale string
}

func (DbConfig)TableName() string {
	return "dbconfigs"
}

