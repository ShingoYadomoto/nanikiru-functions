package data

import (
	"fmt"
	"strings"
)

type PaiStr string

func (ps PaiStr) Parse() ([]Pai, error) {
	if ps == "" {
		return []Pai{}, nil
	}

	var (
		paiStrList = strings.Split(string(ps), ",")
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
