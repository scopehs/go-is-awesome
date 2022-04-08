// User Model Struct
// I love Struct, they are awesome!
package models


type User struct {
	Id        uint
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  []byte
}
