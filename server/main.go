package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"codies-server/server/authorize"
	"codies-server/server/home"
	"codies-server/server/praise"
	"codies-server/server/topic"
	"codies-server/server/user"
	"codies-server/skeleton/route"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
)

const (
	serverPort = "8888"
	// mysqlAddr  = "172.17.0.3:3306"
	mysqlAddr = "docker_mysql_1:3306"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "2")
	flag.Parse()

	initDB()
}

func initDB() {
	glog.Infoln("connect mysql...")
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("root:codies-pwd@tcp(%s)/codies?charset=utf8", mysqlAddr), 30)
	orm.RegisterModel(new(user.User))
	orm.RegisterModel(new(authorize.LocalAuth))
	orm.RegisterModel(new(topic.Post))
	orm.RegisterModel(new(praise.Praise))
	glog.Infoln("mysql connected")
}

func main() {
	// check.CheckMySQL()

	regRoutes := home.NewRoute()
	topicRoutes := topic.NewRoute()

	handler := route.BuildHandler(
		regRoutes,
		topicRoutes,
	)

	glog.Infof("start serving at %s", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, handler))
}
