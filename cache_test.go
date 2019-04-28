package cache

import "fmt"
import "testing"

type MyLess1_t struct {}

func (* MyLess1_t) Less(a * Value_t, b * Value_t) bool {
	return a.Key().(int) < b.Key().(int)
}

func ExampleSort1() {
	cc := New()
	cc.PushFront(1, 10)
	cc.PushFront(5, 50)
	cc.PushFront(9, 90)
	cc.PushFront(7, 70)
	cc.PushFront(3, 30)
	
	cc.InsertionSortFront(&MyLess1_t{})
	for it := cc.Front(); it != cc.End(); it = it.Next() {
		fmt.Printf("%v\n", it.Key())
	}
/* Output:
1
3
5
7
9
*/
}

func ExampleSort2() {
	cc := New()
	cc.PushFront(1, 10)
	cc.PushFront(5, 50)
	cc.PushFront(9, 90)
	cc.PushFront(7, 70)
	cc.PushFront(3, 30)
	
	cc.InsertionSortBack(&MyLess1_t{})
	for it := cc.Front(); it != cc.End(); it = it.Next() {
		fmt.Printf("%v\n", it.Key())
	}
/* Output:
9
7
5
3
1
*/
}

func Test1(t * testing.T) {

}
