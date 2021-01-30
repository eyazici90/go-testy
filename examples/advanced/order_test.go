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

type orderFixture struct {
	ordr Order
}

func an() *orderFixture { return &orderFixture{} }

func (o *orderFixture) order() *orderFixture {
	o.ordr = Order{}
	return o
}

func (o *orderFixture) createdAt(t time.Time) *orderFixture {
	o.ordr.created = t
	return o
}

func (o *orderFixture) as(s Status) *orderFixture {
	o.ordr.status = s
	return o
}

func (o *orderFixture) was() *orderFixture { return o }
func (o *orderFixture) toOrder() Order     { return o.ordr }
