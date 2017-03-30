package user_test

import (
	"fmt"
	"testing"

	"codies-server/server/user"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	fmt.Println("testMain...")
	orm.RegisterDataBase("default", "mysql", "root:codies-pwd@tcp(127.0.0.1:3306)/codies?charset=utf8", 30)
	orm.RegisterModel(new(user.User))
	fmt.Println("end register")

	m.Run()
}

func TestUser(t *testing.T) {
	convey.Convey("test User.Insert", t, func() {
		u := user.NewUser("testName")
		u.Email = "test@test.com"
		err := u.Insert()
		convey.So(err, convey.ShouldBeNil)
	})
	// TODO: delete test record
}
