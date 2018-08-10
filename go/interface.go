package main

import "fmt"

var v1 interface{} = 1     // 将int类型赋值给interface{}
var v2 interface{} = "abc" // 将string类型赋值给interface{}
var v3 interface{} = &v2   // 将*interface{}类型赋值给interface{}
var v4 interface{} = struct{ X int }{1}
var v5 interface{} = &struct{ X int }{1}

func Printf(fmt string, args ...interface{}) {}
func Println(args ...interface{})            {}

type Shaper interface {
	Area() float64
	//  Length() float32  //必须实现接口中定义的所有方法，才算实现了接口
}

type Square struct {
	side float64
}
type Rectangle struct {
	side   float64
	height float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}
func (s *Square) Another() float64 {
	return s.side * s.side
}

func (r *Rectangle) Area() float64 {
	return r.side * r.height / 2
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

/* method to determine the value of a stock position */
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
	return c.price
}

/* contract that defines different things that have value */
type valuable interface {
	getValue() float32
}

/* anything that satisfies the “valuable” interface is accepted */
func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

//func main() {
//	sql := &Square{5}
//	rect := &Rectangle{10, 4}
//
//	//接口可以直接引用实现了此接口的类型的方法
//
//	//接口声明,方式一
//	var areaIntf Shaper
//	areaIntf = sql
//
//	//接口声明,方式二
//	areaIntf1 := Shaper(sql)
//
//	//接口声明,方式三
//	areaIntf2 := sql
//
//	fmt.Println(areaIntf.Area())
//	fmt.Println(areaIntf1.Area())
//	fmt.Println(areaIntf2.Area())
//
//	//多态
//	areaIntf3 := Shaper(rect)
//	fmt.Println(areaIntf3.Area())
//
//	//鸭式辩型,类似其他语言的转型
//
//	var o valuable = stockPosition{"GOOG", 577.20, 4}
//	showValue(o)
//	o = car{"BMW", "M3", 66500}
//	showValue(o)
//
//	// 类型断言
//	if t, ok := areaIntf.(*Square); ok { //必须是*Square 而不是Square
//		fmt.Printf("%T", t) //*main.Square
//	}
//
//	// type switch
//	switch areaIntf.(type) {
//	case *Square:
//		fmt.Println("square")
//	case *Rectangle:
//		fmt.Println("Rectangle")
//	default:
//		fmt.Println("not find")
//	}
//
//}
