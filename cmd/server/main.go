package main

import (
	"log"
	"net/http"

	"github.com/ShingoYadomoto/nanikiru-functions/handler"
	"github.com/gorilla/mux"
)

func main() {
	h := handler.Handler{}

	r := mux.NewRouter()
	r.HandleFunc("/questions", h.CORSMiddleware(h.GetRandomQuestionHandler)).Methods("GET", "OPTIONS")
	r.HandleFunc(`/questions/{question_id:\d+}`, h.CORSMiddleware(h.GetAnswerHandler)).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8888", r))
}
