package authorize_test

import (
	"fmt"
	"testing"

	"github.com/CodiesTeam/codies-server/server/authorize"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	fmt.Println("testMain...")
	orm.RegisterDataBase("default", "mysql", "root:codies-pwd@tcp(127.0.0.1:3306)/codies?charset=utf8", 30)
	orm.RegisterModel(new(authorize.LocalAuth))
	fmt.Println("end register")

	m.Run()
}

func TestLocalAuth(t *testing.T) {
	convey.Convey("test LocalAuth.Insert", t, func() {
		local := authorize.NewLocalAuth("test-uuid", "test-email", "", "test-pwd")
		err := local.Insert()
		convey.So(err, convey.ShouldBeNil)
		// TODO: delete test record
	})
}
