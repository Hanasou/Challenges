package main

import (
	"fmt"
	"sandbox/linkedlist"
)

func main() {
	//fmt.Println(queue.WeaveTest())
	n1 := &linkedlist.SingleNode{
		Val:  1,
		Next: nil,
	}

	n1.InsertFirst(3)
	n1.Append(2)

	fmt.Println(n1)
}
