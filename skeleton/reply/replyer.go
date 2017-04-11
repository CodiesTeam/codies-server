package reply

import (
	"encoding/json"
	"net/http"

	"codies-server/skeleton/common"
)

// Replyer write result to r
type Replyer func(w http.ResponseWriter)

func serverJSON(v interface{}) Replyer {
	return func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(v); err != nil {
			panic(err)
		}
	}
}

func JSON(v interface{}) Replyer {
	return serverJSON(v)
}

func EmptyJSON() Replyer {
	return serverJSON(nil)
}

func Err(err error) Replyer {
	return func(w http.ResponseWriter) {
		e, ok := err.(*common.BaseErr)
		if !ok {
			e = common.WrapeInternalError(err)
		}
		switch e.StatusCode {
		case http.StatusInternalServerError:
			http.Error(w, e.Message, http.StatusInternalServerError)
		case http.StatusForbidden:
			http.Error(w, "", http.StatusForbidden)
		case http.StatusNotFound:
			http.Error(w, e.Message, http.StatusNotFound)
		case http.StatusBadRequest:
			http.Error(w, e.Message, http.StatusBadRequest)
		}
	}
}

func BasicAuth() Replyer {
	return func(w http.ResponseWriter) {
		w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}
