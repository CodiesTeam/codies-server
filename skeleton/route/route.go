package route

import (
	"net/http"

	"github.com/CodiesTeam/codies-server/skeleton/context"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

// type ProcessRequest func(*http.Request, httprouter.Params) reply.Replyer

type Router struct {
	Pattern string
	Method  string
	Handle  context.ProcessRequest
}

func NewRouter(pattern, method string, handle context.ProcessRequest) *Router {
	return &Router{
		Pattern: pattern,
		Method:  method,
		Handle:  handle,
	}
}

// Route basic route
type Route struct {
	Pattern string
	Method  string
	Handle  httprouter.Handle
}

func NewRoute(pattern, method string, handle httprouter.Handle) *Route {
	return &Route{
		Pattern: pattern,
		Method:  method,
		Handle:  handle,
	}
}

func BuildHandler(routers []*Router, routes ...[]*Route) http.Handler {
	router := httprouter.New()

	for _, rou := range routers {
		handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			ctx := &context.Context{
				Input: context.NewParam(r, ps),
				Resp:  context.NewResponse(w),
			}
			replyer := rou.Handle(ctx)
			ctx.Resp.ReplyFunc = replyer
			ctx.Reply()
			// replyFunc := rou.Handle(r, ps)
			// replyFunc(w)
		}
		router.Handle(rou.Method, rou.Pattern, handle)
	}

	for _, rs := range routes {
		for _, route := range rs {
			router.Handle(route.Method, route.Pattern, route.Handle)
		}
	}

	// TODO: add serverfile route

	// use middleware
	n := negroni.Classic()
	n.UseHandler(router)
	return n
}

// type Router struct {
// 	Pattern string
// 	Method  string
// 	Handle  httprouter.Handle
// }
