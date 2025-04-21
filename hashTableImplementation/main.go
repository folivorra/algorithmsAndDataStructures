package main

import "fmt"

func main() {
	ht := NewHashTable(10)

	ht.Put(123, "hello")
	ht.Put(12, "world")
	ht.Put(90, "!")
	ht.Put(52, "Hello, World!")

	res, err := ht.Get(12)
	if err != nil {
		fmt.Println("error =", err)
	} else {
		fmt.Println("key = 12, value =", res)
	}

	res, err = ht.Get(123)
	if err != nil {
		fmt.Println("error =", err)
	} else {
		fmt.Println("key = 123, value =", res)
	}

	err = ht.Delete(12)
	if err != nil {
		fmt.Println(err)
	}

	res, err = ht.Get(12)
	if err != nil {
		fmt.Println("error =", err)
	} else {
		fmt.Println("key = 123, value =", res)
	}

	err = ht.Update(123, "olleh")
	if err != nil {
		fmt.Println(err)
	}

	res, err = ht.Get(123)
	if err != nil {
		fmt.Println("error =", err)
	} else {
		fmt.Println("key = 123, value =", res)
	}
}
