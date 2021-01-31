package advanced

import "time"

type orderGenerator struct {
	ordr Order
}

func an() *orderGenerator { return &orderGenerator{} }

func (o *orderGenerator) order() *orderGenerator {
	o.ordr = Order{}
	return o
}

func (o *orderGenerator) createdAt(t time.Time) *orderGenerator {
	o.ordr.created = t
	return o
}

func (o *orderGenerator) as(s Status) *orderGenerator {
	o.ordr.status = s
	return o
}

func (o *orderGenerator) was() *orderGenerator { return o }
func (o *orderGenerator) toOrder() *Order      { return &o.ordr }
