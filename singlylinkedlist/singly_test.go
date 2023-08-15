package singlylinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrependSinglyLinkedList(t *testing.T) {
	myList := linkedList{}

	node1 := newNode("foo", nil)
	node2 := newNode("bar", nil)
	node3 := newNode("baz", nil)

	myList.prepend(*node3)
	myList.prepend(*node2)
	myList.prepend(*node1)

	assert.Equal(t, []string{"foo", "bar", "baz"}, myList.getListData())
}

func TestDeleteSinglyLinkedList(t *testing.T) {
	myList := linkedList{}

	node1 := newNode("foo", nil)
	node2 := newNode("bar", nil)
	node3 := newNode("baz", nil)

	myList.prepend(*node3)
	myList.prepend(*node2)
	myList.prepend(*node1)

	assert.Equal(t, []string{"foo", "bar", "baz"}, myList.getListData())

	myList.delete("foo")

	assert.Equal(t, []string{"bar", "baz"}, myList.getListData())

	myList.delete("baz")

	assert.Equal(t, []string{"bar"}, myList.getListData())
}
