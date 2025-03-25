package main

import (
	btree "algoAndDS/BTreeImplementation"
	"fmt"
)

func main() {
	tree := &btree.BTree{}

	tree.Insert(10, "Value 10")
	tree.Insert(20, "Value 20")
	tree.Insert(5, "Value 5")
	tree.Insert(6, "Value 6")
	tree.Insert(12, "Value 12")
	tree.Insert(30, "Value 30")
	tree.Insert(7, "Value 7")
	tree.Insert(17, "Value 17")
	tmp, _ := tree.Search(12)
	fmt.Println(tmp)
	_ = tree.Update(12, "Value 52")
	tmp, _ = tree.Search(12)
	fmt.Println(tmp)
	tree.Delete(6)
	tree.Delete(7)
	tree.Delete(12)
	tree.Delete(10)
}
