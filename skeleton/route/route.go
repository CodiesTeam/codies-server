package route

import (
	"net/http"

	"codies-server/skeleton/context"
	"codies-server/skeleton/kmux"

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

func BuildHandler(routeLists ...[]*Route) http.Handler {
	router := kmux.New()

	for _, routes := range routeLists {
		for _, rou := range routes {
			handler := func(route *Route) func(w http.ResponseWriter, r *http.Request, ps kmux.Params) {
				return func(w http.ResponseWriter, r *http.Request, ps kmux.Params) {
					ctx := &context.Context{
						Input: context.NewParam(r, ps),
						Resp:  context.NewResponse(w),
					}
					replyer := route.Handle(ctx)
					ctx.Resp.ReplyFunc = replyer
					ctx.Reply()
				}
			}(rou)

			router.Register(rou.Pattern, rou.Method, handler)
		}
	}

	// TODO: add serverfile route

	// use middleware
	n := negroni.Classic()
	n.UseHandler(router)
	return n
}
