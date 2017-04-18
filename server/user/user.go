package user

import (
	"codies-server/skeleton/common"
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/golang/glog"
	"github.com/pborman/uuid"
)

type User struct {
	ID    int    `orm:"column(id);auto" json:"-"`
	UUID  string `orm:"column(uuid);size(36)" json:"uuid"`
	Name  string `orm:"column(name);size(50)" json:"name"`
	Phone string `orm:"column(phone);size(16);null" json:"phone,omitempty"`
	Email string `orm:"column(email);size(45);null" json:"email,omitempty"`
	Bio   string `orm:"column(bio);null" json:"bio,omitempty"`
}

func (u *User) TableName() string {
	return "user"
}

func NewUser(name string) *User {
	return &User{
		UUID: uuid.New(),
		Name: name,
	}
}

func (u *User) isValid() bool {
	return u.UUID != "" &&
		u.Name != "" &&
		(u.Email != "" || u.Phone != "")
}

// Insert insert User to DB
// TODO: test
func (u *User) Insert() error {
	if !u.isValid() {
		return fmt.Errorf("%s record %#v is not valid", u.TableName(), u)
	}
	o := orm.NewOrm()
	_, err := o.Insert(u)
	if err != nil {
		return err
	}
	return nil
}

func UserByUUID(uuid string) (*User, error) {
	sql := fmt.Sprintf(`select * from user where uuid = '%s'`,
		uuid)
	glog.V(2).Infof("UserByUUID sql: %s", sql)
	var user User
	err := orm.NewOrm().Raw(sql).QueryRow(&user)
	glog.Infof("UserByUUID, user: %#v, err: %v", user, err)
	if err != nil {
		return nil, err
	}
	if !user.isValid() {
		return nil, common.InvalidArgumentErr("user %#v is not valid", user)
	}
	return &user, nil
}
