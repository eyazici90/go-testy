package v3

import (
	"testing"
	. "testy/testy"

	"testy/examples"

	"github.com/go-playground/assert"
)

func TestSum_v3(t *testing.T) {
	Given(t, func(ctx Tctx) {
		fixture := a().fake().has().number(2)

		WhenR(ctx, func() interface{} {
			return fixture.fk.Sum(4)
		}).
			Then(func(r Resolver) {
				assert.Equal(r.T(), 6, r.Got())
			})
	})
}

type fakeFixture struct {
	fk examples.Fake
}

func a() *fakeFixture { return &fakeFixture{} }

func (f *fakeFixture) fake() *fakeFixture {
	f.fk = examples.Fake{}
	return f
}

func (f *fakeFixture) has() *fakeFixture { return f }

func (f *fakeFixture) number(n int) *fakeFixture {
	f.fk.Number = n
	return f
}
