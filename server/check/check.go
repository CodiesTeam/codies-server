/*
this package is just use as a test
*/
package check

import (
	"codies-server/skeleton/context"

	"fmt"

	"codies-server/skeleton/reply"
	"codies-server/skeleton/route"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
)

type M map[string]interface{}

func NewRoutes() []*route.Router {
	return []*route.Router{
		route.NewRouter(
			"/",
			"GET",
			welcome,
		),
		route.NewRouter(
			"/redis",
			"GET",
			checkRedis,
		),
		route.NewRouter(
			"/any/:name",
			"GET",
			hi,
		),
		route.NewRouter(
			"/match/*filepath",
			"GET",
			matchAll,
		),
		route.NewRouter(
			"/hello/:abcd/world",
			"GET",
			hello,
		),
		// route.NewRouter(
		// 	"/protected",
		// 	"GET",
		// 	basicAuth(protected, "kang", "123!"),
		// ),
	}
}

func welcome(ctx *context.Context) reply.Replyer {
	glog.Infof("this is welcome, input is: %v", ctx.Input)
	return reply.ReplyJSON(M{"wel": "this is powered by httprouter"})
}

func hello(ctx *context.Context) reply.Replyer {
	// glog.Infof("ps: %#v", ctx.Param.)
	return reply.ReplyJSON(M{"abc": "123"})
}

// func welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	reply.JSON(w, M{"wel": "this is powered by httprouter"})
// }

func checkRedis(ctx *context.Context) reply.Replyer {
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	v, err := conn.Do("SET", "name", "red")
	if err != nil {
		glog.Error(err)
	}
	glog.Infoln("v: ", v)

	v, err = redis.String(conn.Do("GET", "name"))
	if err != nil {
		glog.Error(err)
	}
	return reply.ReplyJSON(M{"result": fmt.Sprintf("check redis, the value got is %s", v)})
}

// func checkRedis(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	conn, err := redis.Dial("tcp", "redis:6379")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()
// 	v, err := conn.Do("SET", "name", "red")
// 	if err != nil {
// 		glog.Error(err)
// 	}
// 	glog.Infoln("v: ", v)

// 	v, err = redis.String(conn.Do("GET", "name"))
// 	if err != nil {
// 		glog.Error(err)
// 	}
// 	reply.JSON(w, M{"result": fmt.Sprintf("check redis, the value got is %s", v)})
// }

func hi(ctx *context.Context) reply.Replyer {
	glog.Infof("params: %#v\n", ctx.Input)
	var name string
	if err := ctx.Input.Var("name", &name).Error(); err != nil {
		return reply.Err(err)
	}
	return reply.ReplyJSON(M{"result": fmt.Sprintf("hello %s", name)})
}

// func hi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	glog.Infof("params: %#v\n", ps)
// 	reply.JSON(w, M{"result": fmt.Sprintf("hello %s", ps.ByName("name"))})
// }

func matchAll(ctx *context.Context) reply.Replyer {
	var filepath string
	if err := ctx.Input.Var("filepath", &filepath).Error(); err != nil {
		return reply.Err(err)
	}
	return reply.ReplyJSON(M{"result": fmt.Sprintf("the filepath is %s", filepath)})
}

// func matchAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	reply.JSON(w, M{"result": fmt.Sprintf("the filepath is %s", ps.ByName("filepath"))})
// }

/*func basicAuth(ctx *context.Context) context.ProcessRequest {
	return func(ctx *context.Context)reply.Replyer {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()
		glog.Infoln("user, pwd and hasAuth: ", user, password, hasAuth)

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
func basicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()
		glog.Infoln("user, pwd and hasAuth: ", user, password, hasAuth)

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
*/
func protected(ctx *context.Context) reply.Replyer {
	return reply.ReplyJSON(M{"result": "this is protected ~"})
}
