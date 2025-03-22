package main

import (
	"algoAndDS/singleLinkedList"
	"fmt"
)

func main() {
	l := singleLinkedList.LinkedList{}
	l.AddInTail(1)
	l.AddInTail(2)
	l.AddInTail(3)
	l.AddInTail(4)
	l.AddInTail(5)
	l.AddInTail(6)
	l.AddInTail(7)
	fmt.Println((middleOfList(&l)).Value)
}

func middleOfList(l *singleLinkedList.LinkedList) *singleLinkedList.Node {
	slow := l.Head
	fast := l.Head
	// два указателя: один двигается инкрементировано по списку, второй перепрыгивает через один
	// когда второй дошел до последнего элемента списка заканчиваем цикл и возвращаем первый указатель
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
