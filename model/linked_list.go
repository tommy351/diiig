package model

// LinkedList is a ordered doubly linked list.
type LinkedList struct {
	head *linkedListNode
	len  int
}

func (l *LinkedList) Insert(element *Element) {
	cursor := l.head
	l.len++

	for cursor != nil {
		next := cursor.next

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

func (l *LinkedList) Each() Enumerator {
	return &linkedListEnumerator{
		node: l.head,
	}
}

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

type linkedListEnumerator struct {
	node  *linkedListNode
	value *Element
}

func (l *linkedListEnumerator) Next() bool {
	if l.node == nil {
		return false
	}

	l.value = l.node.value
	l.node = l.node.next
	return true
}

func (l *linkedListEnumerator) Value() *Element {
	return l.value
}
