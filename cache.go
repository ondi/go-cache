//
// Use exclusive lock for all operations
// Key = key_type, Value = mapped_type, struct{Key, Value} = value_type
// to get Value from iterator use it.Mapped()
// PushFront() and PushBack() return value is same as stl map::insert
// {iterator, true} = inserted {key, value} 
// {intertor, false} = key already exists, no changes made
//

package cache

type Value_t struct {
	Key interface{}
	Value interface{}
	_prev * Value_t
	_next * Value_t
}

type Iterator struct {
	element * Value_t
}

func (self * Iterator) Key() interface{} {
	return self.element.Key
}

func (self * Iterator) Mapped() interface{} {
	return self.element.Value
}

func (self * Iterator) Value() Value_t {
	return *self.element
}

func (self * Iterator) Update(value interface{}) {
	self.element.Value = value
}

func (self * Iterator) Valid() bool {
	return self.element != nil
}

func (self * Iterator) Next() {
	self.element = self.element._next
}

func (self * Iterator) Prev() {
	self.element = self.element._prev
}

// list must not be empty
func cut_list(it * Value_t, first ** Value_t, last ** Value_t) (* Value_t) {
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
func set_first(it * Value_t, first ** Value_t, last ** Value_t) (* Value_t) {
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
func set_last(it * Value_t, first ** Value_t, last ** Value_t) (* Value_t) {
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

func move_first(it * Value_t, first ** Value_t, last ** Value_t) (* Value_t) {
	return set_first(cut_list(it, first, last), first, last)
}

func move_last(it * Value_t, first ** Value_t, last ** Value_t) (* Value_t) {
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

func (self * Cache) PushFront(key interface{}, value interface{}) (Iterator, bool) {
	if it, ok := self.dict[key]; ok {
		move_first(it, &self._first, &self._last)
		return Iterator{it}, false
	}
	it := &Value_t{Key: key, Value: value}
	self.dict[key] = it
	set_first(it, &self._first, &self._last)
	return Iterator{it}, true
}

func (self * Cache) PushBack(key interface{}, value interface{}) (Iterator, bool) {
	if it, ok := self.dict[key]; ok {
		move_last(it, &self._first, &self._last)
		return Iterator{it}, false
	}
	it := &Value_t{Key: key, Value: value}
	self.dict[key] = it
	set_last(it, &self._first, &self._last)
	return Iterator{it}, true
}

func (self * Cache) FindFront(key interface{}) (Iterator) {
	if it, ok := self.dict[key]; ok {
		move_first(it, &self._first, &self._last)
		return Iterator{it}
	}
	return Iterator{nil}
}

func (self * Cache) FindBack(key interface{}) (Iterator) {
	if it, ok := self.dict[key]; ok {
		move_last(it, &self._first, &self._last)
		return Iterator{it}
	}
	return Iterator{nil}
}

func (self * Cache) Find(key interface{}) (Iterator) {
	return Iterator{self.dict[key]}
}

func (self * Cache) Remove(key interface{}) {
	if it, ok := self.dict[key]; ok {
		cut_list(it, &self._first, &self._last)
		delete(self.dict, key)
	}
}

func (self * Cache) Front() (Iterator) {
	return Iterator{self._first}
}

func (self * Cache) Back() (Iterator) {
	return Iterator{self._last}
}

func (self * Cache) Size() (int) {
	return len(self.dict)
}
