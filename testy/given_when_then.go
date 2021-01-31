package testy

import "testing"

type (
	Do  func(Tctx)
	DoR func(Tctx) interface{}
)

func Given(t *testing.T, fn Do) GivenResult {
	ctx := NewTestContext(t)
	fn(ctx)
	return GivenResult{
		ctx: ctx,
	}
}

func (g GivenResult) When(fn Do) WhenResult {
	copy := g.ctx.copy()
	fn(copy)
	return WhenResult{
		ctx: copy,
	}
}

func (g GivenResult) WhenR(fn DoR) WhenResult {
	copy := g.ctx.copy()
	got := fn(copy)
	copy.got = got

	return WhenResult{
		ctx: copy,
	}
}

func (w WhenResult) Then(fn func(Resolver)) {
	fn(w.ctx)
}

func When(context Tctx, fn func()) WhenResult {
	fn()
	copy := context.(*TestContext).copy()
	copy.got = copy.Subject()
	return WhenResult{
		copy,
	}
}

func WhenR(context Tctx, fn func() interface{}) WhenResult {
	got := fn()
	ctx := context.(*TestContext)
	copy := ctx.copy()
	copy.got = got

	return WhenResult{
		copy,
	}
}

func WhenRErr(context Tctx, fn func() (interface{}, error)) WhenResult {
	got, err := fn()
	ctx := context.(*TestContext)
	ctx.got = got
	ctx.err = err

	return WhenResult{
		ctx.copy(),
	}
}

func Then(ctx *TestContext, fn func()) {
	fn()
}
