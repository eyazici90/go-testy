package testy

import "testing"

type (
	Do  func(Tctx)
	DoR func(Tctx) interface{}
)

func Given(t *testing.T, d Do) GivenResult {
	ctx := NewTestContext(t)
	d(ctx)
	return GivenResult{
		ctx: ctx,
	}
}

func (g GivenResult) When(d Do) WhenResult {
	d(g.ctx)
	return WhenResult{
		ctx: g.ctx,
	}
}

func (g GivenResult) WhenR(d DoR) WhenResult {
	ctx := g.ctx
	got := d(ctx)
	ctx.got = got
	return WhenResult{
		ctx: ctx,
	}
}

func (w WhenResult) Then(fn func(Resolver)) {
	fn(w.ctx)
}

func When(context Tctx, fn func()) WhenResult {
	fn()
	ctx := context.(*TestContext)
	return WhenResult{
		ctx,
	}
}

func WhenR(context Tctx, fn func() interface{}) WhenResult {
	got := fn()
	ctx := context.(*TestContext)
	ctx.got = got

	return WhenResult{
		ctx,
	}
}

func Then(ctx *TestContext, fn func()) {
	fn()
}
