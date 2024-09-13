package linkedlist

import "fmt"

// Allows to push to bottom
func (n *node) append(next *node) {
	if n.Next == nil {
		n.Next = next
	} else {
		n.Next.append(next)
	}
}

func (n *node) String() string {
	if n.Next == nil {
		return fmt.Sprintf("%d->nil", n.Value)
	} else {
		return fmt.Sprintf("%d->", n.Value) + n.Next.String()
	}
}
