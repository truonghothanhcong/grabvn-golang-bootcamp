package processor

import (
	df "../common"
	itf "../interfaces"

	"log"
	"net/http"
)

type Handler struct {
	DataLoader itf.DataLoader
	Responder itf.Responder
}

func (h Handler) HandleError(writer http.ResponseWriter, errorMsg string, err error) {
	log.Println("get posts failed with error: ", err)
	writer.WriteHeader(500)
}

func (h Handler) Processing(writer http.ResponseWriter, request *http.Request) {
	posts, err := h.DataLoader.GetPost()
	if err != nil {
		h.HandleError(writer, "get posts failed with error: ", err)
		return
	}

	comments, err := h.DataLoader.GetComment()
	if err != nil {
		h.HandleError(writer, "get comments failed with error: ", err)
		return
	}

	// Combine and return response
	postWithComments := h.CombinePostWithComments(posts, comments)
	h.Responder.Response(writer, postWithComments)
}

func (h Handler) CombinePostWithComments(posts []df.Post, comments []df.Comment) []df.PostWithComments {
	commentsByPostID := map[int64][]df.Comment{}
	for _, comment := range comments {
		commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
	}

	result := make([]df.PostWithComments, 0, len(posts))
	for _, post := range posts {
		result = append(result, df.PostWithComments{
			ID:       post.ID,
			Title:    post.Title,
			Comments: commentsByPostID[post.ID],
		})
	}

	return result
}
