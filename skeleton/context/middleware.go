package context

import (
	"codies-server/skeleton/reply"

	"github.com/golang/glog"
)

// type Middleware func(ctx *Context)

// /*
// wrap should like this:
// wrap(ProcessRequest, ...Middleware)ProcessRequest
// */

// // TODO: seem not correct
// func WrapMiddleWare(ctx *Context, middles ...Middleware) {
// 	if ctx.Resp.ReplyFunc != nil {
// 		ctx.Reply()
// 		return
// 	}
// 	for _, mid := range middles {
// 		mid(ctx)
// 	}
// }

func BasicAuth(user, password string, f ProcessRequest) ProcessRequest {
	return func(ctx *Context) reply.Replyer {
		inputUser, inputPWD, hasAuth := ctx.Input.Req.BasicAuth()
		glog.Infoln("user, pwd and hasAuth: ", user, password, hasAuth)
		if hasAuth && inputUser == user && inputPWD == password {
			return f(ctx)
		}
		return reply.BasicAuth()
	}
}
