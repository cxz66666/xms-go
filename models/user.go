package models

import "strconv"

type Role uint8
const (
	//干事
	Staff Role =iota
	//部长
	Minister
	//会长
	President
	//副会长
	VicePresident
	//技术指导
	TechAdviser
	//已退休
	Retired
	//维护管理员
	Sysadmin
	//副部长
	ViceMinister
)
func (role Role) GetIndex() int {
	return int(role)
}

func (role Role) ToDisplayName() string {
	switch role {
	case Staff:
		return "干事"
	case Minister:
		return "部长"
	case President:
		return "会长"
	case VicePresident:
		return "副会长"
	case TechAdviser:
		return "技术指导"
	case Retired:
		return "已退休"
	case Sysadmin:
		return "系统管理员"
	case ViceMinister:
		return "副部长"

	default:
		return "未知"
	}
}



type Department uint8

const (
	//电器部
	DQ Department =iota
	//电脑部
	DN
	//人资部
	RZ
	//文宣部
	WX
	//财外部
	CW
)

func (department Department) GetIndex() int {
	return int(department)
}

func (department Department) ToDisplayName() string {
	switch department {
	case DQ:
		return "电器部"
	case DN:
		return "电脑部"
	case RZ:
		return "人资部"
	case WX:
		return "文宣部"
	case CW:
		return "财外部"

	default:
		return "未知"
	}
}



// User is the table users
type User struct {
	ID int `gorm:"primaryKey;column:Id"`

	//学号
	StudentId int `gorm:"column:StudentId"`

	//名字
	Name string `gorm:"column:Name;index:indexName"`

	//密码 hash后的
	Secret string `gorm:"column:Secret"`

	//角色Role
	Role Role `gorm:"column:Role"`

	//部门
	Department Department `gorm:"column:Department;index:indexDep"`

	//电脑维修件数
	ComputerFixedCount int `gorm:"column:ComputerFixedCount"`

	//电器件数
	ApplianceFixedCount int `gorm:"column:ApplianceFixedCount"`

	//头像URL
	AvatarURL string `gorm:"column:AvatarURL"`
}

func (User) TableName() string {
	return "users"
}

// NewUser return a user ptr for the default user
func NewUser() *User {
	return &User{
		ComputerFixedCount: 0,
		ApplianceFixedCount: 0,
	}
}

func NewAdminUser() *User {
	return &User{
		StudentId: 10086,
		Name:"系统维护管理员",
		Role: Sysadmin,
		Department: DN,
		ID: 1,
		ComputerFixedCount: 0,
		ApplianceFixedCount: 0,
	}
}

func (user *User)GetAvatarURL() string {
	//result:=int64(user.ID)
	//result = ((result ^ 1242458739) + 1984) ^ 4281719956;
	//return  strconv.Itoa(int(result))
	return strconv.Itoa(user.ID)
}

// GetUserById return the note with specially id
func GetUserById(id int) (*User,error)  {
	user:=User{}
	if err:=db.First(&user,id).Error;err!=nil{
		return &user,err
	}
	return &user,nil
}
// StuId2Id convert StudentId to ID
func StuId2Id(stuid int) int  {
	user:=User{}
	if err:=db.First(&user,"StudentId = ?",stuid).Error;err!=nil{
		return -1
	}
	return user.ID
}

// UpdateUser is used to update all the fields for user
func UpdateUser(user *User) error  {
	if err:=db.Select("*").Updates(user).Error;err!=nil{
		return err
	}
	return nil
}
