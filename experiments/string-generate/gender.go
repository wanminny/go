package string_generate

// 注意了!!!!:  go:generate 前面不能有空格
//
//go:generate stringer -type=Gender
type Gender uint8

const (
	Male Gender = iota
	FeMale
	Secret
)

/*
在当前目录执行 go generate 即可！
*/
