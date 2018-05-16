package main

type Base struct {
}

func (base *Base) Foo() {}
func (base *Base) Bar() {}

//type Foo struct {
//	Base
//}

type Foo struct {
	*Base
}

func (foo *Foo) Bar() {
	foo.Base.Bar()
}
