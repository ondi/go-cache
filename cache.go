//
// Create/Push/Find return value:
// {iterator, true} = inserted/found {key, value}
// {intertor, false} = key already exists/not found, no changes made
// iterate over cache:
// for it := c.Front(); it != c.End(); it = it.Next() {
//	fmt.Printf("%v=%v\n", it.Key, it.Value)
//}
//

package cache

type Value_t[Key_t comparable, Mapped_t any] struct {
	Key   Key_t // read only
	Value Mapped_t
	prev  *Value_t[Key_t, Mapped_t]
	next  *Value_t[Key_t, Mapped_t]
}

func (self *Value_t[Key_t, Mapped_t]) Next() *Value_t[Key_t, Mapped_t] {
	return self.next
}

func (self *Value_t[Key_t, Mapped_t]) Prev() *Value_t[Key_t, Mapped_t] {
	return self.prev
}

func (it *Value_t[Key_t, Mapped_t]) cut_list() *Value_t[Key_t, Mapped_t] {
	it.prev.next = it.next
	it.next.prev = it.prev
	return it
}

func (it *Value_t[Key_t, Mapped_t]) set_before(at *Value_t[Key_t, Mapped_t]) *Value_t[Key_t, Mapped_t]  {
	it.prev = at.prev
	at.prev.next = it
	at.prev = it
	it.next = at
	return it
}

func (it *Value_t[Key_t, Mapped_t]) set_after(at *Value_t[Key_t, Mapped_t]) *Value_t[Key_t, Mapped_t] {
	it.next = at.next
	at.next.prev = it
	at.next = it
	it.prev = at
	return it
}

func (a *Value_t[Key_t, Mapped_t]) Swap(b *Value_t[Key_t, Mapped_t]) {
	if a.next == b {
		a.prev.next = b
		b.next.prev = a
		b.prev = a.prev
		a.prev = b
		a.next = b.next
		b.next = a
		return
	}
	if a.prev == b {
		a.next.prev = b
		b.prev.next = a
		b.next = a.next
		a.next = b
		a.prev = b.prev
		b.prev = a
		return
	}
	a.next.prev = b
	a.prev.next = b
	b.next.prev = a
	b.prev.next = a
	a.prev, b.prev = b.prev, a.prev
	a.next, b.next = b.next, a.next
}

func (it *Value_t[Key_t, Mapped_t]) MoveAfter(at *Value_t[Key_t, Mapped_t]) *Value_t[Key_t, Mapped_t] {
	if it != at {
		it.cut_list().set_after(at)
	}
	return it
}

func (it *Value_t[Key_t, Mapped_t]) MoveBefore(at *Value_t[Key_t, Mapped_t]) *Value_t[Key_t, Mapped_t] {
	if it != at {
		it.cut_list().set_before(at)
	}
	return it
}

type Cache_t[Key_t comparable, Mapped_t any] struct {
	dict map[Key_t]*Value_t[Key_t, Mapped_t]
	root *Value_t[Key_t, Mapped_t]
}

func New[Key_t comparable, Mapped_t any]() (self *Cache_t[Key_t, Mapped_t]) {
	self = &Cache_t[Key_t, Mapped_t]{}
	self.Clear()
	return
}

func (self *Cache_t[Key_t, Mapped_t]) Clear() {
	self.dict = map[Key_t]*Value_t[Key_t, Mapped_t]{}
	self.root = &Value_t[Key_t, Mapped_t]{}
	self.root.prev = self.root
	self.root.next = self.root
}

func (self *Cache_t[Key_t, Mapped_t]) CreateFront(key Key_t, value func() Mapped_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		return it, false
	}
	it = &Value_t[Key_t, Mapped_t]{Key: key, Value: value()}
	self.dict[key] = it
	it.set_after(self.root)
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) CreateFront2(key Key_t, value func() *Value_t[Key_t, Mapped_t]) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		return it, false
	}
	if it = value(); it != nil {
		it.Key = key
		self.dict[key] = it
		it.set_after(self.root)
	}
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) CreateBack(key Key_t, value func() Mapped_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		return it, false
	}
	it = &Value_t[Key_t, Mapped_t]{Key: key, Value: value()}
	self.dict[key] = it
	it.set_before(self.root)
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) CreateBack2(key Key_t, value func() *Value_t[Key_t, Mapped_t]) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		return it, false
	}
	if it = value(); it != nil {
		it.Key = key
		self.dict[key] = it
		it.set_before(self.root)
	}
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) PushFront(key Key_t, value func() Mapped_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		it.cut_list().set_after(self.root)
		return it, false
	}
	it = &Value_t[Key_t, Mapped_t]{Key: key, Value: value()}
	self.dict[key] = it
	it.set_after(self.root)
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) PushFront2(key Key_t, value func() *Value_t[Key_t, Mapped_t]) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		it.cut_list().set_after(self.root)
		return it, false
	}
	if it = value(); it != nil {
		it.Key = key
		self.dict[key] = it
		it.set_after(self.root)
	}
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) PushBack(key Key_t, value func() Mapped_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		it.cut_list().set_before(self.root)
		return it, false
	}
	it = &Value_t[Key_t, Mapped_t]{Key: key, Value: value()}
	self.dict[key] = it
	it.set_before(self.root)
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) PushBack2(key Key_t, value func() *Value_t[Key_t, Mapped_t]) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		it.cut_list().set_before(self.root)
		return it, false
	}
	if it = value(); it != nil {
		it.Key = key
		self.dict[key] = it
		it.set_before(self.root)
	}
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) FindFront(key Key_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		it.cut_list().set_after(self.root)
	}
	return
}

func (self *Cache_t[Key_t, Mapped_t]) FindBack(key Key_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		it.cut_list().set_before(self.root)
	}
	return
}

func (self *Cache_t[Key_t, Mapped_t]) Find(key Key_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	it, ok = self.dict[key]
	return
}

func (self *Cache_t[Key_t, Mapped_t]) Remove(key Key_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		delete(self.dict, key)
		it.cut_list()
	}
	return
}

func (self *Cache_t[Key_t, Mapped_t]) Front() *Value_t[Key_t, Mapped_t] {
	return self.root.next
}

func (self *Cache_t[Key_t, Mapped_t]) Back() *Value_t[Key_t, Mapped_t] {
	return self.root.prev
}

func (self *Cache_t[Key_t, Mapped_t]) End() *Value_t[Key_t, Mapped_t] {
	return self.root
}

func (self *Cache_t[Key_t, Mapped_t]) Size() int {
	return len(self.dict)
}

// linear if sorted before
func (self *Cache_t[Key_t, Mapped_t]) InsertionSortFront(cmp MyLess[Key_t, Mapped_t]) {
	for it1 := self.Front().Next(); it1 != self.End(); it1 = it1.Next() {
		for it2 := it1; it2.Prev() != self.End() && cmp.Less(it2, it2.Prev()); {
			it2.cut_list().set_before(it2.Prev())
		}
	}
}

// linear if sorted before
func (self *Cache_t[Key_t, Mapped_t]) InsertionSortBack(cmp MyLess[Key_t, Mapped_t]) {
	for it1 := self.Back().Prev(); it1 != self.End(); it1 = it1.Prev() {
		for it2 := it1; it2.Next() != self.End() && cmp.Less(it2, it2.Next()); {
			it2.cut_list().set_after(it2.Next())
		}
	}
}

type MyLess[Key_t comparable, Mapped_t any] interface {
	Less(a *Value_t[Key_t, Mapped_t], b *Value_t[Key_t, Mapped_t]) bool
}

type reverse[Key_t comparable, Mapped_t any] struct {
	MyLess[Key_t, Mapped_t]
}

func (self *reverse[Key_t, Mapped_t]) Less(a *Value_t[Key_t, Mapped_t], b *Value_t[Key_t, Mapped_t]) bool {
	return self.Less(b, a)
}

func Reverse[Key_t comparable, Mapped_t any](cmp MyLess[Key_t, Mapped_t]) MyLess[Key_t, Mapped_t] {
	return &reverse[Key_t, Mapped_t]{cmp}
}
