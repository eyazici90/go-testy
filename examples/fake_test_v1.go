package examples

import (
	"testing"
	. "testy/pkg/testy"

	"github.com/go-playground/assert"
)

func TestSum_v1(t *testing.T) {
	Given(t, func(ctx *TestContext) {
		ctx.Use().A(Fake{
			number: 1,
		})
	}).WhenR(func(ctx *TestContext) interface{} {
		tF := ctx.The(Fake{})
		f := tF.(Fake)

		return f.Sum(1)
	}).Then(func(ctx *TestContext) {
		assert.Equal(ctx.T(), 2, ctx.Got())
	})
}
