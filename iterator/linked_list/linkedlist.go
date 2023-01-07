package main

import (
	"errors"
	"fmt"
)

const (
	LINEAR = iota
	DESCENDING
)

type Node struct {
	data int
	next *Node
}

func NewNode(v int) *Node {
	return &Node{data: v}
}

type LinkedList interface {
	Add(int)
	Pop()
	Size() int
	CreateLinkedListIterator(uint) (LinkedListIterator, error)
}

type linkedList struct {
	node *Node
	size int
}

func NewLinkedList() LinkedList {
	return &linkedList{}
}

func (l *linkedList) Add(value int) {
	node := NewNode(value)
	if l.node == nil {
		l.node = node
	} else if l.node.next == nil {
		l.node.next = node
	} else {
		cur := l.node
		for cur.next != nil {
			cur = cur.next
		}

		cur.next = node
	}
	l.size++
}

func (l *linkedList) Pop() {
	if l.node == nil {
		return
	}

	if l.node.next == nil {
		l.node = nil
		return
	}

	cur := l.node
	prev := cur
	for cur.next != nil {
		prev = cur
		cur = cur.next
	}

	prev.next = nil
	l.size--
}

func (l *linkedList) Print() {
	if l.node == nil {
		return
	}

	if l.node.next == nil {
		fmt.Println(l.node.data)
	}

	cur := l.node
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}
}

func (l *linkedList) Size() int {
	return l.size
}

// creates an iterator based on the provided type
func (l *linkedList) CreateLinkedListIterator(iteratorType uint) (LinkedListIterator, error) {
	var iterator LinkedListIterator
	var err error
	switch iteratorType {
	case LINEAR:
		iterator, err = NewLinearLinkedListIterator(l)
	case DESCENDING:
		iterator, err = NewDescendingLinkedListIterator(l)
	default:
		err = errors.New("Invalid terator type")
	}

	//linkedListIterator, err := NewLinearLinkedListIterator(l)
	if err != nil {
		return nil, err
	}

	return iterator, nil
}

type LinkedListIterator interface {
	Iterate() bool
	Next() *Node
}

// iterator that goes one element after another (ascending order)
type LinearLinkedListIterator struct {
	currentPosition *Node
}

func NewLinearLinkedListIterator(list *linkedList) (*LinearLinkedListIterator, error) {
	if list == nil {
		return nil, errors.New("Invalid linked list")
	}
	return &LinearLinkedListIterator{
		currentPosition: list.node,
	}, nil
}

func (l *LinearLinkedListIterator) Iterate() bool {
	return l.currentPosition != nil
}

func (l *LinearLinkedListIterator) Next() *Node {
	result := l.currentPosition
	l.currentPosition = l.currentPosition.next
	return result
}

// iterator that goes in descending order
type DescendingLinkedListIterator struct {
	currentPosition *Node
	list            *linkedList
}

func NewDescendingLinkedListIterator(l *linkedList) (*DescendingLinkedListIterator, error) {
	if l == nil {
		return nil, errors.New("Invalid linked list")
	}

	// moving to the last element
	cur := l.node
	for cur.next != nil {
		cur = cur.next
	}

	return &DescendingLinkedListIterator{list: l, currentPosition: cur}, nil
}

func (l *DescendingLinkedListIterator) Iterate() bool {
	return l.currentPosition != nil
}

func (l *DescendingLinkedListIterator) Next() *Node {
	cur := l.list.node
	var prev *Node
	for cur != l.currentPosition {
		prev = cur
		cur = cur.next
	}

	l.currentPosition = prev
	return cur
}
