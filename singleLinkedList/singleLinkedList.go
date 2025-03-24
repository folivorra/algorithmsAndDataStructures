package singleLinkedList

import (
	"errors"
	"fmt"
	"strconv"
)

type Node struct {
	Next  *Node
	Value int
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

// AddInTail добавляет node в конец списка со значением переданным в аргементе
func (l *LinkedList) AddInTail(Value int) {
	newNode := Node{Value: Value, Next: nil}

	// если список пуст, то присваиваем Head новую node,
	// иначе добавляем в Tail.Next node
	if l.Head == nil {
		l.Head = &newNode
	} else {
		l.Tail.Next = &newNode
	}

	// Tail это наша новая node
	l.Tail = &newNode
	l.Size++
}

// AddInHead добавляет node в начало списка со значением переданным в аргументе
func (l *LinkedList) AddInHead(Value int) {
	newNode := Node{Value: Value, Next: nil}

	// если список пуст, то назначаем Head и Tail = node,
	// иначе помещаем node в начало списка подменой указателей
	if l.Head == nil {
		l.Head, l.Tail = &newNode, &newNode
	} else {
		newNode.Next, l.Head = l.Head, &newNode
	}
	l.Size++
}

// Remove удаляет node со значением переданным в аргументе
func (l *LinkedList) Remove(Value int) error {
	// если список пуст, то удалить нечего
	if l.Head == nil {
		return errors.New("list is empty")
	}

	// если первый элемент списка искомый, то осуществляем нужные замены
	if l.Head.Value == Value {
		l.Head = l.Head.Next
		// если это был единственный элемент списка - обнуляем Tail
		if l.Head == nil {
			l.Tail = nil
		}
		l.Size--
		return nil
	}

	// идем двумя указателями для проверки на соответствие и в случае положительного сравнения прокидывания указателя с prev на Next
	iter := l.Head
	var iterPrev *Node = nil

	for iter != nil {
		if iter.Value == Value {
			iterPrev.Next = iter.Next
			// если удаляемый элемент оказался последним присваиваем Tail = prev
			if iter.Next == nil {
				l.Tail = iterPrev
			}
			iter.Next = nil
			l.Size--
			return nil
		}
		iterPrev = iter
		iter = iter.Next
	}

	// если не нашли элемент возвращаем соответствующую ошибку
	return errors.New("item not found")
}

// Search ищет и возвращает node со значением переданным в аргументе
func (l *LinkedList) Search(Value int) (*Node, error) {
	// если список пустой возвращаем пустой указатель и ошибку
	if l.Head == nil {
		return nil, errors.New("list is empty")
	}

	iter := l.Head
	for iter != nil {
		if iter.Value == Value {
			return iter, nil
		}
		iter = iter.Next
	}

	// если мы не нашли искомый элемент
	return nil, errors.New("item not found")
}

// Print последовательно выводит весь список
func (l *LinkedList) Print() {
	res := ""
	iter := l.Head
	for iter != nil {
		res += strconv.Itoa(iter.Value) + " "
		iter = iter.Next
	}
	fmt.Println(res)
}
