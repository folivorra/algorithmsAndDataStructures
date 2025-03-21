package main

import (
	"algoAndDS/singleLinkedList"
	"fmt"
)

func main() {
	// имитация зацикленного списка
	l := singleLinkedList.LinkedList{}
	l.AddInTail(1)
	l.AddInTail(2)
	l.AddInTail(3)
	l.AddInTail(4)
	l.AddInTail(5)
	l.AddInTail(6)
	l.Tail.Next, _ = l.Search(3)

	fmt.Println(checkLoop(&l))

	l1 := singleLinkedList.LinkedList{}
	l1.AddInTail(1)
	l1.AddInTail(2)
	l1.AddInTail(3)
	l1.AddInTail(4)
	l1.AddInTail(5)
	l1.AddInTail(6)

	fmt.Println(checkLoop(&l1))
}

func checkLoop(l *singleLinkedList.LinkedList) bool {
	// у нас есть быстрый и медленный указатели
	// медленный итерируется по каждому элементу списка, быстрый же итерируется через один элемент
	slow := l.Head
	fast := l.Head
	for {
		slow = slow.Next
		// если не можем перепрыгнуть = значит существует конец списка и мы возвращаем false
		// если можем идем дальше
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		// если fast стал nil значит конец существует и мы возвращаем false
		if fast == nil {
			return false
		}
		// если slow и fast встретились значит гарантировано существует зацикливание и мы возвращаем true
		if fast == slow {
			return true
		}
	}
}

/*

output:
true
false

*/
