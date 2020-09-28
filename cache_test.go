package cache

import (
	"fmt"
	"testing"
)

type MyLess1_t struct{}

func (*MyLess1_t) Less(a *Value_t, b *Value_t) bool {
	return a.Key.(int) < b.Key.(int)
}

func Example_sort10() {
	cc := New()
	cc.CreateFront(1, func() interface{} { return 10 })
	cc.PushFront(5, func() interface{} { return 50 })
	it, _ := cc.FindFront(1)
	it.Value = 100
	cc.CreateFront(7, func() interface{} { return 70 })
	cc.PushFront(3, func() interface{} { return 30 })
	it, _ = cc.FindFront(5)
	it.Value = 500

	cc.InsertionSortFront(&MyLess1_t{})
	for it := cc.Front(); it != cc.End(); it = it.Next() {
		fmt.Printf("%v %v\n", it.Key, it.Value)
	}
	// Output:
	// 1 100
	// 3 30
	// 5 500
	// 7 70
}

func Example_sort20() {
	cc := New()
	cc.CreateFront(1, func() interface{} { return 10 })
	cc.PushFront(5, func() interface{} { return 50 })
	it, _ := cc.FindFront(1)
	it.Value = 100
	cc.CreateFront(7, func() interface{} { return 70 })
	cc.PushFront(3, func() interface{} { return 30 })
	it, _ = cc.FindFront(7)
	it.Value = 700
	cc.InsertionSortBack(&MyLess1_t{})
	for it := cc.Front(); it != cc.End(); it = it.Next() {
		fmt.Printf("%v %v\n", it.Key, it.Value)
	}
	// Output:
	// 7 700
	// 5 50
	// 3 30
	// 1 100
}

func Example_swap10() {
	cc := New()
	cc.CreateFront(1, func() interface{} { return 10 })
	cc.PushFront(5, func() interface{} { return 50 })
	it, _ := cc.FindFront(1)
	it.Value = 100
	cc.CreateFront(7, func() interface{} { return 70 })
	cc.PushFront(3, func() interface{} { return 30 })
	it, _ = cc.FindFront(7)
	it.Value = 700
	it1, _ := cc.Find(1)
	it2, _ := cc.Find(7)
	Swap(it1, it2)
	// 7, 3, 1, 5
	for it := cc.Front(); it != cc.End(); it = it.Next() {
		fmt.Printf("%v %v\n", it.Key, it.Value)
	}
	// Output:
	// 1 100
	// 3 30
	// 7 700
	// 5 50
}

var c1 = New()
var c2 = New()
var c3 = New()

func Benchmark_cache10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1.CreateFront(i, func() interface{} { return i })
	}
}

func Benchmark_cache20(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c2.PushFront(i, func() interface{} { return i })
	}
}

func Benchmark_cache30(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		it, _ := c3.FindFront(i)
		it.Value = i
	}
}
