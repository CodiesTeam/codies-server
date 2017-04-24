package context

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"codies-server/skeleton/common"
	"codies-server/skeleton/kmux"
)

type Param struct {
	Req  *http.Request
	vars map[string]string
	errs []string
}

func NewParam(req *http.Request, vars kmux.Params) *Param {
	params := make(map[string]string, len(vars))
	for _, p := range vars {
		params[p.Key] = p.Value
	}
	return &Param{
		Req:  req,
		vars: params,
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
	ret, ok := p.vars[key]
	if !ok {
		p.AddError(fmt.Sprintf("path var %s not set", key))
		return p
	}
	*result = ret
	return p
}

func (p *Param) Optional(key string, result *string) *Param {
	ret, ok := p.vars[key]
	if ok {
		*result = ret
	}
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
