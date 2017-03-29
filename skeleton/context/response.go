package context

import (
	"net/http"

	"github.com/CodiesTeam/codies-server/skeleton/reply"
)

type Response struct {
	http.ResponseWriter
	ReplyFunc reply.Replyer
}

func NewResponse(w http.ResponseWriter) Response {
	return Response{
		ResponseWriter: w,
	}
}

// func (rep *Response) Reply() {
// 	rep.ReplyFunc(rep.ResponseWriter)
// }
