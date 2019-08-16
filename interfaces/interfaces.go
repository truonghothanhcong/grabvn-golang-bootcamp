package interfaces

import (
	df "../common"
	"net/http"
)

type DataLoader interface {
	GetPost() ([]df.Post, error)
	GetComment() ([]df.Comment, error)
}


type Responder interface {
	Response(http.ResponseWriter, []df.PostWithComments)
}


type Processor interface {
	HandleError(http.ResponseWriter, string, error)
	Processing(http.ResponseWriter, *http.Request)
	CombinePostWithComments([]df.Post, []df.Comment) []df.PostWithComments
}
