package algorithms

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

var head = new(Node)

func addNode(t *Node, v int) int {
	if head == nil {
		t = &Node{v, nil}
		head = t
		return 0
	}

	if v == t.Value {
		fmt.Println("The node already exists: ", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> empty!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println()
}

func lookupNode(t *Node, v int) bool {
	if head == nil {
		t = &Node{v, nil}
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> empty!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}

	return i
}

func SinglyLinkedList() {
	fmt.Println(head)
	head = nil
	traverse(head)

	addNode(head, 1)
	addNode(head, -1)
	traverse(head)

	addNode(head, 10)
	addNode(head, 5)
	addNode(head, 45)
	traverse(head)

	addNode(head, 5)

	if lookupNode(head, 5) {
		fmt.Println("Already exist")
	} else {
		fmt.Println("Not exist")
	}

	if lookupNode(head, -100) {
		fmt.Println("Already exist")
	} else {
		fmt.Println("Not exist")
	}

	fmt.Println("size: ", size(head))
}
