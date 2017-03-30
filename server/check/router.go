package check

import (
	"codies-server/skeleton/context"
	"codies-server/skeleton/reply"
	"codies-server/skeleton/route"
)

func NewRouters() []*route.Router {
	return []*route.Router{
		&route.Router{
			"/hello/:abcd/world",
			"GET",
			hello,
		},
	}
}

func hello(ctx *context.Context) reply.Replyer {
	// glog.Infof("ps: %#v", ctx.Param.)
	return reply.ReplyJSON(M{"abc": "123"})
}
