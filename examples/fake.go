package examples

type Fake struct {
	number int
}

func (f Fake) Sum(n int) int {
	return f.number + n
}
