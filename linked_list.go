package main

import "fmt"

type node struct {
	data string
	next *node
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

func (l *linkedList) printList() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%#v -> ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	println()
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
