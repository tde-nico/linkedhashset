package linkedhashset

type LinkedHashSet[T comparable] struct {
	end *LinkedHashSetNode[T]
	set map[T]*LinkedHashSetNode[T]
}

func (l *LinkedHashSet[T]) PushBack(key T) {
	if _, ok := l.set[key]; ok {
		return
	}
	end := l.end
	curr := end.Prev
	curr.Next = &LinkedHashSetNode[T]{Key: key, Prev: curr, Next: end}
	end.Prev = curr.Next
	l.set[key] = curr.Next
}

func (l *LinkedHashSet[T]) PushFront(key T) {
	if _, ok := l.set[key]; ok {
		return
	}
	end := l.end
	curr := end.Next
	curr.Prev = &LinkedHashSetNode[T]{Key: key, Prev: end, Next: curr}
	end.Next = curr.Prev
	l.set[key] = curr.Prev
}

func (l *LinkedHashSet[T]) PopBack() T {
	if l.end == nil {
		var zeroValue T
		return zeroValue
	}
	var key T
	if l.end.Prev != nil {
		key = l.end.Prev.Key
		delete(l.set, l.end.Prev.Key)
		l.end.Prev = l.end.Prev.Prev
		l.end.Prev.Next = l.end
	}
	return key
}

func (l *LinkedHashSet[T]) PopFront() T {
	if l.end == nil {
		var zeroValue T
		return zeroValue
	}
	var key T
	if l.end.Next != nil {
		key = l.end.Next.Key
		delete(l.set, l.end.Next.Key)
		l.end.Next = l.end.Next.Next
		l.end.Next.Prev = l.end
	}
	return key
}

func (l *LinkedHashSet[T]) PushBackArr(keys []T) {
	for _, key := range keys {
		l.PushBack(key)
	}
}

func (l *LinkedHashSet[T]) PushFrontArr(keys []T) {
	for i := len(keys) - 1; i >= 0; i-- {
		l.PushFront(keys[i])
	}
}

func (l *LinkedHashSet[T]) PopBackArr(amount int) []T {
	length := min(amount, len(l.set))
	res := make([]T, 0, length)
	for i := 0; i < length; i++ {
		key := l.PopBack()
		res = append(res, key)
	}
	return res
}

func (l *LinkedHashSet[T]) PopFrontArr(amount int) []T {
	length := min(amount, len(l.set))
	res := make([]T, 0, length)
	for i := 0; i < length; i++ {
		key := l.PopFront()
		res = append(res, key)
	}
	return res
}

func (l *LinkedHashSet[T]) Contains(key T) bool {
	_, ok := l.set[key]
	return ok
}

func (l *LinkedHashSet[T]) Remove(key T) {
	if _, ok := l.set[key]; !ok {
		return
	}
	node := l.set[key]
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	delete(l.set, key)
}

func (l *LinkedHashSet[T]) RemoveArr(keys []T) {
	for _, key := range keys {
		l.Remove(key)
	}
}

func (l *LinkedHashSet[T]) Clear() {
	l.end.Prev = l.end
	l.end.Next = l.end
	l.set = make(map[T]*LinkedHashSetNode[T])
}

func (l *LinkedHashSet[T]) GetAll() []T {
	res := make([]T, 0, len(l.set))
	for node := l.end.Next; node != l.end; node = node.Next {
		res = append(res, node.Key)
	}
	return res
}

func (l *LinkedHashSet[T]) GetAllReverse() []T {
	res := make([]T, 0, len(l.set))
	for node := l.end.Prev; node != l.end; node = node.Prev {
		res = append(res, node.Key)
	}
	return res
}

func (l *LinkedHashSet[T]) Size() int {
	return len(l.set)
}

func (l *LinkedHashSet[T]) Empty() bool {
	return len(l.set) == 0
}
