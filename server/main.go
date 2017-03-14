package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/CodiesTeam/codies-server/server/check"
	"github.com/CodiesTeam/codies-server/skeleton/route"
	"github.com/golang/glog"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()
}

const (
	serverPort = "8888"
)

func main() {
	// check.CheckMySQL()

	routes := check.NewRoutes()
	handler := route.BuildHandler(routes)

	glog.Infof("start serving at %s", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, handler))
}
