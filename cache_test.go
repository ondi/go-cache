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

func ExampleSort20() {
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

func BenchmarkCache10(b * testing.B) {
	var c1 = New()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1.CreateFront(i, func() interface{} {return i})
	}
}

func BenchmarkCache20(b * testing.B) {
	var c2 = New()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c2.PushFront(i, func() interface{} {return i})
	}
}

func BenchmarkCache30(b * testing.B) {
	var c3 = New()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c3.UpdateFront(i, func() interface{} {return i})
	}
}
