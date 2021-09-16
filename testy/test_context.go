package testy

import (
	"testing"
)

var doNothing func() = func() {}

type TestContext struct {
	t          *testing.T
	got        interface{}
	err        error
	subject    interface{}
	beforeEach func()
	afterEach  func()
}

func NewTestContext(t *testing.T) *TestContext {
	return &TestContext{
		t:          t,
		beforeEach: doNothing,
		afterEach:  doNothing,
	}
}

func (t *TestContext) SetThe(i interface{}) {
	t.subject = i
}

func (t *TestContext) BeforeEach(fn func()) {
	t.beforeEach = fn
}

func (t *TestContext) AfterEach(fn func()) {
	t.afterEach = fn
}

func (t *TestContext) Got() interface{} { return t.got }

func (t *TestContext) Err() error { return t.err }

func (t *TestContext) T() *testing.T { return t.t }

func (t *TestContext) Subject() interface{} { return t.subject }

func (t *TestContext) copy() *TestContext {
	return &TestContext{
		err:        t.err,
		got:        t.got,
		subject:    t.subject,
		t:          t.t,
		beforeEach: t.beforeEach,
		afterEach:  t.afterEach,
	}
}
