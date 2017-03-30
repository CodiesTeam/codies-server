package context

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"codies-server/skeleton/common"

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

func (p *Param) AddError(msg string) {
	p.errs = append(p.errs, msg)
}

func (p *Param) Error() error {
	if len(p.errs) == 0 {
		return nil
	}
	return common.InvalidArgumentErr(strings.Join(p.errs, "\n"))
}

func (p *Param) Var(key string, result *string) *Param {
	var ret string
	if ret = p.vars.ByName(key); ret == "" {
		p.AddError(fmt.Sprintf("path var %s not set", key))
		return p
	}
	*result = ret
	return p
}

func (p *Param) JSONBody(obj interface{}) *Param {
	b, err := ioutil.ReadAll(p.Req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, obj)
	if err != nil {
		p.AddError(fmt.Sprintf("invalid body: %v", err.Error()))
	}
	return p
}

// TODO: optional, data, required, and so on

func (p *Param) Required(key string, ret *string) *Param {

	return nil
}
