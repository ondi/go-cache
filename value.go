//
//
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