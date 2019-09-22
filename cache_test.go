package cache

import "fmt"
import "testing"

type MyLess1_t struct {}

func (* MyLess1_t) Less(a * Value_t, b * Value_t) bool {
	return a.Key().(int) < b.Key().(int)
}

func ExampleSort1() {
	cc := New()
	cc.CreateFront(1, func() interface{} {return 10})
	cc.PushFront(5, func() interface{} {return 50})
	cc.UpdateFront(9, func() interface{} {return 90})
	cc.CreateFront(7, func() interface{} {return 70})
	cc.PushFront(3, func() interface{} {return 30})
	cc.UpdateFront(2, func() interface{} {return 20})
	
	cc.InsertionSortFront(&MyLess1_t{})
	for it := cc.Front(); it != cc.End(); it = it.Next() {
		fmt.Printf("%v\n", it.Key())
	}
/* Output:
1
2
3
5
7
9
*/
}

func ExampleSort2() {
	cc := New()
	cc.CreateFront(1, func() interface{} {return 10})
	cc.PushFront(5, func() interface{} {return 50})
	cc.UpdateFront(9, func() interface{} {return 90})
	cc.CreateFront(7, func() interface{} {return 70})
	cc.PushFront(3, func() interface{} {return 30})
	cc.UpdateFront(2, func() interface{} {return 20})
	
	cc.InsertionSortBack(&MyLess1_t{})
	for it := cc.Front(); it != cc.End(); it = it.Next() {
		fmt.Printf("%v\n", it.Key())
	}
/* Output:
9
7
5
3
2
1
*/
}

func BenchmarkCache1(b * testing.B) {
	c := New()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c.CreateFront(i, func() interface{} {return i})
	}
}

func BenchmarkCache2(b * testing.B) {
	c := New()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c.PushFront(i, func() interface{} {return i})
	}
}

func BenchmarkCache3(b * testing.B) {
	c := New()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c.UpdateFront(i, func() interface{} {return i})
	}
}
