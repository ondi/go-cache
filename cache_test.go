package cache

import "fmt"
import "testing"

type MyLess1_t struct {}

func (* MyLess1_t) Less(a * Value_t, b * Value_t) bool {
	return a.Key().(int) < b.Key().(int)
}

func ExampleSort10() {
	cc := New()
	cc.CreateFront(1, func() interface{} {return 10})
	cc.PushFront(5, func() interface{} {return 50})
	cc.UpdateFront(1, func(interface{}) interface{} {return 100})
	cc.CreateFront(7, func() interface{} {return 70})
	cc.PushFront(3, func() interface{} {return 30})
	cc.UpdateFront(5, func(interface{}) interface{} {return 500})
	
	cc.InsertionSortFront(&MyLess1_t{})
	for it := cc.Front(); it != cc.End(); it = it.Next() {
		fmt.Printf("%v %v\n", it.Key(), it.Value())
	}
/* Output:
1 100
3 30
5 500
7 70
*/
}

func ExampleSort20() {
	cc := New()
	cc.CreateFront(1, func() interface{} {return 10})
	cc.PushFront(5, func() interface{} {return 50})
	cc.UpdateFront(1, func(interface{}) interface{} {return 100})
	cc.CreateFront(7, func() interface{} {return 70})
	cc.PushFront(3, func() interface{} {return 30})
	cc.UpdateFront(7, func(interface{}) interface{} {return 700})
	
	cc.InsertionSortBack(&MyLess1_t{})
	for it := cc.Front(); it != cc.End(); it = it.Next() {
		fmt.Printf("%v %v\n", it.Key(), it.Value())
	}
/* Output:
7 700
5 50
3 30
1 100
*/
}

var c1 = New()
var c2 = New()
var c3 = New()

func BenchmarkCache10(b * testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1.CreateFront(i, func() interface{} {return i})
	}
}

func BenchmarkCache20(b * testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c2.PushFront(i, func() interface{} {return i})
	}
}

func BenchmarkCache30(b * testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c3.UpdateFront(i, func(interface{}) interface{} {return i})
	}
}
