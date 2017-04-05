package route

import (
	"net/http"

	"codies-server/skeleton/context"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

type Route struct {
	Pattern string
	Method  string
	Handle  context.ProcessRequest
}

func NewRoute(pattern, method string, handle context.ProcessRequest) *Route {
	return &Route{
		Pattern: pattern,
		Method:  method,
		Handle:  handle,
	}
}

func BuildHandler(routers []*Route) http.Handler {
	router := httprouter.New()

	for _, rou := range routers {
		handler := func(route *Route) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
				ctx := &context.Context{
					Input: context.NewParam(r, ps),
					Resp:  context.NewResponse(w),
				}
				replyer := route.Handle(ctx)
				ctx.Resp.ReplyFunc = replyer
				ctx.Reply()
			}
		}(rou)

		router.Handle(rou.Method, rou.Pattern, handler)
	}

	// TODO: add serverfile route

	// use middleware
	n := negroni.Classic()
	n.UseHandler(router)
	return n
}
