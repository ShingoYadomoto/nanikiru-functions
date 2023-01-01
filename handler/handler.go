package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ShingoYadomoto/nanikiru-functions/data"
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

func (h Handler) convertPaiList(pl []data.Pai) ([]QuestionPai, error) {
	paiList := make([]QuestionPai, len(pl))
	for i, parsedPai := range pl {

		idx, err := parsedPai.Index.ToUint8()
		if err != nil {
			return nil, err
		}

		paiList[i] = QuestionPai{
			Type:    parsedPai.Type,
			Index:   idx,
			IsFolou: parsedPai.IsFolou,
			IsBonus: parsedPai.IsBonus,
		}
	}

	return paiList, nil
}

const excludeIDQueryKey = "exclude_id"

func (h Handler) GetRandomQuestionHandler(rw http.ResponseWriter, r *http.Request) {
	h.response(rw, func() ([]byte, error) {
		var (
			query            = r.URL.Query()
			excludeIDCSV     = query.Get(excludeIDQueryKey)
			excludeIDStrList = strings.Split(excludeIDCSV, ",")
			excludeIDList    = make([]data.QuestionID, 0, len(excludeIDStrList))
		)

		for _, excludeIDStr := range excludeIDStrList {
			if excludeIDStr == "" {
				continue
			}

			id, err := data.NewQuestionIDFromStr(excludeIDStr)
			if err != nil {
				return nil, err
			}

			excludeIDList = append(excludeIDList, id)
		}

		question, err := data.GetQuestioner().GetRandomQuestion(excludeIDList)
		if err != nil {
			return nil, err
		}

		parsedPaiList, err := question.Hands.Parse()
		if err != nil {
			return nil, err
		}

		paiList, err := h.convertPaiList(parsedPaiList)
		if err != nil {
			return nil, err
		}

		j, err := json.Marshal(QuestionResponse{
			ID:      question.ID,
			PaiList: paiList,
			Page:    question.Page,
		})
		if err != nil {
			return nil, err
		}

		return j, nil
	})
}

func (h Handler) PostAnswerHandler(rw http.ResponseWriter, r *http.Request) {
	h.response(rw, func() ([]byte, error) {
		id, err := data.NewQuestionIDFromStr(mux.Vars(r)["question_id"])
		if err != nil {
			return nil, err
		}

		var request AnswerRequest
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return nil, err
		}

		userAnswerPai, err := data.NewPai(request.UserAnswer.Type, request.UserAnswer.Index, request.UserAnswer.IsFolou, request.UserAnswer.IsBonus)
		if err != nil {
			return nil, err
		}

		correctAnswer, err := data.GetQuestioner().GetQuestion(id)
		if err != nil {
			return nil, err
		}

		parsedPaiList, err := correctAnswer.Answer.Parse()
		if err != nil {
			return nil, err
		}

		var isCorrect bool
		for _, correctPai := range parsedPaiList {
			if correctPai.Equal(userAnswerPai) {
				isCorrect = true
				break
			}
		}

		answer, err := h.convertPaiList(parsedPaiList)
		if err != nil {
			return nil, err
		}

		j, err := json.Marshal(AnswerResponse{
			IsCorrect:     isCorrect,
			CorrectAnswer: answer,
		})
		if err != nil {
			return nil, err
		}

		return j, nil
	})
}
