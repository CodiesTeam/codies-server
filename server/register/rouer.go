package register

import (
	"net/http"

	"github.com/CodiesTeam/codies-server/skeleton/input/body"
	"github.com/CodiesTeam/codies-server/skeleton/reply"
	"github.com/CodiesTeam/codies-server/skeleton/route"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
)

func NewRoutes() []*route.Route {
	return []*route.Route{
		route.NewRoute(
			"/reg",
			"POST",
			register,
		),
	}
}

func register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	p := struct {
		Email string `json:"email"`
		PWD   string `json:"pwd"`
	}{}
	err := body.JSONBody(r, &p)
	if err != nil {
		glog.Errorf("JSONBody failed: %v", err)
	}
	glog.Infof("email: %s, pwd: %s", p.Email, p.PWD)
	user, err := regByEmail(p.Email, p.PWD)
	if err != nil {
		glog.Errorf("regByEmail failed: %v", err)
	}
	response := reply.JSON(map[string]interface{}{"user": user})
	response(w)

	// response := reply.JSON(M{"wel": "this is powered by httprouter"})
	// response(w)
}
