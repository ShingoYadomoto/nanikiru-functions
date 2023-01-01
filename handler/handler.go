package handler

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ShingoYadomoto/nanikiru-functions/data"
	"github.com/aws/aws-lambda-go/events"
)

type (
	QuestionPai struct {
		Type    data.PaiType `json:"type"`
		Index   uint8        `json:"index"`
		IsFolou bool         `json:"isFolou"`
		IsBonus bool         `json:"isBonus"`
	}

	QuestionResponse struct {
		ID        data.QuestionID `json:"id"`
		PaiList   []QuestionPai   `json:"paiList"`
		Situation string          `json:"situation"`
		Page      uint            `json:"page"`
	}

	AnswerRequest struct {
		UserAnswer QuestionPai `json:"userAnswer"`
	}
	AnswerResponse struct {
		IsCorrect     bool          `json:"isCorrect"`
		CorrectAnswer []QuestionPai `json:"correctAnswer"`
	}
)

var commonHeader = map[string]string{
	"Content-Type":                     "application/json",
	"Access-Control-Allow-Origin":      os.Getenv("ALLOW_ORIGIN"),
	"Access-Control-Allow-Headers":     "Content-Type",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Methods":     "GET,OPTIONS",
}

type Handler struct{}

func (h Handler) CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for k, v := range commonHeader {
			w.Header().Set(k, v)
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func (h Handler) response(rw http.ResponseWriter, f func() ([]byte, error)) {
	b, err := f()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to get byte data. err: %s", err.Error())
	}

	_, err = rw.Write(b)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to write resuponse. err: %s", err.Error())
	}
}

func (h Handler) responseLambda(httpMethod string, f func() ([]byte, error)) (*events.APIGatewayProxyResponse, error) {
	if httpMethod == http.MethodOptions {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers:    commonHeader,
		}, nil
	}

	b, err := f()
	if err != nil {
		log.Printf("failed to get byte data. err: %s", err.Error())

		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    commonHeader,
		}, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    commonHeader,
		Body:       string(b),
	}, nil
}

const excludeIDQueryKey = "exclude_id"

func (h Handler) GetRandomQuestionHandler(rw http.ResponseWriter, r *http.Request) {
	h.response(rw, func() ([]byte, error) {
		excludeIDStr := r.URL.Query().Get(excludeIDQueryKey)

		bg := BodyGenerator{}
		return bg.GetRandomQuestion(excludeIDStr)
	})
}

func (h Handler) GetRandomQuestionLambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return h.responseLambda(request.HTTPMethod, func() ([]byte, error) {
		excludeIDStr := request.QueryStringParameters[excludeIDQueryKey]

		bg := BodyGenerator{}
		return bg.GetRandomQuestion(excludeIDStr)
	})
}

const (
	questionIDQueryKey = "question_id"
	answerQueryKey     = "answer"
)

func (h Handler) GetAnswerHandler(rw http.ResponseWriter, r *http.Request) {
	h.response(rw, func() ([]byte, error) {
		var (
			idStr     = r.URL.Query().Get(questionIDQueryKey)
			answerStr = r.URL.Query().Get(answerQueryKey)
		)

		bg := BodyGenerator{}
		return bg.GetAnswerHandler(idStr, answerStr)
	})
}

func (h Handler) GetAnswerLambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return h.responseLambda(request.HTTPMethod, func() ([]byte, error) {
		var (
			idStr     = request.QueryStringParameters[questionIDQueryKey]
			answerStr = request.QueryStringParameters[answerQueryKey]
		)

		bg := BodyGenerator{}
		return bg.GetAnswerHandler(idStr, answerStr)
	})
}
