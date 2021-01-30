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

func TestOrder_should_be_paid_when_pay(t *testing.T) {
	Given(t, func(ctx Tctx) {
		order := an().order().was().createdAt(time.Now()).as(Submitted).toOrder()

		WhenR(ctx, func() interface{} {
			return order.Pay()
		}).
			Then(func(r Resolver) {
				assert.Equal(r.T(), Paid, r.Got())
			})

	})
}

func TestOrder_should_not_valid(t *testing.T) {
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

type orderFactory struct {
	ordr Order
}

func an() *orderFactory { return &orderFactory{} }

func (o *orderFactory) order() *orderFactory {
	o.ordr = Order{}
	return o
}

func (o *orderFactory) createdAt(t time.Time) *orderFactory {
	o.ordr.created = t
	return o
}

func (o *orderFactory) as(s Status) *orderFactory {
	o.ordr.status = s
	return o
}

func (o *orderFactory) was() *orderFactory { return o }
func (o *orderFactory) toOrder() Order     { return o.ordr }
