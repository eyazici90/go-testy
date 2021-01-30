package testy

import (
	"reflect"
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
		Subject() interface{}
		T() *testing.T
	}
)

type TestContext struct {
	t       *testing.T
	val     map[reflect.Type]interface{}
	got     interface{}
	subject interface{}
}

func NewTestContext(t *testing.T) *TestContext {
	return &TestContext{
		t:   t,
		val: make(map[reflect.Type]interface{}),
	}
}

func (t *TestContext) UseThe(i interface{}) {
	tp := reflect.TypeOf(i)
	t.val[tp] = i
}

func (t *TestContext) SetThe(i interface{}) {
	t.subject = i
}

func (t *TestContext) The(i interface{}) interface{} {
	tp := reflect.TypeOf(i)
	return t.val[tp]
}

func (t *TestContext) Got() interface{} { return t.got }

func (t *TestContext) T() *testing.T { return t.t }

func (t *TestContext) Subject() interface{} { return t.subject }

func (t *TestContext) Use() *TestContext { return t }

func (t *TestContext) A(i interface{}) {
	//fixture generate
	tp := reflect.TypeOf(i)
	t.val[tp] = i
}
