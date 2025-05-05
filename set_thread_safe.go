package linkedhashset

import (
	"sync"
)

type LinkedHashSetSafe[T comparable] struct {
	sync.Mutex
	LinkedHashSet[T]
}

func (l *LinkedHashSetSafe[T]) PushBack(key T) {
	l.Lock()
	defer l.Unlock()
	l.LinkedHashSet.PushBack(key)
}

func (l *LinkedHashSetSafe[T]) PushFront(key T) {
	l.Lock()
	defer l.Unlock()
	l.LinkedHashSet.PushFront(key)
}

func (l *LinkedHashSetSafe[T]) PopBack() T {
	l.Lock()
	defer l.Unlock()
	return l.LinkedHashSet.PopBack()
}

func (l *LinkedHashSetSafe[T]) PopFront() T {
	l.Lock()
	defer l.Unlock()
	return l.LinkedHashSet.PopFront()
}

func (l *LinkedHashSetSafe[T]) PushBackArr(keys []T) {
	l.Lock()
	defer l.Unlock()
	l.LinkedHashSet.PushBackArr(keys)
}

func (l *LinkedHashSetSafe[T]) PushFrontArr(keys []T) {
	l.Lock()
	defer l.Unlock()
	l.LinkedHashSet.PushFrontArr(keys)
}

func (l *LinkedHashSetSafe[T]) PopBackArr(amount int) []T {
	l.Lock()
	defer l.Unlock()
	return l.LinkedHashSet.PopBackArr(amount)
}

func (l *LinkedHashSetSafe[T]) PopFrontArr(amount int) []T {
	l.Lock()
	defer l.Unlock()
	return l.LinkedHashSet.PopFrontArr(amount)
}

func (l *LinkedHashSetSafe[T]) Contains(key T) bool {
	l.Lock()
	defer l.Unlock()
	return l.LinkedHashSet.Contains(key)
}

func (l *LinkedHashSetSafe[T]) Remove(key T) {
	l.Lock()
	defer l.Unlock()
	l.LinkedHashSet.Remove(key)
}

func (l *LinkedHashSetSafe[T]) RemoveArr(keys []T) {
	l.Lock()
	defer l.Unlock()
	l.LinkedHashSet.RemoveArr(keys)
}

func (l *LinkedHashSetSafe[T]) Clear() {
	l.Lock()
	defer l.Unlock()
	l.LinkedHashSet.Clear()
}

func (l *LinkedHashSetSafe[T]) GetAll() []T {
	l.Lock()
	defer l.Unlock()
	return l.LinkedHashSet.GetAll()
}

func (l *LinkedHashSetSafe[T]) GetAllReverse() []T {
	l.Lock()
	defer l.Unlock()
	return l.LinkedHashSet.GetAllReverse()
}

func (l *LinkedHashSetSafe[T]) Size() int {
	l.Lock()
	defer l.Unlock()
	return l.LinkedHashSet.Size()
}

func (l *LinkedHashSetSafe[T]) Empty() bool {
	l.Lock()
	defer l.Unlock()
	return l.LinkedHashSet.Empty()
}
