package singleLinkedList

import (
	"fmt"
	"strconv"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// AddInTail добавляет node в конец списка со значением переданным в аргементе
func (l *LinkedList) AddInTail(value int) {
	newNode := Node{value: value, next: nil}

	// если список пуст, то присваиваем head новую node,
	// иначе добавляем в tail.next node
	if l.head == nil {
		l.head = &newNode
	} else {
		l.tail.next = &newNode
	}

	// tail это наша новая node
	l.tail = &newNode
	l.size++
}

// Remove удаляет node со значением переданным в аргументе
func (l *LinkedList) Remove(value int) {
	// если
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		}
		l.size--
		return
	}

	iter := l.head
	var iterPrev *Node = nil
	for iter != nil {
		if iter.value == value {
			iterPrev.next = iter.next
			if iter.next == nil {
				l.tail = iterPrev
			}
			iter.next = nil
			l.size--
			return
		}
		iterPrev = iter
		iter = iter.next
	}
}

// Search ищет и возвращает node со значением переданным в аргументе
func (l *LinkedList) Search(value int) *Node {
	// если список пустой
	if l.head == nil {
		return nil
	}

	// если список не пустой
	iter := l.head
	for iter != nil {
		if iter.value == value {
			return iter
		}
	}

	// если мы не нашли искомый элемент
	return nil
}

func (l *LinkedList) Print() {
	res := ""
	iter := l.head
	for iter != nil {
		res += strconv.Itoa(iter.value) + " "
		iter = iter.next
	}
	fmt.Println(res)
}
