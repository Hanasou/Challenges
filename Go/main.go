package main

import (
	"fmt"
	"sandbox/linkedlist"
)

func main() {
	list := &linkedlist.SingleNode{
		Val:  0,
		Next: nil,
	}
	list.InsertFirst(1)
	list.InsertFirst(2)
	fmt.Println(list)

	list.PointerTest()
	fmt.Println(list)
}
