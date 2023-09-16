package structs

type BigData struct {
	a      int
	b      int
	c      string
	secret float32
	d      bool
}

var a = BigData{1, 2, "secret", 1, true}
var b = BigData{1, 2, "secret", 1, true}

func Compare() bool {
	return a == b
}
