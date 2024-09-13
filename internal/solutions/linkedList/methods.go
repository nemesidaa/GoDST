package linkedlist

import "errors"

// Allows to push to top (list[0]). tc O(1)
func (l *LinkedList) PushTop(value int) {
	n := &node{Value: value}
	n.Next = l.head
	l.head = n
	l.len++
}

// Allows to pop from top (list[0]), returns value of Popped item and error. tc O(1)
func (l *LinkedList) PopTop() (int, error) {
	if l.head == nil {
		return 0, errors.New("nil reference")
	}
	val := l.head.Value
	l.head = l.head.Next
	l.len--
	return val, nil
}

// Allows to push to bottom. tc O(n)
func (l *LinkedList) PushBack(value int) {
	n := &node{Value: value}
	if l.head == nil {
		l.head = n
	} else {
		l.head.append(n)
	}
	l.len++
}

// Allows to pop from bottom, returns value of Popped item and error. tc O(n)
func (l *LinkedList) PopBack() (int, error) {
	if l.head == nil {
		return 0, errors.New("nil reference")
	} else if l.head.Next == nil {
		return l.PopTop()
	} else {
		var pn, curr *node
		curr = l.head
		for curr.Next != nil {
			pn = curr
			curr = curr.Next
		}
		pn.Next = nil

		l.len--
		return curr.Value, nil
	}
}

// Allows to get length of list.
func (l *LinkedList) Len() uint {
	return l.len
}

func (l *LinkedList) String() string {
	if l.head == nil {
		return "nil"
	}
	return l.head.String()
}
