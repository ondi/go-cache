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

func CutList[Key_t comparable, Mapped_t any](it *Value_t[Key_t, Mapped_t]) {
	it.prev.next = it.next
	it.next.prev = it.prev
}

func SetPrev[Key_t comparable, Mapped_t any](it, at *Value_t[Key_t, Mapped_t]) {
	it.prev = at.prev
	at.prev.next = it
	at.prev = it
	it.next = at
}

func SetNext[Key_t comparable, Mapped_t any](it, at *Value_t[Key_t, Mapped_t]) {
	it.next = at.next
	at.next.prev = it
	at.next = it
	it.prev = at
}

func Swap[Key_t comparable, Mapped_t any](a, b *Value_t[Key_t, Mapped_t]) {
	a.next, b.next = b.next, a.next
	a.next.prev = a
	b.next.prev = b
	a.prev, b.prev = b.prev, a.prev
	a.prev.next = a
	b.prev.next = b
}

func MoveNext[Key_t comparable, Mapped_t any](it, at *Value_t[Key_t, Mapped_t]) {
	if it != at {
		CutList(it)
		SetNext(it, at)
	}
}

func MovePrev[Key_t comparable, Mapped_t any](it, at *Value_t[Key_t, Mapped_t]) {
	if it != at {
		CutList(it)
		SetPrev(it, at)
	}
}
