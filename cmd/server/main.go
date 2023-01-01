package main

import (
	"log"
	"net/http"

	"github.com/ShingoYadomoto/nanikiru-functions/handler"
)

func main() {
	h := handler.Handler{}

	http.HandleFunc("/question", h.CORSMiddleware(h.GetRandomQuestionHandler))
	http.HandleFunc(`/answer`, h.CORSMiddleware(h.GetAnswerHandler))

	log.Fatal(http.ListenAndServe(":8888", nil))
}
