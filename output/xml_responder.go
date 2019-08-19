package output

import (
	df "../common"
	"encoding/xml"
	"log"
	"net/http"
)

type XmlResponder struct {}

func (jr XmlResponder) Response(writer http.ResponseWriter, postWithComments []df.PostWithComments) {
	resp := df.PostWithCommentsResponse{Posts: postWithComments}
	buf, err := xml.Marshal(resp)
	if err != nil {
		log.Println("get posts failed with error: ", err)
		writer.WriteHeader(500)
		return
	}

	writer.Header().Set("Content-Type", "application/xml")
	_, err = writer.Write(buf)
}

