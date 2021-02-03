package examples

type Fake struct {
	Number int
}

func (f Fake) Sum(n int) int {
	return f.Number + n
}
