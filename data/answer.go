package data

import (
	"fmt"
	"strings"
)

type Answer string

func (ans Answer) Parse() ([]Pai, error) {
	var (
		paiStrList = strings.Split(string(ans), ",")
		ret        = make([]Pai, len(paiStrList))
	)

	for i, paiStr := range paiStrList {
		p, err := NewAnswerPai(paiStr)
		if err != nil {
			return nil, fmt.Errorf("failed to NewAnswerPai(). %s", err)
		}

		ret[i] = *p
	}

	return ret, nil
}
