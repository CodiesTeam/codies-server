package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/CodiesTeam/codies-server/server/authorize"
	"github.com/CodiesTeam/codies-server/server/check"
	"github.com/CodiesTeam/codies-server/server/register"
	"github.com/CodiesTeam/codies-server/server/user"
	"github.com/CodiesTeam/codies-server/skeleton/route"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	orm.RegisterDataBase("default", "mysql", "root:codies-pwd@tcp(mysql:3306)/codies?charset=utf8", 30)
	orm.RegisterModel(new(user.User))
	orm.RegisterModel(new(authorize.LocalAuth))
}

const (
	serverPort = "8888"
)

func main() {
	// check.CheckMySQL()

	routes := check.NewRoutes()
	regRoutes := register.NewRoutes()
	handler := route.BuildHandler(routes, regRoutes)

	glog.Infof("start serving at %s", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, handler))
}
