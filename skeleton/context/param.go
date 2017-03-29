package context

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Param struct {
	Req  *http.Request
	vars httprouter.Params
	errs []string
}

func NewParam(req *http.Request, vars httprouter.Params) *Param {
	return &Param{
		Req:  req,
		vars: vars,
	}
}

// TODO: optional, data, required, and so on

func (p *Param) Required(key string, ret *string) *Param {

	return nil
}
