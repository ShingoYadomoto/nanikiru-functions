package data

import (
	"errors"
	"math/rand"
	"time"
)

type Questioner struct {
	data map[QuestionID]Question
}

func newQuestioner() *Questioner {
	dl := newData()

	m := map[QuestionID]Question{}
	for i, d := range dl {
		id := QuestionID(i + 1)
		d.ID = id

		m[id] = d
	}
	return &Questioner{data: m}
}

var nanikiruQuestioner *Questioner

func init() {
	nanikiruQuestioner = newQuestioner()
	rand.Seed(time.Now().UnixNano())
}

func GetQuestioner() *Questioner {
	return nanikiruQuestioner
}

var ErrNotFound = errors.New("err: Not Found")

func (q *Questioner) GetRandomQuestion(excludeIDList []QuestionID) (*Question, error) {
	var (
		max          = len(q.data)
		excludeIDMap = map[QuestionID]struct{}{}
	)

	if max == len(excludeIDList) {
		return nil, ErrNotFound
	}

	for _, excludeID := range excludeIDList {
		excludeIDMap[excludeID] = struct{}{}
	}

	failCount := 0
	for {
		if failCount == max {
			return nil, errors.New("over max fail count")
		}

		randomID := QuestionID(rand.Intn(max)) + 1
		if _, ng := excludeIDMap[randomID]; ng {
			failCount++
			continue
		}

		ret := q.data[randomID]

		return &ret, nil
	}
}

func (q *Questioner) GetQuestion(id QuestionID) (*Question, error) {
	if d, ok := q.data[id]; ok {
		return &d, nil
	}

	return nil, ErrNotFound
}
