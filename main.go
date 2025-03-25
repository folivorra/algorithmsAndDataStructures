package main

import (
	btree "algoAndDS/BTreeImplementation"
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
	tree.Insert(56, "F")
	tree.Insert(77, "F")
	tree.Insert(81, "F")
	tree.Insert(60, "F")
	tree.Insert(99, "F")
	tree.Insert(84, "F")
	tree.Insert(59, "F")
	tree.Insert(85, "F")

}
