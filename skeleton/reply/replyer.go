package reply

import (
	"encoding/json"
	"net/http"

	"github.com/CodiesTeam/codies-server/skeleton/common"
)

// Replyer write result to r
type Replyer func(w http.ResponseWriter)

func JSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

func Error(w http.ResponseWriter, err error) {
	if common.IsForbiddenError(err) {
		http.Error(w, "", http.StatusForbidden)
	} else if common.IsNotFoundError(err) {
		http.Error(w, err.(*common.BaseErr).Msg, http.StatusNotFound)
	} else if common.IsInvalidArgumentError(err) {
		http.Error(w, err.(*common.BaseErr).Msg, http.StatusBadRequest)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func serverJSON(v interface{}) Replyer {
// 	return func(w http.ResponseWriter) {
// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		w.WriteHeader(http.StatusOK)
// 		if err := json.NewEncoder(w).Encode(v); err != nil {
// 			panic(err)
// 		}
// 	}
// }

// func JSON(v interface{}) Replyer {
// 	return serverJSON(v)
// }

// func EmptyJSON() Replyer {
// 	return serverJSON(nil)
// }
