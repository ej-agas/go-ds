package doublylinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrependDoublyLinkedList(t *testing.T) {
	linkedList := doublyLinkedList[string]{}

	node1 := "John"
	node2 := "Jake"
	node3 := "Jane"
	node4 := "Jack"
	node5 := "Fred"

	linkedList.prepend(node5)
	linkedList.prepend(node4)
	linkedList.prepend(node3)
	linkedList.prepend(node2)
	linkedList.prepend(node1)

	assert.Equal(t, []string{node1, node2, node3, node4, node5}, linkedList.getData())
	assert.Equal(t, node5, linkedList.tail.data)
	assert.Equal(t, node1, linkedList.head.data)
}

func TestTraverseForwardLinkedList(t *testing.T) {
	linkedList := doublyLinkedList[string]{}

	node1 := "John"
	node2 := "Jake"
	node3 := "Jane"
	node4 := "Jack"
	node5 := "Fred"

	linkedList.prepend(node5)
	linkedList.prepend(node4)
	linkedList.prepend(node3)
	linkedList.prepend(node2)
	linkedList.prepend(node1)

	expected := []string{node1, node2, node3, node4, node5}

	currentNode := linkedList.head
	for i := 0; i < linkedList.length; i++ {
		assert.Equal(t, expected[i], currentNode.data)
		currentNode = currentNode.next
	}
}

func TestTraverseBackwardLinkedList(t *testing.T) {
	linkedList := doublyLinkedList[string]{}

	node1 := "John"
	node2 := "Jake"
	node3 := "Jane"
	node4 := "Jack"
	node5 := "Fred"

	linkedList.prepend(node5)
	linkedList.prepend(node4)
	linkedList.prepend(node3)
	linkedList.prepend(node2)
	linkedList.prepend(node1)

	expected := []string{node5, node4, node3, node2, node1}
	currentNode := linkedList.tail
	for i := 0; i < linkedList.length; i++ {
		assert.Equal(t, expected[i], currentNode.data)
		currentNode = currentNode.prev
	}
}
