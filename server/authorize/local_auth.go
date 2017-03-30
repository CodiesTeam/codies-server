package authorize

import (
	"fmt"

	"codies-server/server/util/encrypt"

	"github.com/astaxie/beego/orm"
)

type LocalAuth struct {
	UUID     string `orm:"column(uuid);pk"`
	Email    string `orm:"column(email);size(45);null"`
	Phone    string `orm:"column(phone);size(16);null"`
	Password string `orm:"column(password);size(8)"`
}

func (t *LocalAuth) TableName() string {
	return "local_auth"
}

func NewLocalAuth(uuid, email, phone, pwd string) *LocalAuth {
	return &LocalAuth{
		UUID:     uuid,
		Email:    email,
		Phone:    phone,
		Password: password(uuid, pwd),
	}
}

func (l *LocalAuth) isValid() bool {
	return l.UUID != "" &&
		l.Password != "" &&
		(l.Email != "" || l.Phone != "")
}

func (l *LocalAuth) GetByEmail(email string) error {

	err := orm.NewOrm().QueryTable(l.TableName()).
		Filter("email", email).
		One(l)
	fmt.Printf("GetByEmail err: %v, l: %#v\n", l)
	if err != nil {
		return err
	}
	return nil
}

// Insert insert to DB
// TODO: complete , test
func (l *LocalAuth) Insert() error {
	if !l.isValid() {
		return fmt.Errorf("% record %#v is not valid", l.TableName(), l)
	}
	o := orm.NewOrm()
	_, err := o.Insert(l)
	if err != nil {
		return err
	}
	return nil
}

// TODO: test it
func password(uuid, pwd string) string {
	return encrypt.MD5Sum(uuid, pwd)[:8]
}
