//
// Use exclusive lock for all operations
// PushFront() and PushBack() output:
// {iterator, true} = inserted {key, value}
// {intertor, false} = key already exists, no changes made
// iterate over cache: for i := c.Front(); i != c.End(); i = i.Next() {...}
//

package cache

type Less_t interface {
	Less(a * Value_t, b * Value_t) bool
}

type Value_t struct {
	key interface{}
	value interface{}
	prev * Value_t
	next * Value_t
}

func (self * Value_t) Key() interface{} {
	return self.key
}

func (self * Value_t) Value() interface{} {
	return self.value
}

func (self * Value_t) Update(value interface{}) {
	self.value = value
}

func (self * Value_t) Next() * Value_t {
	return self.next
}

func (self * Value_t) Prev() * Value_t {
	return self.prev
}

func cut_list(it * Value_t) * Value_t {
	it.prev.next = it.next
	it.next.prev = it.prev
	return it
}

func set_before(it * Value_t, at * Value_t) * Value_t {
	it.prev = at.prev
	at.prev.next = it
	at.prev = it
	it.next = at
	return it
}

func set_after(it * Value_t, at * Value_t) * Value_t {
	it.next = at.next
	at.next.prev = it
	at.next = it;
	it.prev = at
	return it
}

// very complicated
func Swap(a * Value_t, b * Value_t) {
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
	b.next.prev = a
	a.prev.next = b
	b.prev.next = a
	
	temp := a.prev
	a.prev = b.prev
	b.prev = temp
	
	temp = a.next
	a.next = b.next
	b.next = temp
}

func MoveAfter(it * Value_t, at * Value_t) {
	if it != at {
		set_after(cut_list(it), at)
	}
}

func MoveBefore(it * Value_t, at * Value_t) {
	if it != at {
		set_before(cut_list(it), at)
	}
}

type Cache_t struct {
	dict map[interface{}]*Value_t
	root * Value_t
}

func New() (self * Cache_t) {
	self = &Cache_t{}
	self.Clear()
	return
}

func (self * Cache_t) Clear() {
	self.dict = map[interface{}]*Value_t{}
	self.root = &Value_t{}
	self.root.prev = self.root
	self.root.next = self.root
}

func (self * Cache_t) CreateFront(key interface{}, value func() interface{}) (it * Value_t, ok bool) {
	if it, ok = self.dict[key]; ok {
		return it, false
	}
	it = &Value_t{key: key, value: value()}
	self.dict[key] = it
	set_after(it, self.root)
	return it, true
}

func (self * Cache_t) CreateBack(key interface{}, value func() interface{}) (it * Value_t, ok bool) {
	if it, ok = self.dict[key]; ok {
		return it, false
	}
	it = &Value_t{key: key, value: value()}
	self.dict[key] = it
	set_before(it, self.root)
	return it, true
}

func (self * Cache_t) PushFront(key interface{}, value func() interface{}) (it * Value_t, ok bool) {
	if it, ok = self.CreateFront(key, value); !ok {
		set_after(cut_list(it), self.root)
	}
	return
}

func (self * Cache_t) PushBack(key interface{}, value func() interface{}) (it * Value_t, ok bool) {
	if it, ok = self.CreateBack(key, value); !ok {
		set_before(cut_list(it), self.root)
	}
	return
}

func (self * Cache_t) UpdateFront(key interface{}, value func() interface{}) (it * Value_t, ok bool) {
	if it, ok = self.dict[key]; ok {
		it.Update(value())
		set_after(cut_list(it), self.root)
		return it, false
	}
	it = &Value_t{key: key, value: value()}
	self.dict[key] = it
	set_after(it, self.root)
	return it, true
}

func (self * Cache_t) UpdateBack(key interface{}, value func() interface{}) (it * Value_t, ok bool) {
	if it, ok = self.dict[key]; ok {
		it.Update(value())
		set_before(cut_list(it), self.root)
		return it, false
	}
	it = &Value_t{key: key, value: value()}
	self.dict[key] = it
	set_before(it, self.root)
	return it, true
}

func (self * Cache_t) FindFront(key interface{}) * Value_t {
	if it, ok := self.dict[key]; ok {
		set_after(cut_list(it), self.root)
		return it
	}
	return self.End()
}

func (self * Cache_t) FindBack(key interface{}) * Value_t {
	if it, ok := self.dict[key]; ok {
		set_before(cut_list(it), self.root)
		return it
	}
	return self.End()
}

func (self * Cache_t) Find(key interface{}) * Value_t {
	if it, ok := self.dict[key]; ok {
		return it
	}
	return self.End()
}

func (self * Cache_t) Remove(key interface{}) {
	if it, ok := self.dict[key]; ok {
		cut_list(it)
		delete(self.dict, key)
	}
}

func (self * Cache_t) Front() * Value_t {
	return self.root.next
}

func (self * Cache_t) Back() * Value_t {
	return self.root.prev
}

func (self * Cache_t) End() * Value_t {
	return self.root
}

func (self * Cache_t) Size() int {
	return len(self.dict)
}

// takes linear time if sorted before
func (self * Cache_t) InsertionSortFront(less Less_t) {
	for it1 := self.Front().Next(); it1 != self.End(); it1 = it1.Next() {
		for it2 := it1; it2.Prev() != self.End() && less.Less(it2, it2.Prev()); {
			set_before(cut_list(it2), it2.Prev())
		}
	}
}

// takes linear time if sorted before
func (self * Cache_t) InsertionSortBack(less Less_t) {
	for it1 := self.Back().Prev(); it1 != self.End(); it1 = it1.Prev() {
		for it2 := it1; it2.Next() != self.End() && less.Less(it2, it2.Next()); {
			set_after(cut_list(it2), it2.Next())
		}
	}
}

type reverse struct {
	Less_t
}

func (self * reverse) Less(a * Value_t, b * Value_t) bool {
	return self.Less_t.Less(b, a)
}

func Reverse(less Less_t) Less_t {
	return &reverse{less}
}
