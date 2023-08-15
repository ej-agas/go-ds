package doublylinkedlist

type data interface {
	int | int32 | int64 | string | float32 | float64
}

type node[T data] struct {
	data T
	prev *node[T]
	next *node[T]
}

type doublyLinkedList[T data] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func (l *doublyLinkedList[T]) prepend(data T) {
	newNode := &node[T]{data: data}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length++
		return
	}

	second := l.head
	l.head = newNode
	l.head.next = second
	second.prev = l.head

	if second.next == nil {
		l.tail = second
	}

	l.length++
}

func (l doublyLinkedList[T]) getData() []T {
	toPrint := l.head

	data := make([]T, 0, l.length)

	for l.length != 0 {
		data = append(data, toPrint.data)
		toPrint = toPrint.next
		l.length--
	}

	return data
}
