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

type Cache_t[Key_t comparable, Mapped_t any] struct {
	dict map[Key_t]*Value_t[Key_t, Mapped_t]
	root *Value_t[Key_t, Mapped_t]
}

type Less_t[Key_t comparable, Mapped_t any] func(a, b *Value_t[Key_t, Mapped_t]) bool

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
	SetNext(it, self.root)
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) CreateBack(key Key_t, value func() Mapped_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		return it, false
	}
	it = &Value_t[Key_t, Mapped_t]{Key: key, Value: value()}
	self.dict[key] = it
	SetPrev(it, self.root)
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) CreateSorted(key Key_t, value func() Mapped_t, where func(it, root *Value_t[Key_t, Mapped_t], less Less_t[Key_t, Mapped_t]), less Less_t[Key_t, Mapped_t]) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		return it, false
	}
	it = &Value_t[Key_t, Mapped_t]{Key: key, Value: value()}
	self.dict[key] = it
	where(it, self.root, less)
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) PushFront(key Key_t, value func() Mapped_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		CutList(it)
		SetNext(it, self.root)
		return it, false
	}
	it = &Value_t[Key_t, Mapped_t]{Key: key, Value: value()}
	self.dict[key] = it
	SetNext(it, self.root)
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) PushBack(key Key_t, value func() Mapped_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		CutList(it)
		SetPrev(it, self.root)
		return it, false
	}
	it = &Value_t[Key_t, Mapped_t]{Key: key, Value: value()}
	self.dict[key] = it
	SetPrev(it, self.root)
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) PushSorted(key Key_t, value func() Mapped_t, where func(it, root *Value_t[Key_t, Mapped_t], less Less_t[Key_t, Mapped_t]), less Less_t[Key_t, Mapped_t]) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		CutList(it)
		where(it, self.root, less)
		return it, false
	}
	it = &Value_t[Key_t, Mapped_t]{Key: key, Value: value()}
	self.dict[key] = it
	where(it, self.root, less)
	return it, true
}

func (self *Cache_t[Key_t, Mapped_t]) FindFront(key Key_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		CutList(it)
		SetNext(it, self.root)
	}
	return
}

func (self *Cache_t[Key_t, Mapped_t]) FindBack(key Key_t) (it *Value_t[Key_t, Mapped_t], ok bool) {
	if it, ok = self.dict[key]; ok {
		CutList(it)
		SetPrev(it, self.root)
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
		CutList(it)
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
func (self *Cache_t[Key_t, Mapped_t]) InsertionSortFront(less Less_t[Key_t, Mapped_t]) {
	for it1 := self.Front().Next(); it1 != self.End(); it1 = it1.Next() {
		for it2 := it1; it2.Prev() != self.End() && less(it2, it2.Prev()); {
			CutList(it2)
			SetPrev(it2, it2.Prev())
		}
	}
}

// linear if sorted before
func (self *Cache_t[Key_t, Mapped_t]) InsertionSortBack(less Less_t[Key_t, Mapped_t]) {
	for it1 := self.Back().Prev(); it1 != self.End(); it1 = it1.Prev() {
		for it2 := it1; it2.Next() != self.End() && less(it2, it2.Next()); {
			CutList(it2)
			SetNext(it2, it2.Next())
		}
	}
}

func SortValueFront[Key_t comparable, Mapped_t any](it, root *Value_t[Key_t, Mapped_t], less Less_t[Key_t, Mapped_t]) {
	for v := root.Next(); v != root; v = v.Next() {
		if less(it, v) {
			SetPrev(it, v)
			return
		}
	}
	SetPrev(it, root)
}

func SortValueBack[Key_t comparable, Mapped_t any](it, root *Value_t[Key_t, Mapped_t], less Less_t[Key_t, Mapped_t]) {
	for v := root.Prev(); v != root; v = v.Prev() {
		if less(it, v) {
			SetNext(it, v)
			return
		}
	}
	SetNext(it, root)
}
