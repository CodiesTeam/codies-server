package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/pborman/uuid"
)

type User struct {
	ID    int    `orm:"column(id);auto"`
	UUID  string `orm:"column(uuid);size(36)"`
	Name  string `orm:"column(name);size(50)"`
	Phone string `orm:"column(phone);size(16);null"`
	Email string `orm:"column(email);size(45);null"`
	Bio   string `orm:"column(bio);null"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

func NewUser(name string) *User {
	return &User{
		UUID: uuid.New(),
		Name: name,
	}
}
