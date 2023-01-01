package handler

import (
	"encoding/json"
	"strings"

	"github.com/ShingoYadomoto/nanikiru-functions/data"
)

type BodyGenerator struct{}

func (bg BodyGenerator) convertPaiList(pl []data.Pai) ([]QuestionPai, error) {
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

func (bg BodyGenerator) GetRandomQuestion(excludeIDCSV string) ([]byte, error) {
	var (
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

	paiList, err := bg.convertPaiList(parsedPaiList)
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
}

func (bg BodyGenerator) GetAnswerHandler(idStr string, request AnswerRequest) ([]byte, error) {
	id, err := data.NewQuestionIDFromStr(idStr)
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

	answer, err := bg.convertPaiList(parsedPaiList)
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
}
