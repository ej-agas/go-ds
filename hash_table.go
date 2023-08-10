package main

const ArraySize = 7

type bucketNode struct {
	key  string
	next *bucketNode
}

func (b *bucket) insert(k string) {
	if b.search(k) == true {
		panic("error: key already exists")
	}

	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

func (b *bucket) search(k string) bool {
	current := b.head

	for current != nil {
		if current.key == k {
			return true
		}
		current = current.next
	}

	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previous := b.head
	for previous.next != nil {
		if previous.next.key == k {
			previous.next = previous.next.next
		}

		previous = previous.next
		break
	}
}

type bucket struct {
	head *bucketNode
}

type HashTable struct {
	array [ArraySize]*bucket
}

func NewHashTable() *HashTable {
	h := &HashTable{}

	for i := range h.array {
		h.array[i] = &bucket{}
	}

	return h
}

func (h *HashTable) Insert(key string) {
	h.array[hash(key)].insert(key)
}

func (h *HashTable) Search(key string) bool {
	return h.array[hash(key)].search(key)
}

func (h *HashTable) Delete(key string) {
	h.array[hash(key)].delete(key)
}

func hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}
