package singleLinkedList

import (
	"errors"
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
func (l *LinkedList) Remove(value int) error {
	// если список пуст, то удалить нечего
	if l.head == nil {
		return errors.New("list is empty")
	}

	// если первый элемент списка искомый, то осуществляем нужные замены
	if l.head.value == value {
		l.head = l.head.next
		// если это был единственный элемент списка - обнуляем tail
		if l.head == nil {
			l.tail = nil
		}
		l.size--
		return nil
	}

	// идем двумя указателями для проверки на соответствие и в случае положительного сравнения прокидывания указателя с prev на next
	iter := l.head
	var iterPrev *Node = nil

	for iter != nil {
		if iter.value == value {
			iterPrev.next = iter.next
			// если удаляемый элемент оказался последним присваиваем tail = prev
			if iter.next == nil {
				l.tail = iterPrev
			}
			iter.next = nil
			l.size--
			return nil
		}
		iterPrev = iter
		iter = iter.next
	}

	// если не нашли элемент возвращаем соответствующую ошибку
	return errors.New("item not found")
}

// Search ищет и возвращает node со значением переданным в аргументе
func (l *LinkedList) Search(value int) (*Node, error) {
	// если список пустой возвращаем пустой указатель и ошибку
	if l.head == nil {
		return nil, errors.New("list is empty")
	}

	iter := l.head
	for iter != nil {
		if iter.value == value {
			return iter, nil
		}
	}

	// если мы не нашли искомый элемент
	return nil, errors.New("item not found")
}

// Print последовательно выводит весь список
func (l *LinkedList) Print() {
	res := ""
	iter := l.head
	for iter != nil {
		res += strconv.Itoa(iter.value) + " "
		iter = iter.next
	}
	fmt.Println(res)
}
