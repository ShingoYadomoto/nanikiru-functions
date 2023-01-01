package data

import (
	"fmt"
)

type (
	Hand        string
	HandPaiType int
)

const (
	HandPaiTypePai HandPaiType = iota
	HandPaiTypeBonus
	HandPaiTypeFolou
	HandPaiTypeType

	paiCount  = 14
	bonusFlag = "*"
	folouFlag = "_"
)

func (h Hand) reverseHandStr() string {
	runes := []rune(h)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (h Hand) reversePaiList(pl []Pai) []Pai {
	for i, j := 0, len(pl)-1; i < j; i, j = i+1, j-1 {
		pl[i], pl[j] = pl[j], pl[i]
	}
	return pl
}

func (h Hand) Parse() ([]Pai, error) {
	var (
		ret                            = make([]Pai, 0, paiCount)
		currentPaiType                 = PaiTypeZi
		currentIsBonus, currentIsFolou bool
	)
	for _, r := range h.reverseHandStr() {
		s := string(r)

		switch h.checkHandType(s) {
		case HandPaiTypeBonus:
			currentIsBonus = true
		case HandPaiTypeFolou:
			currentIsFolou = true
		case HandPaiTypeType:
			currentPaiType = PaiType(s)
		case HandPaiTypePai:
			p, err := NewHandPai(s, currentPaiType, currentIsFolou, currentIsBonus)
			if err != nil {
				return nil, err
			}

			ret = append(ret, *p)

			currentIsBonus = false
			currentIsFolou = false
		default:
			return nil, fmt.Errorf("unexpected pai: %s", s)
		}
	}

	return h.reversePaiList(ret), nil
}

func (h Hand) checkHandType(s string) HandPaiType {
	if s == bonusFlag {
		return HandPaiTypeBonus
	}
	if s == folouFlag {
		return HandPaiTypeFolou
	}

	if s == string(PaiTypeManzu) ||
		s == string(PaiTypePinzu) ||
		s == string(PaiTypeSozu) {
		return HandPaiTypeType
	}

	return HandPaiTypePai
}
