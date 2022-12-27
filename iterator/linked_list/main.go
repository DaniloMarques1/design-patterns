package main

import "log"

func main() {
	l := NewLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	//l.Pop()
	l.Add(12)
	//l.Pop()
	//l.Print()
	//iterator, err := l.CreateLinkedListIterator(LINEAR)
	iterator, err := l.CreateLinkedListIterator(DESCENDING)
	if err != nil {
		log.Fatal(err)
	}

	for iterator.Iterate() {
		element := iterator.Next()
		log.Printf("This is the element = %v\n", element.data)
	}
}
