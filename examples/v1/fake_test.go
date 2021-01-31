package examples

import (
	"testing"

	"testy/examples"
	. "testy/testy"

	"github.com/go-playground/assert"
)

func TestSum_v1(t *testing.T) {
	Given(t, func(ctx Tctx) {
		ctx.SetThe(examples.Fake{
			Number: 1,
		})
	}).WhenR(func(ctx Tctx) interface{} {
		return ctx.Subject().(examples.Fake).Sum(1)

	}).Then(func(r Resolver) {
		assert.Equal(r.T(), 2, r.Got())
	})
}
