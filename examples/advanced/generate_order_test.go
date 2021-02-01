package advanced

import (
	"time"

	"testy/testy"
)

type generators struct {
	registrar      testy.Registrar
	orderGenerator *orderGenerator
}

type orderGenerator struct {
	ordr Order
}

func a(registrar testy.Registrar) *generators {
	return &generators{
		registrar: registrar,
	}
}
func an(registrar testy.Registrar) *generators { return a(registrar) }

func (o *generators) order() *orderGenerator {
	return &orderGenerator{}
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

func (o *orderGenerator) generate() *Order { return &o.ordr }
