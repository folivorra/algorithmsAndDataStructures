package main

import (
	btree "algoAndDS/BTreeImplementation"
	"fmt"
)

func main() {
	tree := &btree.BTree{}

	tree.Insert(65, "A")
	tree.Insert(66, "B")
	tree.Insert(80, "P")
	tree.Insert(69, "E")
	tree.Insert(67, "C")
	tree.Insert(71, "G")
	tree.Insert(68, "D")
	tree.Insert(56, "a")
	tree.Insert(77, "F")
	tree.Insert(81, "F")
	tree.Insert(60, "b")
	tree.Insert(99, "F")
	tree.Insert(84, "F")
	tree.Insert(59, "c")
	tree.Insert(85, "F")
	val, err := tree.SearchRange(50, 90)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}
