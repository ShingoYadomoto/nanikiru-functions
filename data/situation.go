package data

type Fan int

const (
	FanTon Fan = iota + 1
	FanNan
	FanSha
	FanPei
)

type FieldInning uint

const (
	FieldInning1 FieldInning = 1
	FieldInning2 FieldInning = 2
	FieldInning3 FieldInning = 3
	FieldInning4 FieldInning = 4
)

type Field struct {
	Fan     Fan         // 場風
	Inning  FieldInning // 局
	Stack   uint        // 本場
	Deposit uint        // 供託
}

type Situation struct {
	Bonus     PaiStr // ドラ
	Field     Field
	PlayerFan Fan    // 自風
	Other     string // その他（アガリトップなど）
}
