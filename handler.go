package chain

import (
	"errors"
	"reflect"
)

var (
	ErrTypeNotFunction = errors.New("argument type not function")
	ErrInArgsMissMatch = errors.New("input arguments count not match")
	ErrOutCntMissMatch = errors.New("output parameter count not match")
)

type Handler struct {
	f    interface{}
	args []interface{}
}

func NewHandler(f interface{}, args ...interface{}) *Handler {
	res := new(Handler)
	res.f = f
	res.args = args
	return res
}

func (h *Handler) Call() []reflect.Value {
	f := reflect.ValueOf(h.f)
	typ := f.Type()
	if typ.Kind() != reflect.Func {
		panic(ErrTypeNotFunction)
	}
	// variable parameter, h.args less..
	if typ.NumIn() > len(h.args) {
		panic(ErrInArgsMissMatch)
	}
	inputs := make([]reflect.Value, len(h.args))
	for i := 0; i < len(h.args); i++ {
		if h.args[i] == nil {
			inputs[i] = reflect.Zero(f.Type().In(i))
		} else {
			inputs[i] = reflect.ValueOf(h.args[i])
		}
	}
	return f.Call(inputs)
}

func (h *Handler) BoolCall() bool {
	ret := h.Call()
	if len(ret) == 0 {
		panic(ErrOutCntMissMatch)
	}
	return ret[0].Bool()
}
