package main

import "algoAndDS/singleLinkedList"

func main() {
	l := singleLinkedList.LinkedList{}
	l.AddInTail(1)
	l.AddInTail(2)
	l.AddInTail(3)
	l.AddInTail(4)
	l.AddInTail(5)
	l.AddInTail(6)
	l.Print()

	reverseList(&l)

	l.Print()
}

func reverseList(l *singleLinkedList.LinkedList) {
	// создаем три указателя: next, current и prev
	// изначально prev = nil, current = l.head, а next динамически передвигает на позицию current.next
	current := l.Head
	var next, prev *singleLinkedList.Node
	// заранее ставим tail в новый конец списка
	l.Tail = current
	for current != nil {
		// инициализация next на след. node от current
		next = current.Next
		// разворачиваем две node
		current.Next = prev
		// передвигаем prev и current на 1 позицию по списку
		prev = current
		current = next
	}
	// записываем в head prev так как current стал nil
	l.Head = prev
}

/*

input:
1 2 3 4 5 6

output:
6 5 4 3 2 1

*/
