package body

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"codies-server/skeleton/common"
)

func JSONBody(req *http.Request, obj interface{}) error {
	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, obj)
	if err != nil {
		return common.InvalidArgumentErr("invalid body: %v", err)
	}
	return nil
}
