package singlylinkedlist

type node struct {
	data string
	next *node
}

func newNode(data string, next *node) *node {
	return &node{data: data, next: next}
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n node) {
	second := l.head
	l.head = &n
	l.head.next = second
	l.length++
}

func (l linkedList) getListData() []string {
	toPrint := l.head

	data := make([]string, 0, l.length)

	for l.length != 0 {
		data = append(data, toPrint.data)
		toPrint = toPrint.next
		l.length--
	}

	return data
}

func (l *linkedList) delete(data string) {
	if l.length == 0 {
		return
	}

	prev := l.head

	if prev.data == data {
		l.head = l.head.next
		l.length--

		return
	}

	for prev.next.data != data {
		if prev.next.next == nil {
			return
		}

		prev = prev.next
	}

	prev.next = prev.next.next
	l.length--
}
