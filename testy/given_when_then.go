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
	fn(g.ctx)
	return WhenResult{
		ctx: g.ctx.copy(),
	}
}

func (g GivenResult) WhenR(fn DoR) WhenResult {
	got := fn(g.ctx)
	g.ctx.got = got
	return WhenResult{
		ctx: g.ctx.copy(),
	}
}

func (w WhenResult) Then(fn func(Resolver)) {
	fn(w.ctx)
}

func When(context Tctx, fn func()) WhenResult {
	fn()
	ctx := context.(*TestContext)
	return WhenResult{
		ctx.copy(),
	}
}

func WhenR(context Tctx, fn func() interface{}) WhenResult {
	got := fn()
	ctx := context.(*TestContext)
	ctx.got = got

	return WhenResult{
		ctx.copy(),
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
