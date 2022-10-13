//
// go clean -testcache && go test ./...
//

package cache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func IntKey[Mapped_t any](a, b *Value_t[int, Mapped_t]) bool {
	return a.Key < b.Key
}

func IntKeyValue(a, b *Value_t[int, int]) bool {
	return a.Value < b.Value
}

func Example_create10() {
	cc := New[int, int]()
	values := []int{0,1,1,2,3,3,3,4,5,6,7,8,9}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(values), func(i, j int) {values[i], values[j] = values[j], values[i]})
	for _, v := range values {
		cc.CreateSorted(v, func() int { return v }, SortValueFront[int, int], IntKeyValue)
	}
	for it := cc.Front(); it != cc.End(); it = it.Next() {
		fmt.Printf("%v %v\n", it.Key, it.Value)
	}
	// Output:
	// 0 0
	// 1 1
	// 2 2
	// 3 3
	// 4 4
	// 5 5
	// 6 6
	// 7 7
	// 8 8
	// 9 9
}

func Example_push10() {
	cc := New[int, int]()
	values := []int{0,1,1,2,3,3,3,4,5,6,7,8,9}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(values), func(i, j int) {values[i], values[j] = values[j], values[i]})
	for _, v := range values {
		cc.PushSorted(v, func() int { return v }, SortValueBack[int, int], IntKeyValue)
	}
	for it := cc.Front(); it != cc.End(); it = it.Next() {
		fmt.Printf("%v %v\n", it.Key, it.Value)
	}
	// Output:
	// 9 9
	// 8 8
	// 7 7
	// 6 6
	// 5 5
	// 4 4
	// 3 3
	// 2 2
	// 1 1
	// 0 0
}

func Example_sort10() {
	cc := New[int, int]()
	cc.CreateFront(1, func() int { return 10 })
	cc.PushFront(5, func() int { return 50 })
	it, _ := cc.FindFront(1)
	it.Value = 100
	cc.CreateFront(7, func() int { return 70 })
	cc.PushFront(3, func() int { return 30 })
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
	cc.CreateFront(1, func() int { return 10 })
	cc.PushFront(5, func() int { return 50 })
	it, _ := cc.FindFront(1)
	it.Value = 100
	cc.CreateFront(7, func() int { return 70 })
	cc.PushFront(3, func() int { return 30 })
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
	cc.CreateFront(1, func() int { return 10 })
	cc.PushFront(5, func() int { return 50 })
	it, _ := cc.FindFront(1)
	it.Value = 100
	cc.CreateFront(7, func() int { return 70 })
	cc.PushFront(3, func() int { return 30 })
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
		c1.CreateFront(i, func() int { return i })
	}
}

func Benchmark_cache20(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c2.PushFront(i, func() int { return i })
	}
}

func Benchmark_cache30(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		it, _ := c3.FindFront(i)
		it.Value = i
	}
}
