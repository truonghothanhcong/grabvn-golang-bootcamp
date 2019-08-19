package output

import (
	df "../common"
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponder struct {}

func (jr JsonResponder) Response(writer http.ResponseWriter, postWithComments []df.PostWithComments) {
	resp := df.PostWithCommentsResponse{Posts: postWithComments}
	buf, err := json.Marshal(resp)
	if err != nil {
		log.Println("get posts failed with error: ", err)
		writer.WriteHeader(500)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(buf)
}
