package test

import "testing"

func head[T any](slice []T) (*T, bool) {
	if len(slice) > 0 {
		return &slice[0], true
	}

	return nil, false
}

func TestHead(t *testing.T) {
	var slice = []int{1, 2, 3, 4, 5}
	v, ok := head(slice)
	if !ok {
		t.Error("head fail")
	}

	t.Log(*v)
}

type Getter[T any] interface {
	Get() T
}

type Box[T any] struct {
	value T
}

func (b Box[T]) Get() T {
	return b.value
}

func TestStruct(t *testing.T) {
	var getter Getter[int] = Box[int]{10}
	t.Log(getter.Get())

	var tox = Box[string]{}
	t.Log(tox.Get())
}

type Defaulter interface {
	InitDefault()
}

type Foo struct {
	Msg   string
	Value int64
}

func (f *Foo) InitDefault() {
	f.Msg = "foo default"
	f.Value = 1024
}

type Bar struct {
	Msg   string
	Value float64
}

func (b *Bar) InitDefault() {
	b.Msg = "bar default"
	b.Value = 3.14
}

type defaultptr[T any] interface {
	*T
	Defaulter
}

// func Default[T Defaulter]() T {
// 	var v T
// 	v.InitDefault()
// 	return v
// }

func Default[T any, P defaultptr[T]]() T {
	var v T
	P.InitDefault(&v)
	return v
}

func TestInterface(t *testing.T) {
	var foo = Default[Foo]()
	var bar = Default[Bar]()

	t.Logf("foo: %+v", foo)
	t.Logf("bar: %+v", bar)
}
