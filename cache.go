//
// Use exclusive lock for all operations
// PushFront() and PushBack() output:
// {iterator, true} = inserted {key, value}
// {intertor, false} = key already exists, no changes made
// iterate over cache: for i := c.Front(); i != c.End(); i = i.Next() {...}
//

package cache

type Value_t struct {
	key interface{}
	mapped interface{}
	prev * Value_t
	next * Value_t
}

func (self * Value_t) Key() interface{} {
	return self.key
}

func (self * Value_t) Mapped() interface{} {
	return self.mapped
}

func (self * Value_t) Update(mapped interface{}) {
	self.mapped = mapped
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

func move_after(it * Value_t, at * Value_t) * Value_t {
	return set_after(cut_list(it), at)
}

func move_before(it * Value_t, at * Value_t) * Value_t {
	return set_before(cut_list(it), at)
}

type Cache struct {
	dict map[interface{}]*Value_t
	_root * Value_t
}

func New() (self * Cache) {
	self = &Cache{}
	self.dict = map[interface{}]*Value_t{}
	self._root = &Value_t{}
	self._root.prev = self._root
	self._root.next = self._root
	return
}

func (self * Cache) PushFront(key interface{}, value interface{}) (it * Value_t, ok bool) {
	if it, ok = self.dict[key]; ok {
		move_after(it, self._root)
		return it, false
	}
	it = &Value_t{key: key, mapped: value}
	set_after(it, self._root)
	self.dict[key] = it
	return it, true
}

func (self * Cache) PushBack(key interface{}, value interface{}) (it * Value_t, ok bool) {
	if it, ok = self.dict[key]; ok {
		move_before(it, self._root)
		return it, false
	}
	it = &Value_t{key: key, mapped: value}
	set_before(it, self._root)
	self.dict[key] = it
	return it, true
}

func (self * Cache) FindFront(key interface{}) * Value_t {
	if it, ok := self.dict[key]; ok {
		move_after(it, self._root)
		return it
	}
	return self.End()
}

func (self * Cache) FindBack(key interface{}) * Value_t {
	if it, ok := self.dict[key]; ok {
		move_before(it, self._root)
		return it
	}
	return self.End()
}

func (self * Cache) Find(key interface{}) * Value_t {
	if it, ok := self.dict[key]; ok {
		return it
	}
	return self.End()
}

func (self * Cache) Remove(key interface{}) {
	if it, ok := self.dict[key]; ok {
		cut_list(it)
		delete(self.dict, key)
	}
}

func (self * Cache) Front() * Value_t {
	return self._root.next
}

func (self * Cache) Back() * Value_t {
	return self._root.prev
}

func (self * Cache) End() * Value_t {
	return self._root
}

func (self * Cache) Size() (int) {
	return len(self.dict)
}
