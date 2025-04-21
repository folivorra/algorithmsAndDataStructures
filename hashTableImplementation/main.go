package main

import "fmt"

func main() {
	fmt.Println("------------------\nChain method\n------------------")
	var ht HashTable = NewHashTableChain(11)

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

	fmt.Println("------------------\nDouble hash method\n------------------")

	ht = NewHashTableRehash(11)

	ht.Put(123, "hello")
	ht.Put(12, "world")
	ht.Put(90, "!")
	ht.Put(52, "Hello, World!")

	res, err = ht.Get(12)
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
