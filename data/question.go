package data

import (
	"strconv"
)

type (
	QuestionID uint

	Question struct {
		ID     QuestionID
		Hands  Hand
		Answer Answer
		Page   uint
	}
)

func NewQuestionIDFromStr(s string) (QuestionID, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return QuestionID(i), nil
}

func newData() []Question {
	return []Question{
		{Hands: "56m5689p44667s中中中", Answer: "8p", Page: 9},
		{Hands: "45m2344779p23367s", Answer: "4p", Page: 10},
		{Hands: "45*78m1111234467s", Answer: "7m,8m", Page: 11},
		{Hands: "577m23677p245577s", Answer: "2s", Page: 12},
		{Hands: "5*666m567p99p35567s", Answer: "9p", Page: 13},
		{Hands: "3467m5*568p3478s北北", Answer: "北", Page: 14},
		{Hands: "234677m388p23478s", Answer: "7m,3p", Page: 22},
		{Hands: "234677m388p23468s", Answer: "7m", Page: 23},
		{Hands: "788m11379p1389s中中", Answer: "8m", Page: 25},
		{Hands: "788m99p1378s發發中中_中", Answer: "3s", Page: 25},
		{Hands: "12367m24p789p22s88s", Answer: "2p", Page: 31},
	}
}
