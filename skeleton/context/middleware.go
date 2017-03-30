package context

type Middleware func(ctx *Context)

/*
wrap should like this:
wrap(ProcessRequest, ...Middleware)ProcessRequest
*/

func WrapMiddleWare(ctx *Context, middles ...Middleware) {
	if ctx.Resp.ReplyFunc != nil {
		ctx.Reply()
		return
	}
	for _, mid := range middles {
		mid(ctx)
	}
}
