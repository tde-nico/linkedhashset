package linkedhashset

type LinkedHashSetNode[T comparable] struct {
	Key  T
	Prev *LinkedHashSetNode[T]
	Next *LinkedHashSetNode[T]
}

type LinkedHashSetInterface[T any] interface {
	PushBack(key T)
	PushFront(key T)
	PopBack() T
	PopFront() T
	PushBackArr(keys []T)
	PushFrontArr(keys []T)
	PopBackArr(amount int) []T
	PopFrontArr(amount int) []T
	Contains(key T) bool
	Remove(key T)
	RemoveArr(keys []T)
	Clear()
	GetAll() []T
	GetAllReverse() []T
	Size() int
	Empty() bool
}

func NewLinkedHashSet[T comparable]() *LinkedHashSet[T] {
	end := LinkedHashSetNode[T]{}
	end.Prev = &end
	end.Next = &end
	return &LinkedHashSet[T]{
		end: &end,
		set: make(map[T]*LinkedHashSetNode[T]),
	}
}

func NewLinkedHashSetSafe[T comparable]() *LinkedHashSetSafe[T] {
	return &LinkedHashSetSafe[T]{
		LinkedHashSet: *NewLinkedHashSet[T](),
	}
}
