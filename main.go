package main

import (
	in "./input"
	itf "./interfaces"
	out "./output"
	pc "./processor"
	"log"
	"net/http"
)


func newProcessor(Loader itf.DataLoader, Responder itf.Responder) itf.Processor {
	return &pc.Handler{
		DataLoader: Loader,
		Responder: Responder,
	}
}


//TODO: how to separate API logic, business logic and response format logic
func main() {

	processor := newProcessor(&in.ServiceLoader{}, &out.XmlResponder{})

	http.HandleFunc("/postWithComments", func(writer http.ResponseWriter, request *http.Request) {
		processor.Processing(writer, request)
	})

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
