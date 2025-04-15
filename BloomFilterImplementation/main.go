package main

import "fmt"

func main() {
	bf := NewBloomFilter(100, 3)

	bf.Add("hello")
	bf.Add("world")
	bf.Add("!")

	if bf.Contains("hello") {
		fmt.Println("true")
	}
	if bf.Contains("world") {
		fmt.Println("true")
	}
	if bf.Contains("!") {
		fmt.Println("true")
	}
	if bf.Contains("hello world!") {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
