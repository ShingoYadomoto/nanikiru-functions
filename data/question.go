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
		{Hands: "56m5689p44667s中中中", Answer: "8p", Page: 9, Comment: "ターツオーバー"},
		{Hands: "45m2344779p23367s", Answer: "4p", Page: 10, Comment: "タンヤオを見て9pも悪くないが、5ブロックあるのでメンツ固定"},
		{Hands: "45*78m1111234467s", Answer: "7m,8m", Page: 11, Comment: "ターツオーバーなので2度受け解消"},
		{Hands: "577m23677p245577s", Answer: "2s", Page: 12, Comment: "索子2ターツとすると2sが余剰牌"},
		{Hands: "5*666m567p99p35567s", Answer: "9p", Page: 13, Comment: "マンズをヘッド候補に、打点・形込みで打9p"},
		{Hands: "3467m5*568p3478s北北", Answer: "北", Page: 14, Comment: "打8pが受け入れ最大だが次順迷うので、順目込みで北対子落としで5ブロックに", Situation: &Situation{Other: "2順目"}},
		{Hands: "234677m388p23478s", Answer: "7m,3p", Page: 22, Comment: "7m切りで平和確保・タンヤオ/三色をみる or 打3pで素直に完全一向聴"},
		{Hands: "234677m388p23468s", Answer: "7m", Page: 23, Comment: "カンチャンが不満。リャンメンを固定しつつ、3pくっつきをみるバランス"},
		{Hands: "788m11379p1389s中中", Answer: "8m", Page: 25, Comment: "一向聴ピーク理論。向聴数に拘らず、チャンタ三色を見て6ブロックに"},
		{Hands: "788m99p1378s發發中中_中", Answer: "3s", Page: 25, Comment: "二向聴の場合、5ブロックに絞るのが基本"},
		{Hands: "12367m24789p22s88s", Answer: "2p", Page: 30, Comment: "2ヘッドを残してリャンメン変化枚数優先。5p4枚 vs 3s7s8枚"},

		{Hands: "2346m23346p45678s", Answer: "6m", Page: 100, Comment: "中ぶくれ+1。6pは5pの受け入れがある分6mより重要"},
		{Hands: "234m233468p45678s", Answer: "8p", Page: 101, Comment: "3p:10種33枚, 8p:6種20枚 好形優先で打8p（即聴牌効率だけなら打3p）"},
		{Hands: "789m13445p45678s北", Answer: "北", Page: 101, Comment: "ヘッドレスでは中ぶくれ+1は重要"},
		{Hands: "3456m6778p4456s北北", Answer: "4s", Page: 102, Comment: "複合ターツが混在したら大体亜リャンメンを嫌う"},
		{Hands: "6678m34p22344568s", Answer: "6m", Page: 104, Comment: "8sは7s8s受け入れの重要牌で打6mが受け入れ最大 + 打点面も有利"},
		{Hands: "34568m6678p34888s", Answer: "6p", Page: 105, Comment: "8mは7m8m受け入れの重要牌で打6pが受け入れ最大"},
		{Hands: "33678m3344p23478s", Answer: "4p", Page: 112, Comment: "平和系なので並び対子ほぐし"},
		{Hands: "33678m3344p23468s", Answer: "8s", Page: 112, Comment: "25p引きの高目タンピン一盃口まで見て打8s"},
		{Hands: "33678m8899p23468s", Answer: "9p", Page: 113, Comment: "端の並び対子は価値低。3枚構成に構える"},
		{Hands: "3368m3344p456667s", Answer: "8m", Page: 113, Comment: "対子手とタンピン一盃口の両天秤"},
		{Hands: "7799m12345*6p5578s", Answer: "9m,7m", Page: 114, Comment: "飛び対子から1枚切って3枚構成に（打9mが若干有利。打9m: リャンメン・タンヤオ変化 vs 打7m: 最終系シャンポンリーチ・索子のの日対応）"},
		{Hands: "6688m123789p4578s", Answer: "8m", Page: 115, Comment: "飛び対子から1枚切って平和一向聴に。789三色を見て打8m"},
		{Hands: "237799m123567p68s", Answer: "8s", Page: 115, Comment: "568m引きでの変化を見てカンチャン外し"},
		{Hands: "678m224466p12378s", Answer: "4p", Page: 116, Comment: "飛び対子は中を切れ"},
		{Hands: "678m224466p12399s", Answer: "9s", Page: 116, Comment: "どの対子を切っても受け入れ変わらないので内寄せ+一盃口狙い"},
		{Hands: "678m224466p23477s", Answer: "7s,4p", Page: 117, Comment: "打7sが期待値最大だが、鳴きとリャンメン変化を見て打4pもあり"},
		{Hands: "678m455679p2399s北", Answer: "北", Page: 118, Comment: "リャンメンカンチャン。打9pは4枚ロス"},
		{Hands: "67m11456p2356778s", Answer: "2s", Page: 118, Comment: "リャンメンカンチャンの定番。打2s"},
		{Hands: "678m455679p22388s", Answer: "9p", Page: 119, Comment: "打9p打2sで受け入れは同じなので、タンヤオ優先"},
		{Hands: "3556778m45p12388s", Answer: "3m,5m", Page: 120, Comment: "打5m: 一盃口は消えるが平和確定（最低打点の確保）, 打3m: 一盃口が残り最高形を見れるが、リーチドラ1の最安値になる可能性あり（最高打点を見る）"},
		{Hands: "348m34458p35*7788s", Answer: "8m", Page: 122, Comment: "ツモ6pで8pが生きるので打8m"},
		{Hands: "348m44568p35*7788s", Answer: "8m", Page: 122, Comment: "ツモ3pで8pが生きるので打8m"},
		{Hands: "789m122346p34889s", Answer: "9s", Page: 123, Comment: "端のリャンメンカンチャン。打9sで目一杯"},
		{Hands: "3488m22234p12378s", Answer: "8m", Page: 130, Comment: "エントツ形はなるべく残す。一手先で完全一向聴になると5pの受け入れ分打8mが有利"},
		{Hands: "3488m44456p22278s", Answer: "7s,8s", Page: 131, Comment: "対子落としと受け入れがほとんど変わらないので、タンヤオ・三暗刻・一盃口を見て6-9s払い"},
		{Hands: "78m4556778p11345s", Answer: "7p", Page: 132, Comment: "一盃口より形優先。赤引きでの聴牌を考慮して打7p"},
		{Hands: "78m5667889p11345s", Answer: "9p", Page: 132, Comment: "打689pの受け入れ枚数は変わらないので、一盃口見て打9p"},
		{Hands: "67m34599p2234445s", Answer: "2s", Page: 133, Comment: "打2s: 受け入れ22枚で最大, 打4s: 受け入れ19枚だが一盃口が残る 打5s: 受け入れ20枚だが平和になりづらく損"},
		{Hands: "67m34599p1123334s", Answer: "3s", Page: 133, Comment: "受け入れ最大かつ一盃口が残る"},
		{Hands: "677m222334p115*67s", Answer: "7m", Page: 134, Comment: "打2pは平和・一盃口目があるが、受け入れ5枚差をひっくり返すほどではない。特に今回は赤5p引きでの聴牌を逃すことにもなる"},
		{Hands: "677m111223p115*67s", Answer: "7m", Page: 134, Comment: "受け入れ最大かつ、最終形1-4p・1sを見る。打1pで一盃口は枚数と打点上昇の折り合いがつかない"},
		{Hands: "789m33445667p557s", Answer: "7s", Page: 136, Comment: "多面待ちになりやすく一盃口も見れる打7s。打6pは受け入れ最大だが、愚形聴牌率が高すぎる。"},
		{Hands: "789m11223445p557s", Answer: "7s", Page: 136, Comment: "多面待ちになりやすく一盃口も見れる打7s。打4pは受け入れ最大だが、愚形聴牌率が高すぎる"},
		{Hands: "789m1122345*5p557s", Answer: "5p", Page: 137, Comment: "打5pで横伸びに柔軟な形に構え、一盃口目も残す。"},
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
