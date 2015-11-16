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
	value interface{}
	_prev * Value_t
	_next * Value_t
}

func (self * Value_t) Key() interface{} {
	return self.key
}

func (self * Value_t) Mapped() interface{} {
	return self.value
}

func (self * Value_t) Value() Value_t {
	return *self
}

func (self * Value_t) Update(value interface{}) {
	self.value = value
}

func (self * Value_t) Next() * Value_t {
	return self._next
}

func (self * Value_t) Prev() * Value_t {
	return self._prev
}

func cut_list(it * Value_t) * Value_t {
	it._prev._next = it._next
	it._next._prev = it._prev
	it._prev = nil	// be on the safe side
	it._next = nil	// be on the safe side
	return it
}

func set_before(it * Value_t, at * Value_t) * Value_t {
	it._prev = at._prev
	at._prev = it
	it._prev._next = it
	it._next = at
	return it
}

func set_after(it * Value_t, at * Value_t) * Value_t {
	it._next = at._next
	at._next = it;
	it._next._prev = it
	it._prev = at
	return it
}

func move_first(it * Value_t, root * Value_t) * Value_t {
	return set_after(cut_list(it), root)
}

func move_last(it * Value_t, root * Value_t) * Value_t {
	return set_before(cut_list(it), root)
}

type Cache struct {
	dict map[interface{}]*Value_t
	_root * Value_t
}

func New() (self * Cache) {
	self = &Cache{}
	self.dict = map[interface{}]*Value_t{}
	self._root = &Value_t{}
	self._root._prev = self._root
	self._root._next = self._root
	return
}

func (self * Cache) PushFront(key interface{}, value interface{}) (it * Value_t, ok bool) {
	if it, ok = self.dict[key]; ok {
		move_first(it, self._root)
		return it, false
	}
	it = &Value_t{key: key, value: value}
	set_after(it, self._root)
	self.dict[key] = it
	return it, true
}

func (self * Cache) PushBack(key interface{}, value interface{}) (it * Value_t, ok bool) {
	if it, ok = self.dict[key]; ok {
		move_last(it, self._root)
		return it, false
	}
	it = &Value_t{key: key, value: value}
	set_before(it, self._root)
	self.dict[key] = it
	return it, true
}

func (self * Cache) FindFront(key interface{}) * Value_t {
	if it, ok := self.dict[key]; ok {
		move_first(it, self._root)
		return it
	}
	return self.End()
}

func (self * Cache) FindBack(key interface{}) * Value_t {
	if it, ok := self.dict[key]; ok {
		move_last(it, self._root)
		return it
	}
	return self.End()
}

func (self * Cache) Find(key interface{}) * Value_t {
	return self.dict[key]
}

func (self * Cache) Remove(key interface{}) {
	if it, ok := self.dict[key]; ok {
		cut_list(it)
		delete(self.dict, key)
	}
}

func (self * Cache) Front() * Value_t {
	return self._root._next
}

func (self * Cache) Back() * Value_t {
	return self._root._prev
}

func (self * Cache) End() * Value_t {
	return self._root
}

func (self * Cache) Size() (int) {
	return len(self.dict)
}
