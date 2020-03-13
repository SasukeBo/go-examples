package usegorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB _
var DB *gorm.DB

// Object _
type Object struct {
	Data interface{}
}

// User _
type User struct {
	Name string
	Age  int
	// Data string `gorm:"column:data; type:json" json:"data"`
}

// TableName _
func (u User) TableName() string {
	return "user"
}

// ConnectDB _
func ConnectDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:Wb922149@...S@/test_db")
	if err != nil {
		panic(err)
	}
	DB.LogMode(true)
}

// CreateAUser _
func CreateAUser() User {
	user := User{"sasukebo", 25}
	err := DB.Create(&user).Error
	if err != nil {
		panic(err)
	}

	return user
}

func DeleteUser(u interface{}) error {
	err := DB.Delete(u).Error
	return err
}

// QueryUser _
func QueryUser() {
	users := make([]User, 0)
	DB.Where("name = ?", "sasukebo").Where("age = ?", 25).Find(&users)
	fmt.Println(users)
}
