package v2

import (
	"testing"

	"testy/examples"
	. "testy/testy"

	"github.com/go-playground/assert"
)

func TestSum_v2(t *testing.T) {
	Given(t, fake).WhenR(sum).Then(verify)
}

func fake(ctx *TestContext) {
	ctx.SetThe(examples.Fake{
		Number: 1,
	})
}

func sum(ctx *TestContext) interface{} {
	return ctx.Subject().(examples.Fake).Sum(1)
}

func verify(ctx *TestContext) {
	assert.Equal(ctx.T(), 2, ctx.Got())
}
