package register

import (
	"codies-server/skeleton/context"
	"codies-server/skeleton/reply"
	"codies-server/skeleton/route"

	"github.com/golang/glog"
)

func NewRouter() []*route.Router {
	return []*route.Router{
		route.NewRouter(
			"/reg",
			"POST",
			register,
		),
	}
}

func register(ctx *context.Context) reply.Replyer {
	p := struct {
		Email string `json:"email"`
		PWD   string `json:"pwd"`
	}{}
	if err := ctx.Input.JSONBody(&p).Error(); err != nil {
		return reply.Err(err)
	}
	glog.Infof("email: %s, pwd: %s", p.Email, p.PWD)

	user, err := regByEmail(p.Email, p.PWD)
	if err != nil {
		glog.Errorf("regByEmail failed: %v", err)
		return reply.Err(err)
	}
	return reply.ReplyJSON(map[string]interface{}{
		"user": user,
	})
}
