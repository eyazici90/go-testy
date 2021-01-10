package testy

import "testing"

type (
	Do  func(*TestContext)
	DoR func(*TestContext) interface{}

	// DoT func(*testing.T, *TestContext)
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
	return WhenResult{}
}

func (g GivenResult) WhenR(d DoR) WhenResult {
	ctx := g.ctx
	got := d(ctx)
	ctx.got = got
	return WhenResult{
		ctx: ctx,
	}
}

func (w WhenResult) Then(d Do) {
	d(w.ctx)
}
