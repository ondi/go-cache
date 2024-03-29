//
// go clean -testcache && go test ./...
//

package cache

import (
	"fmt"
	"testing"
)

func IntKey[Mapped_t any](a, b *Value_t[int, Mapped_t]) bool {
	return a.Key < b.Key
}

func IntKeyValue(a, b *Value_t[int, int]) bool {
	return a.Value < b.Value
}

func Example_sort10() {
	cc := New[int, int]()
	cc.CreateFront(1, func(p *int) { *p = 10 }, func(*int) {})
	cc.PushFront(5, func(p *int) { *p = 50 }, func(*int) {})
	it, _ := cc.FindFront(1)
	it.Value = 100
	cc.CreateFront(7, func(p *int) { *p = 70 }, func(*int) {})
	cc.PushFront(3, func(p *int) { *p = 30 }, func(*int) {})
	it, _ = cc.FindFront(5)
	it.Value = 500

	cc.InsertionSortFront(IntKey[int])
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
	cc := New[int, int]()
	cc.CreateFront(1, func(p *int) { *p = 10 }, func(*int) {})
	cc.PushFront(5, func(p *int) { *p = 50 }, func(*int) {})
	it, _ := cc.FindFront(1)
	it.Value = 100
	cc.CreateFront(7, func(p *int) { *p = 70 }, func(*int) {})
	cc.PushFront(3, func(p *int) { *p = 30 }, func(*int) {})
	it, _ = cc.FindFront(7)
	it.Value = 700
	cc.InsertionSortBack(IntKey[int])
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
	cc := New[int, int]()
	cc.CreateFront(1, func(p *int) { *p = 10 }, func(*int) {})
	cc.PushFront(5, func(p *int) { *p = 50 }, func(*int) {})
	it, _ := cc.FindFront(1)
	it.Value = 100
	cc.CreateFront(7, func(p *int) { *p = 70 }, func(*int) {})
	cc.PushFront(3, func(p *int) { *p = 30 }, func(*int) {})
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

var c1 = New[int, int]()
var c2 = New[int, int]()
var c3 = New[int, int]()

func Benchmark_cache10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c1.CreateFront(i, func(p *int) { *p = i }, func(*int) {})
	}
}

func Benchmark_cache20(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c2.PushFront(i, func(p *int) { *p = i }, func(*int) {})
	}
}

func Benchmark_cache30(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		it, _ := c3.FindFront(i)
		it.Value = i
	}
}
