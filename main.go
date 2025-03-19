package main

import (
	"algoAndDS/singleLinkedList"
	"fmt"
)

func main() {
	l := singleLinkedList.LinkedList{}
	l.AddInTail(52)
	l.AddInHead(90)
	l.AddInHead(34)
	s, _ := l.Search(90)
	fmt.Println(s)
	l.Print()
	err := l.Remove(34)
	if err != nil {
		fmt.Println(err)
	}
	l.Print()
}
