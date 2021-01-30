package testy

import (
	"testing"
)

type (
	Tctx interface {
		Registrar
		Resolver
	}
	Registrar interface {
		SetThe(i interface{})
	}
	Resolver interface {
		Got() interface{}
		Err() error
		Subject() interface{}
		T() *testing.T
	}
)

type TestContext struct {
	t       *testing.T
	got     interface{}
	err     error
	subject interface{}
}

func NewTestContext(t *testing.T) *TestContext {
	return &TestContext{
		t: t,
	}
}

func (t *TestContext) SetThe(i interface{}) {
	t.subject = i
}

func (t *TestContext) Got() interface{} { return t.got }

func (t *TestContext) Err() error { return t.err }

func (t *TestContext) T() *testing.T { return t.t }

func (t *TestContext) Subject() interface{} { return t.subject }

func (t *TestContext) copy() *TestContext {
	return &TestContext{
		err:     t.err,
		got:     t.got,
		subject: t.subject,
		t:       t.t,
	}
}
