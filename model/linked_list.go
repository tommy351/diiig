package model

// LinkedList is a ordered doubly linked list.
type LinkedList struct {
	head *linkedListNode
	len  int
}

// Insert adds an element to the list and make sure the list is ordered.
func (l *LinkedList) Insert(element *Element) {
	cursor := l.head
	l.len++

	for cursor != nil {
		next := cursor.next

		// If the value of cursor is larger than the element.
		// The element should be inserted after the cursor when either:
		//
		// - the cursor is the last element
		// - the value of the next element is smaller than or equals to the element
		if cursor.value.Value > element.Value {
			if next == nil || next.value.Value <= element.Value {
				l.insertAfter(element, cursor)
				return
			}
		}

		cursor = next
	}

	l.prepend(element)
}

// Get returns the element with the specified key.
func (l *LinkedList) Get(key string) *Element {
	cursor := l.head

	for cursor != nil {
		if cursor.value.Key == key {
			return cursor.value
		}

		cursor = cursor.next
	}

	return nil
}

// Remove deletes the element with the specified key from the list.
func (l *LinkedList) Remove(key string) {
	cursor := l.head

	for cursor != nil {
		prev := cursor.prev
		next := cursor.next

		if cursor.value.Key == key {
			if prev != nil {
				prev.next = next
			} else {
				l.head = next
			}

			if next != nil {
				next.prev = prev
			}

			l.len--
		}

		cursor = next
	}
}

// Each returns an iterator to iterate over the list.
func (l *LinkedList) Each() Iterator {
	return &linkedListIterator{
		node:  l.head,
		index: -1,
	}
}

// Length returns the number of elements in the list.
func (l *LinkedList) Length() int {
	return l.len
}

func (l *LinkedList) prepend(element *Element) {
	node := &linkedListNode{
		next:  l.head,
		value: element,
	}

	if l.head != nil {
		l.head.prev = node
	}

	l.head = node
}

func (l *LinkedList) insertAfter(element *Element, cursor *linkedListNode) {
	next := cursor.next
	cursor.next = &linkedListNode{
		prev:  cursor,
		next:  next,
		value: element,
	}

	if next != nil {
		next.prev = cursor.next
	}
}

type linkedListNode struct {
	prev  *linkedListNode
	next  *linkedListNode
	value *Element
}

type linkedListIterator struct {
	node  *linkedListNode
	value *Element
	index int
}

func (l *linkedListIterator) Next() bool {
	if l.node == nil {
		return false
	}

	l.value = l.node.value
	l.node = l.node.next
	l.index++
	return true
}

func (l *linkedListIterator) Value() *Element {
	return l.value
}

func (l *linkedListIterator) Index() int {
	return l.index
}
