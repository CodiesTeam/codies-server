package main

import (
	"flag"
	"log"
	"net/http"

	"codies-server/server/authorize"
	"codies-server/server/check"
	"codies-server/server/home"
	"codies-server/server/praise"
	"codies-server/server/topic"
	"codies-server/server/user"
	"codies-server/skeleton/route"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	initDB()
}

func initDB() {
	glog.Infoln("connect mysql...")
	orm.RegisterDataBase("default", "mysql", "root:codies-pwd@tcp(mysql:3306)/codies?charset=utf8", 30)
	orm.RegisterModel(new(user.User))
	orm.RegisterModel(new(authorize.LocalAuth))
	orm.RegisterModel(new(topic.Post))
	orm.RegisterModel(new(praise.Praise))
	glog.Infoln("mysql connected")
}

const (
	serverPort = "8888"
)

func main() {
	// check.CheckMySQL()

	routes := check.NewRoutes()
	regRoutes := home.NewRoute()
	topicRoutes := topic.NewRoute()

	handler := route.BuildHandler(
		routes,
		regRoutes,
		topicRoutes,
	)

	glog.Infof("start serving at %s", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, handler))
}
