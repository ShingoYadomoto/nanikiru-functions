package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/ShingoYadomoto/nanikiru-functions/data"
	"github.com/aws/aws-lambda-go/events"
	"github.com/gorilla/mux"
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

type Handler struct{}

// for development
func (h Handler) CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Content-Type", "application/json")

		// preflight用に200でいったん返す
		if r.Method == "OPTIONS" {
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

	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write(b)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to write resuponse. err: %s", err.Error())
	}
}

func (h Handler) responseLambda(f func() ([]byte, error)) (*events.APIGatewayProxyResponse, error) {
	headter := map[string]string{"Content-Type": "application/json"}

	b, err := f()
	if err != nil {
		log.Printf("failed to get byte data. err: %s", err.Error())

		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headter,
		}, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    headter,
		Body:       string(b),
	}, nil
}

const excludeIDQueryKey = "exclude_id"

func (h Handler) GetRandomQuestionHandler(rw http.ResponseWriter, r *http.Request) {
	h.response(rw, func() ([]byte, error) {
		excludeIDCSV := r.URL.Query().Get(excludeIDQueryKey)

		bg := BodyGenerator{}
		return bg.GetRandomQuestion(excludeIDCSV)
	})
}

func (h Handler) GetRandomQuestionLambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return h.responseLambda(func() ([]byte, error) {
		excludeIDCSV := request.QueryStringParameters[excludeIDQueryKey]

		bg := BodyGenerator{}
		return bg.GetRandomQuestion(excludeIDCSV)
	})
}

const (
	questionIDQueryKey = "question_id"
	answerQueryKey     = "answer"
)

func (h Handler) GetAnswerHandler(rw http.ResponseWriter, r *http.Request) {
	h.response(rw, func() ([]byte, error) {
		var (
			idStr     = mux.Vars(r)[questionIDQueryKey]
			answerStr = r.URL.Query().Get(answerQueryKey)

			request = AnswerRequest{}
		)

		decodedAnswerStr, err := url.QueryUnescape(answerStr)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(decodedAnswerStr), &request)
		if err != nil {
			return nil, err
		}

		bg := BodyGenerator{}
		return bg.GetAnswerHandler(idStr, request)
	})
}

func (h Handler) GetAnswerLambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return h.responseLambda(func() ([]byte, error) {
		var (
			idStr     = request.QueryStringParameters[questionIDQueryKey]
			answerStr = request.QueryStringParameters[answerQueryKey]
			request   = AnswerRequest{}
		)

		decodedAnswerStr, err := url.QueryUnescape(answerStr)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(decodedAnswerStr), &request)
		if err != nil {
			return nil, err
		}

		bg := BodyGenerator{}
		return bg.GetAnswerHandler(idStr, request)
	})
}
