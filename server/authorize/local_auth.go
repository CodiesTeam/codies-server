package models

import "github.com/astaxie/beego/orm"

type LocalAuth struct {
	UUID     string `orm:"column(uuid);pk"`
	Email    string `orm:"column(email);size(45);null"`
	Phone    string `orm:"column(phone);size(16);null"`
	Password string `orm:"column(password);size(8)"`
}

func (t *LocalAuth) TableName() string {
	return "local_auth"
}

func init() {
	orm.RegisterModel(new(LocalAuth))
}

func NewLocalAuth(uuid, email, phone, pwd string) *LocalAuth {
	return &LocalAuth{
		UUID:     uuid,
		Email:    email,
		Phone:    phone,
		Password: pwd,
	}
}
