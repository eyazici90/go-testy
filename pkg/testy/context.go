package testy

import (
	"reflect"
	"testing"
)

type TestContext struct {
	t   *testing.T
	val map[reflect.Type]interface{}
	got interface{}
}

func NewTestContext(t *testing.T) *TestContext {
	return &TestContext{
		t:   t,
		val: make(map[reflect.Type]interface{}),
	}
}

func (t *TestContext) UseThe(any interface{}) {
	tp := reflect.TypeOf(any)
	t.val[tp] = any
}

func (t *TestContext) The(any interface{}) interface{} {
	tp := reflect.TypeOf(any)
	return t.val[tp]
}

func (t *TestContext) Got() interface{} { return t.got }

func (t *TestContext) Use() *TestContext { return t }

func (t *TestContext) A(any interface{}) {
	//fixture generate
	tp := reflect.TypeOf(any)
	t.val[tp] = any
}
