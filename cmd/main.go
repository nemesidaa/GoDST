package main

import (
	linkedlist "GoDST/internal/solutions/linkedList"
	"fmt"
)

func main() {
	l := linkedlist.NewLinkedList()
	l.PushTop(1)
	l.PushTop(2)
	v, err := l.PopBack()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
