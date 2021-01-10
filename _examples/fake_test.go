package _examples

import (
	"testing"
	. "testy/pkg/testy"

	"github.com/go-playground/assert"
)

type Fake struct {
	number int
}

func Sum(f1, f2 Fake) int {
	return f1.number + f1.number
}

func TestCalculate(t *testing.T) {
	Given(t, fake).WhenR(func(ctx *TestContext) interface{} {
		f := ctx.The(Fake{})
		f1, f2 := f.(Fake), f.(Fake)

		return Sum(f1, f2)
	}).Then(verify)
}

func fake(ctx *TestContext) {
	ctx.Use().A(Fake{
		number: 1,
	})
}

func verify(t *testing.T, ctx *TestContext) {
	assert.Equal(t, 2, ctx.Got())
}
