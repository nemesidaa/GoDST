***Связный список***

Структура:

```go
type LinkedList struct { // size=16 (0x10)
    head *node
    len  uint
}
func (l *LinkedList) Len() uint
func (l *LinkedList) PopBack() (int, error)
func (l *LinkedList) PopTop() (int, error)
func (l *LinkedList) PushBack(value int)
func (l *LinkedList) PushTop(value int)
func (l *LinkedList) String() string
```

Usage:

```go
l := linkedlist.NewLinkedList()

	if l.Len() != 0 {
		fmt.Errorf("l.Len() = %d, want %d", l.Len(), 0)
	}

	for i := 0; i < 10; i++ {
		l.PushTop(i)
	}
	fmt.Println(l.String())
	if l.Len() != 10 {
		fmt.Errorf("l.Len() = %d, want %d", l.Len(), 10)
	}

	for i := 0; i < 10; i++ {
		v, err := l.PopBack()
		fmt.Println(v)
		if err != nil {
			fmt.Errorf("l.PopTop() error = %v", err)
		}
		if v != i {
			fmt.Errorf("l.PopTop() = %d, want %d", v, i)
		}
	}
	fmt.Println(l.Len(), l.String())
```


Построена на рекурсии, низкаяя отпимизированность операций с Back индексами вследствие рекомендуется пользоваться Top-ами, best practice - fifo стэк.
