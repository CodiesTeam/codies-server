package check

import (
	"github.com/CodiesTeam/codies-server/skeleton/context"
	"github.com/CodiesTeam/codies-server/skeleton/reply"
	"github.com/CodiesTeam/codies-server/skeleton/route"
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
