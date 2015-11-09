//
// Use exclusive lock for all operations
// Key = key_type, Value = mapped_type, struct{Key, Value} = value_type
// to get Value from iterator use it.Mapped()
// PushFront() and PushBack() return value is same as stl map::insert
// {iterator, true} = inserted {key, value} 
// {intertor, false} = key already exists, no changes made
//

package cache

import "container/list"

type Value_t struct {
	Key interface{}
	Value interface{}
}

type Iterator struct {
	element * list.Element
}

func (self * Iterator) Key() interface{} {
	return self.element.Value.(Value_t).Key
}

func (self * Iterator) Mapped() interface{} {
	return self.element.Value.(Value_t).Value
}

func (self * Iterator) Value() Value_t {
	return self.element.Value.(Value_t)
}

func (self * Iterator) Update(value interface{}) {
	self.element.Value = Value_t{Key: self.element.Value.(Value_t).Key, Value: value}
}

func (self * Iterator) Valid() bool {
	return self.element != nil
}

func (self * Iterator) Next() {
	self.element = self.element.Next()
}

func (self * Iterator) Prev() {
	self.element = self.element.Prev()
}

type Cache struct {
	dict map[interface{}]Iterator
	lru_list list.List
}

func New() (self * Cache) {
	self = &Cache{}
	self.dict = map[interface{}]Iterator{}
	return
}

func (self * Cache) PushFront(key interface{}, value interface{}) (it Iterator, ok bool) {
	if it, ok = self.dict[key]; ok {
		self.lru_list.MoveToFront(it.element)
		return it, false
	}
	it = Iterator{self.lru_list.PushFront(Value_t{Key: key, Value: value})}
	self.dict[key] = it
	return it, true
}

func (self * Cache) PushBack(key interface{}, value interface{}) (it Iterator, ok bool) {
	if it, ok = self.dict[key]; ok {
		self.lru_list.MoveToBack(it.element)
		return it, false
	}
	it = Iterator{self.lru_list.PushFront(Value_t{Key: key, Value: value})}
	self.dict[key] = it
	return it, true
}

func (self * Cache) FindFront(key interface{}) Iterator {
	if mapped, ok := self.dict[key]; ok {
		self.lru_list.MoveToFront(mapped.element)
		return mapped
	}
	return Iterator{nil}
}

func (self * Cache) FindBack(key interface{}) Iterator {
	if mapped, ok := self.dict[key]; ok {
		self.lru_list.MoveToBack(mapped.element)
		return mapped
	}
	return Iterator{nil}
}

func (self * Cache) Find(key interface{}) Iterator {
	return self.dict[key]
}

func (self * Cache) Remove(key interface{}) {
	if mapped, ok := self.dict[key]; ok {
		self.lru_list.Remove(mapped.element)
		delete(self.dict, key)
	}
}

func (self * Cache) Front() Iterator {
	return Iterator{self.lru_list.Front()}
}

func (self * Cache) Back() Iterator {
	return Iterator{self.lru_list.Back()}
}

func (self * Cache) Size() int {
	return len(self.dict)
}
