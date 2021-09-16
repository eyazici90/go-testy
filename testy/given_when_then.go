package testy

import "testing"

type (
	Do  func(Tctx)
	DoR func(Tctx) interface{}
)

type (
	Tctx interface {
		Registrar
		Resolver
	}
	Registrar interface {
		SetThe(i interface{})
		BeforeEach(fn func())
		AfterEach(fn func())
	}
	Resolver interface {
		Got() interface{}
		Err() error
		Subject() interface{}
		T() *testing.T
	}
)

type (
	GivenResult struct {
		ctx *TestContext
	}
	WhenResult struct {
		ctx *TestContext
	}
)

func (g *GivenResult) When(fn Do) *WhenResult {
	copy := g.ctx.copy()
	copy.beforeEach()
	fn(copy)
	copy.afterEach()

	return &WhenResult{
		ctx: copy,
	}
}

func (g *GivenResult) WhenR(fn DoR) *WhenResult {
	copy := g.ctx.copy()
	copy.beforeEach()
	copy.got = fn(copy)
	copy.afterEach()

	return &WhenResult{
		ctx: copy,
	}
}

func (w *WhenResult) Then(fn func(Resolver)) {
	fn(w.ctx)
}

func Given(t *testing.T, fn Do) *GivenResult {
	ctx := NewTestContext(t)
	fn(ctx)
	return &GivenResult{
		ctx: ctx,
	}
}

func When(ctx Tctx, fn func()) *WhenResult {
	copy := ctx.(*TestContext).copy()
	copy.beforeEach()
	fn()
	copy.got = copy.Subject()
	copy.afterEach()

	return &WhenResult{
		copy,
	}
}

func WhenR(ctx Tctx, fn func() interface{}) *WhenResult {
	copy := ctx.(*TestContext).copy()
	copy.beforeEach()
	copy.got = fn()
	copy.afterEach()

	return &WhenResult{
		copy,
	}
}

func WhenRErr(ctx Tctx, fn func() (interface{}, error)) *WhenResult {
	copy := ctx.(*TestContext).copy()
	copy.beforeEach()
	got, err := fn()
	copy.got = got
	copy.err = err
	copy.afterEach()

	return &WhenResult{
		copy,
	}
}
