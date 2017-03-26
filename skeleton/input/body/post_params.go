package body

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func JSONBody(req *http.Request, obj interface{}) error {
	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, obj)
	if err != nil {
		return err
	}
	return nil
}
