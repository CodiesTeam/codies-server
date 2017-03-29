package context

import "github.com/CodiesTeam/codies-server/skeleton/reply"

type Context struct {
	Input *Param
	Resp  Response
}

type ProcessRequest func(*Context) reply.Replyer

func (c *Context) Reply() {
	c.Resp.ReplyFunc(c.Resp.ResponseWriter)
}

// // TODO: not good
// func ProcessRequest(ctx *Context) {
// 	if ctx.Resp.Reply != nil {
// 		ctx.Resp.ReplyFunc()
// 	}

// }
