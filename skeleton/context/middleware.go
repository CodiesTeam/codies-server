package context

type Middleware func(ctx *Context)

func WrapMiddleWare(ctx *Context, middles ...Middleware) {
	if ctx.Resp.ReplyFunc != nil {
		ctx.Reply()
		return
	}
	for _, mid := range middles {
		mid(ctx)
	}
}
