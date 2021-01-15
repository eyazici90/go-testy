package examples

import (
	"testing"

	"testy/examples"
	. "testy/testy"

	"github.com/go-playground/assert"
)

func TestSum_v1(t *testing.T) {
	Given(t, func(ctx *TestContext) {
		ctx.Use().A(examples.Fake{
			Number: 1,
		})
		ctx.SetThe(examples.Fake{
			Number: 1,
		})
	}).WhenR(func(ctx *TestContext) interface{} {
		return ctx.Subject().(examples.Fake).Sum(1)

	}).Then(func(ctx *TestContext) {
		assert.Equal(ctx.T(), 2, ctx.Got())
	})
}
