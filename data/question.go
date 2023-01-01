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
		{Hands: "12367m24789p22s88s", Answer: "2p", Page: 31},

		{Hands: "2346m23346p45678s", Answer: "6m", Page: 100},
		{Hands: "234m233468p45678s", Answer: "8p", Page: 101},
		{Hands: "789m13445p45678s北", Answer: "北", Page: 101},
		{Hands: "3456m6778p4456s北北", Answer: "4s", Page: 102},
		{Hands: "6678m34p22344568s", Answer: "6m", Page: 104},
		{Hands: "34568m6678p34888s", Answer: "6p", Page: 105},
		{Hands: "33678m3344p23478s", Answer: "4p", Page: 112},
		{Hands: "33678m3344p23468s", Answer: "8s", Page: 112},
		{Hands: "33678m8899p23468s", Answer: "9p", Page: 113},
		{Hands: "3368m3344p456667s", Answer: "8m", Page: 113},
		{Hands: "7799m12345*6p5578s", Answer: "9m,7m", Page: 115},
		{Hands: "6688m123789p4578s", Answer: "8m", Page: 115},
		{Hands: "237799m123567p68s", Answer: "8s", Page: 115},
		{Hands: "678m224466p12378s", Answer: "4p", Page: 116},
		{Hands: "678m224466p12399s", Answer: "9s", Page: 116},
		{Hands: "678m224466p23477s", Answer: "7s,4p", Page: 117},
		{Hands: "678m455679p2399s北", Answer: "北", Page: 118},
		{Hands: "67m11456p2356778s", Answer: "2s", Page: 118},
		{Hands: "678m455679p22388s", Answer: "9p", Page: 119},
		{Hands: "3556778m45p12388s", Answer: "3m", Page: 120},
		{Hands: "348m34458p35*7788s", Answer: "8m", Page: 122},
		{Hands: "348m44568p35*7788s", Answer: "8m", Page: 122},
		{Hands: "789m122346p34889s", Answer: "9s", Page: 123},
		{Hands: "3488m22234p12378s", Answer: "8m", Page: 130},
		{Hands: "3488m44456p22278s", Answer: "7s,8s", Page: 131},
		{Hands: "78m4556778p11345s", Answer: "7p", Page: 132},
		{Hands: "78m5667889p11345s", Answer: "9p", Page: 132},
		{Hands: "67m34599p2234445s", Answer: "2s", Page: 133},
		{Hands: "67m34599p1123334s", Answer: "3s", Page: 133},
		{Hands: "677m222334p115*67s", Answer: "7m", Page: 134},
		{Hands: "677m111223p115*67s", Answer: "7m", Page: 134},
		{Hands: "789m33445667p557s", Answer: "7s", Page: 136},
		{Hands: "789m11223445p557s", Answer: "7s", Page: 137},
		{Hands: "789m1122345*5p557s", Answer: "5p", Page: 137},
		{Hands: "233456778p33488s", Answer: "3s", Page: 138},
		{Hands: "122345667p33488s", Answer: "3s", Page: 139},
		{Hands: "1122345667p3488s", Answer: "1p", Page: 139},
		{Hands: "778m23345667p123s", Answer: "8m", Page: 140},
		{Hands: "778m23345668p123s", Answer: "7m", Page: 141},
		{Hands: "44677889m34457p2s", Answer: "2s", Page: 146},
		{Hands: "34788m222334p678s", Answer: "7m", Page: 152},
		{Hands: "34788m222344p678s", Answer: "4p", Page: 152},
		{Hands: "34688m222344567p", Answer: "6m", Page: 153},
		{Hands: "788m33444556p123s", Answer: "7m", Page: 154},
	}
}
