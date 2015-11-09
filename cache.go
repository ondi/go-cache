//
// Use exclusive lock for all operations
// PushFront() and PushBack() return value:
// {iterator, true} = inserted {key, value}
// {intertor, false} = key already exists, no changes made
// iterate over cache: for i := c.Front(); i.Valid(); i = i.Next() {...}
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

func (self * Value_t) Valid() bool {
	return self != nil
}

func (self * Value_t) Next() * Value_t {
	return self._next
}

func (self * Value_t) Prev() * Value_t {
	return self._prev
}

// list must not be empty
func cut_list(it * Value_t, first ** Value_t, last ** Value_t) * Value_t {
	if it._prev == nil {
		*first = it._next
	} else {
		it._prev._next = it._next
	}
	if it._next == nil {
		*last = it._prev
	} else {
		it._next._prev = it._prev
	}
	return it
}

// list may be empty
func set_first(it * Value_t, first ** Value_t, last ** Value_t) * Value_t {
	it._prev = nil
	it._next = *first
	if *first == nil {
		*last = it
	} else {
		(*first)._prev = it
	}
	*first = it
	return it
}

// list may be empty
func set_last(it * Value_t, first ** Value_t, last ** Value_t) * Value_t {
	it._next = nil
	it._prev = *last
	if *last == nil {
		*first = it
	} else {
		(*last)._next = it
	}
	*last = it
	return it
}

func move_first(it * Value_t, first ** Value_t, last ** Value_t) * Value_t {
	return set_first(cut_list(it, first, last), first, last)
}

func move_last(it * Value_t, first ** Value_t, last ** Value_t) * Value_t {
	return set_last(cut_list(it, first, last), first, last)
}

type Cache struct {
	dict map[interface{}]*Value_t
	_first * Value_t
	_last * Value_t
}

func New() (self * Cache) {
	self = &Cache{}
	self.dict = map[interface{}]*Value_t{}
	return
}

func (self * Cache) PushFront(key interface{}, value interface{}) (it * Value_t, ok bool) {
	if it, ok = self.dict[key]; ok {
		move_first(it, &self._first, &self._last)
		return it, false
	}
	it = &Value_t{key: key, value: value}
	set_first(it, &self._first, &self._last)
	self.dict[key] = it
	return it, true
}

func (self * Cache) PushBack(key interface{}, value interface{}) (it * Value_t, ok bool) {
	if it, ok = self.dict[key]; ok {
		move_last(it, &self._first, &self._last)
		return it, false
	}
	it = &Value_t{key: key, value: value}
	set_last(it, &self._first, &self._last)
	self.dict[key] = it
	return it, true
}

func (self * Cache) FindFront(key interface{}) * Value_t {
	if it, ok := self.dict[key]; ok {
		move_first(it, &self._first, &self._last)
		return it
	}
	return nil
}

func (self * Cache) FindBack(key interface{}) * Value_t {
	if it, ok := self.dict[key]; ok {
		move_last(it, &self._first, &self._last)
		return it
	}
	return nil
}

func (self * Cache) Find(key interface{}) * Value_t {
	return self.dict[key]
}

func (self * Cache) Remove(key interface{}) {
	if it, ok := self.dict[key]; ok {
		cut_list(it, &self._first, &self._last)
		delete(self.dict, key)
	}
}

func (self * Cache) Front() * Value_t {
	return self._first
}

func (self * Cache) Back() * Value_t {
	return self._last
}

func (self * Cache) Size() (int) {
	return len(self.dict)
}
