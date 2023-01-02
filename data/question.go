package data

import (
	"strconv"
)

type (
	QuestionID uint

	Question struct {
		ID        QuestionID
		Hands     Hand
		Answer    PaiStr
		Page      uint
		Situation *Situation
		Comment   string
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
		{Hands: "233456778p33488s", Answer: "3s", Page: 138, Comment: "ウイング9枚系 3s切りで受け入れ最大・必ず平和になる"},
		{Hands: "122345667p33488s", Answer: "3s", Page: 139, Comment: "ウイング9枚系 端にくっつくと受け入れに差があるのでウイングとは呼ばない。ツモ3pによる平和テンパイを見逃しやすいので注意"},
		{Hands: "1122345667p3488s", Answer: "1p", Page: 139, Comment: "ウイング9枚系 全ての受け入れで平和テンパイ、ツモ次第でタンヤオ以降も可能"},
		{Hands: "778m23345667p123s", Answer: "8m", Page: 140, Situation: &Situation{Other: "4巡目"}, Comment: "ウイング8枚系 高確率で平和になりつつ、高目一盃口への変化が多い"},
		{Hands: "778m23345668p123s", Answer: "7m", Page: 141, Comment: "ウイング8枚系 ウイングが崩れたことで平和の可能性も下がり、8mより7m有利"},
		{Hands: "44677889m34457p2s", Answer: "2s", Page: 146, Comment: "有効牌の種類と枚数を数える 4p: 11種39枚・好形率31% / 2s: 9種29枚・好形率62%"},
		{Hands: "34788m222334p678s", Answer: "7m", Page: 152, Comment: "メンツ抜き 234pを抜くとリャンメン・リャンメンの比較。25p受けもある複合系を残す"},
		{Hands: "34788m222344p678s", Answer: "4p", Page: 152, Comment: "メンツ抜き 234pを抜くとリャンメン・カンチャンの比較"},
		{Hands: "34688m222344567p", Answer: "6m", Page: 153, Comment: "メンツ抜き 234p・567pを抜き出すとカンチャン同士を比較し、一盃口の可能性を残す"},
		{Hands: "788m33444556p123s", Answer: "7m", Page: 154, Comment: "メンツ抜き 456pを抜き出すと高目タンピン一盃口のテンパイが見える"},
		{Hands: "77m3456789p22s789s", Answer: "7m", Page: 168, Comment: "浮かせ打ち 7mを浮かせて一通・三色両天秤に構える"},
		{Hands: "77m11345678p3478s", Answer: "1p", Page: 168, Comment: "浮かせ打ち 1pを浮かせてソーズを引けばリーチ・一通が見えたらソーズターツ外し"},
		{Hands: "77m22456789p1134s", Answer: "1s", Page: 169, Comment: "浮かせ打ち ヘッドが複数の場合は浮かすさず連打前提"},
		{Hands: "3478m34566p33678s", Answer: "6p", Page: 170, Comment: "目先の受け入れと2次変化 後に3面チャンになればリャンメンを切る。3s切り三色狙いは非効率"},
		{Hands: "34m34566p2355678s", Answer: "2s", Page: 172, Comment: "目先の受け入れと2次変化 三色は無理には狙わないが自然に狙える時は狙わないと損"},
		{Hands: "67m45677p2355678s", Answer: "2s", Page: 173, Comment: "目先の受け入れと2次変化 タンヤオのなりやすさと、自然に678三色を狙う"},
		{Hands: "67m56788p4456677s", Answer: "8p", Page: 173, Comment: "目先の受け入れと2次変化 三色両天秤で4枚ロスをカバー"},
		{Hands: "334m678p35567778s", Answer: "3m", Page: 174, Comment: "迷った時はリャンメン固定で80点の打牌を量産 3mのロスも2m出上がり率向上でカバー"},
		{Hands: "788m234466p22678s", Answer: "8m", Page: 175, Situation: &Situation{Bonus: "2s"}, Comment: "迷った時はリャンメン固定 明確な差がない場合は面前効率重視"},
		{Hands: "33445*6m5678p4456s", Answer: "4s", Page: 176, Situation: &Situation{Bonus: "5s"}, Comment: "迷った時は亜リャンメン切り"},
		{Hands: "466m34799p234668s", Answer: "9p", Page: 178, Comment: "とりあえず内寄せ"},
		{Hands: "579m34678p233444s", Answer: "9m", Page: 179, Comment: "とりあえず内寄せ"},
		{Hands: "23345667p223789s", Answer: "3s", Page: 180, Situation: &Situation{Bonus: "2s"}, Comment: "ウィング8枚系 ドラ固定 ピンズ中ぶくれ2つを活かす"},
		{Hands: "45688m222344p677s", Answer: "7s", Page: 184, Situation: &Situation{Other: "アガリトップ"}, Comment: "チー2倍速、ポン4倍速 受け入れ最大がそのまま正解"},
		{Hands: "45688m222344p788s", Answer: "4p", Page: 185, Situation: &Situation{Other: "アガリトップ"}, Comment: "チー2倍速、ポン4倍速 9sの受け入れはカウントしない"},
		{Hands: "3445*88m5677p3s中中_中", Answer: "3s", Page: 186, Situation: &Situation{Bonus: "8m"}, Comment: "鳴き・面前で変わる複合系の価値 鳴き効率を最大化"},
		{Hands: "788m2245*6p667s中中_中", Answer: "6s", Page: 188, Comment: "鳴き効率（好形・3ヘッド） ほとんどのケースでリャンメン固定有利"},
		{Hands: "688m2245*6p35*5s7_68p", Answer: "3s", Page: 190, Comment: "鳴き効率（愚形・3ヘッド） 同枚数ならポン効率の対子固定有利"},
		{Hands: "5*678m224567p35*55s", Answer: "8m", Page: 193, Situation: &Situation{Bonus: "3s"}, Comment: "鳴いても満貫は強い 受け入れ枚数-18%を打点2倍でカバー"},
		{Hands: "2388m1234567p中中_中", Answer: "1p", Page: 194, Situation: &Situation{Other: "上家から赤5pが捨てられた時に備えるには？"}, Comment: "喰い替え 2pを端に寄せて58pチー打2p"},
		{Hands: "2388m1234456p中中_中", Answer: "1p", Page: 194, Situation: &Situation{Other: "上家から赤5pが捨てられた時に備えるには？"}, Comment: "喰い替え 2pを端に寄せてカン5pチー打2p"},
		{Hands: "2388m2334456p中中_中", Answer: "3p", Page: 195, Situation: &Situation{Other: "上家から赤5pが捨てられた時に備えるには？"}, Comment: "喰い替え 2pを孤立させてカン5pチー打2p"},
		{Hands: "2388m2344556p中中_中", Answer: "5p", Page: 195, Situation: &Situation{Other: "上家から赤5pが捨てられた時に備えるには？"}, Comment: "喰い替え 5pを一枚にしておいてカン5pチー打2p"},
		{Hands: "2388m2345567p中中_中", Answer: "5p", Page: 195, Situation: &Situation{Other: "上家から赤5pが捨てられた時に備えるには？"}, Comment: "喰い替え 5pを一枚にしておいて58pチー打2p"},
	}
}
