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

func fake(ctx Tctx) {
	ctx.SetThe(examples.Fake{
		Number: 1,
	})
}

func sum(ctx Tctx) interface{} {
	return ctx.Subject().(examples.Fake).Sum(1)
}

func verify(r Resolver) {
	assert.Equal(r.T(), 2, r.Got())
}
