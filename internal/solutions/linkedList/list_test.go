package linkedlist_test

import (
	linkedlist "GoDST/internal/solutions/linkedList"
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := linkedlist.NewLinkedList()

	if l.Len() != 0 {
		t.Errorf("l.Len() = %d, want %d", l.Len(), 0)
	}

	for i := 0; i < 10; i++ {
		l.PushTop(i)
	}
	fmt.Println(l.String())
	if l.Len() != 10 {
		t.Errorf("l.Len() = %d, want %d", l.Len(), 10)
	}

	for i := 0; i < 10; i++ {
		v, err := l.PopBack()
		fmt.Println(v)
		if err != nil {
			t.Errorf("l.PopTop() error = %v", err)
		}
		if v != i {
			t.Errorf("l.PopTop() = %d, want %d", v, i)
		}
	}
	fmt.Println(l.Len(), l.String())
}

func BenchmarkLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := linkedlist.NewLinkedList()

		if l.Len() != 0 {
			b.Errorf("l.Len() = %d, want %d", l.Len(), 0)
		}

		for i := 0; i < 10; i++ {
			l.PushTop(i)
		}

		if l.Len() != 10 {
			b.Errorf("l.Len() = %d, want %d", l.Len(), 10)
		}

		for i := 0; i < 9; i++ {
			v, err := l.PopBack()
			fmt.Println(v)
			if err != nil {
				b.Errorf("l.PopBack() error = %v", err)
			}
			if v != i {
				b.Errorf("l.PopBack() = %d, want %d", v, i)
			}
		}
	}
}
