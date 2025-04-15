package main

import "fmt"

func main() {
	tree := BTree{}

	tree.Insert(65, "a")
	tree.Insert(66, "b")
	tree.Insert(80, "c")
	tree.Insert(69, "d")
	tree.Insert(67, "e")
	tree.Insert(71, "f")
	tree.Insert(68, "g")
	tree.Insert(56, "h")
	tree.Insert(77, "m")
	tree.Insert(81, "l")
	tree.Insert(60, "k")
	tree.Insert(99, "n")
	tree.Insert(84, "o")
	tree.Insert(59, "p")
	tree.Insert(85, "t")

	res, err := tree.Search(80)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	err = tree.Update(80, "C")

	if err != nil {
		fmt.Println(err)
	}

	val, err := tree.SearchRange(80, 90)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	tree.Delete(99)
	tree.Delete(60)
	tree.Delete(59)
	tree.Delete(56)
	tree.Delete(68)
}
