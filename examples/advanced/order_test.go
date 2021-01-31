package advanced

import (
	"testing"
	"time"

	. "testy/testy"

	"github.com/stretchr/testify/assert"
)

func TestOrder_should_be_shipped_when_ship(t *testing.T) {
	Given(t, func(ctx Tctx) {
		order := an().order().was().createdAt(time.Now()).as(Submitted).toOrder()

		When(ctx, func() {
			order.Ship()
		}).
			Then(func(r Resolver) {
				assert.Equal(r.T(), Shipped, order.status)
			})

	})
}

func TestOrder_Specs(t *testing.T) {
	GivenOrder(t, func(ctx Tctx) {

		WhenR(ctx, func() interface{} {
			return ctx.Subject().(*Order).Pay()
		}).
			Then(func(r Resolver) {
				assert.Equal(r.T(), Paid, r.Got())
			})

		When(ctx, func() {
			ctx.Subject().(*Order).Ship()
		}).
			Then(func(r Resolver) {
				assert.Equal(r.T(), Shipped, r.Got().(*Order).status)
			})

	})
}

func TestOrder_should_be_not_valid(t *testing.T) {
	Given(t, func(ctx Tctx) {
		order := an().order().was().createdAt(time.Now()).as(Submitted).toOrder()

		WhenRErr(ctx, func() (interface{}, error) {
			return order.Valid()
		}).
			Then(func(r Resolver) {
				assert.False(r.T(), r.Got().(bool))
				assert.Error(r.T(), r.Err())
			})

	})
}

func GivenOrder(t *testing.T, fn func(ctx Tctx)) {
	Given(t, func(ctx Tctx) {
		ctx.SetThe(
			an().order().was().createdAt(time.Now()).as(Submitted).toOrder(),
		)
		fn(ctx)
	})
}
