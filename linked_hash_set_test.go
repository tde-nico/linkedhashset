package linkedhashset

import (
	"testing"
)

func test(t *testing.T, set LinkedHashSetInterface[string]) {
	// Front
	set.PushFront("a")
	set.PushFront("b")
	set.PushFront("c")
	size := set.Size()
	if size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}
	element := set.PopBack()
	if element != "a" {
		t.Errorf("Expected 'a', got %s", element)
	}
	element = set.PopFront()
	if element != "c" {
		t.Errorf("Expected 'c', got %s", element)
	}
	element = set.PopBack()
	if element != "b" {
		t.Errorf("Expected 'b', got %s", element)
	}
	element = set.PopBack()
	if element != "" {
		t.Errorf("Expected '', got %s", element)
	}
	size = set.Size()
	if size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Back
	set.PushBack("a")
	set.PushBack("b")
	set.PushBack("c")
	size = set.Size()
	if size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}
	element = set.PopFront()
	if element != "a" {
		t.Errorf("Expected 'a', got %s", element)
	}
	element = set.PopBack()
	if element != "c" {
		t.Errorf("Expected 'c', got %s", element)
	}
	element = set.PopFront()
	if element != "b" {
		t.Errorf("Expected 'b', got %s", element)
	}
	element = set.PopFront()
	if element != "" {
		t.Errorf("Expected '', got %s", element)
	}
	size = set.Size()
	if size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Arr
	set.PushBackArr([]string{"a", "b", "c"})
	size = set.Size()
	if size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}
	elements := set.PopFrontArr(2)
	if len(elements) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(elements))
	}
	if elements[0] != "a" {
		t.Errorf("Expected 'a', got %s", elements[0])
	}
	if elements[1] != "b" {
		t.Errorf("Expected 'b', got %s", elements[1])
	}
	elements = set.PopBackArr(2)
	if len(elements) != 1 {
		t.Errorf("Expected 1 element, got %d", len(elements))
	}
	if elements[0] != "c" {
		t.Errorf("Expected 'c', got %s", elements[0])
	}

	// Reverse
	set.PushFrontArr([]string{"a", "b", "c"})
	size = set.Size()
	if size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}
	elements = set.GetAllReverse()
	if len(elements) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(elements))
	}
	if elements[0] != "c" {
		t.Errorf("Expected 'c', got %s", elements[0])
	}
	if elements[1] != "b" {
		t.Errorf("Expected 'b', got %s", elements[1])
	}
	if elements[2] != "a" {
		t.Errorf("Expected 'a', got %s", elements[2])
	}

	// Contains
	set.PushBack("a")
	set.PushBack("b")
	set.PushBack("c")
	if !set.Contains("a") {
		t.Errorf("Expected to contain 'a'")
	}
	if !set.Contains("b") {
		t.Errorf("Expected to contain 'b'")
	}
	if !set.Contains("c") {
		t.Errorf("Expected to contain 'c'")
	}
	if set.Contains("d") {
		t.Errorf("Expected not to contain 'd'")
	}

	// Remove
	set.Remove("a")
	if set.Contains("a") {
		t.Errorf("Expected not to contain 'a'")
	}
	set.RemoveArr([]string{"b", "c"})
	if set.Contains("b") {
		t.Errorf("Expected not to contain 'b'")
	}
	if set.Contains("c") {
		t.Errorf("Expected not to contain 'c'")
	}
	if set.Contains("d") {
		t.Errorf("Expected not to contain 'd'")
	}

	// Clear
	set.PushBack("a")
	set.PushBack("b")
	set.PushBack("c")
	set.Clear()
	size = set.Size()
	if size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Empty
	if !set.Empty() {
		t.Errorf("Expected empty set")
	}
	set.PushBack("a")
	if set.Empty() {
		t.Errorf("Expected not empty set")
	}
}

func TestLinkedHashSet(t *testing.T) {
	set := NewLinkedHashSet[string]()
	test(t, set)
}

func TestLinkedHashSetSafe(t *testing.T) {
	set := NewLinkedHashSetSafe[string]()
	test(t, set)
}
