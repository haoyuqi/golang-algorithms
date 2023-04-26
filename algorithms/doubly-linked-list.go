package algorithms

import "fmt"

type DoublyNode struct {
	Value    int
	Previous *DoublyNode
	Next     *DoublyNode
}

var doublyHead = new(DoublyNode)

func doublyAddNode(t *DoublyNode, v int) int {
	if doublyHead == nil {
		t = &DoublyNode{v, nil, nil}
		doublyHead = t
		return 0
	}

	if v == t.Value {
		fmt.Println("The node already exists: ", v)
		return -1
	}

	if t.Next == nil {
		temp := t
		t.Next = &DoublyNode{v, temp, nil}
		return -2
	}

	return doublyAddNode(t.Next, v)
}

func doublyTraverse(t *DoublyNode) {
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

func doublyReverse(t *DoublyNode) {
	if t == nil {
		fmt.Println("-> empty!")
		return
	}

	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}
	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

func doublySize(t *DoublyNode) int {
	if t == nil {
		fmt.Println("-> empty!")
		return 0
	}

	n := 0
	for t != nil {
		n++
		t = t.Next
	}

	return n
}

func doublyLookupNode(t *DoublyNode, v int) bool {
	if doublyHead == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return doublyLookupNode(t.Next, v)
}

func DoublyLinkedList() {
	fmt.Println(doublyHead)

	doublyHead = nil
	doublyTraverse(doublyHead)

	doublyAddNode(doublyHead, 1)
	doublyTraverse(doublyHead)

	doublyAddNode(doublyHead, 5)
	doublyAddNode(doublyHead, 10)
	doublyAddNode(doublyHead, 20)
	doublyTraverse(doublyHead)

	doublyAddNode(doublyHead, 20)
	fmt.Println("size", doublySize(doublyHead))

	doublyReverse(doublyHead)

	if doublyLookupNode(doublyHead, 5) {
		fmt.Println("Already exist")
	} else {
		fmt.Println("Not exist")
	}
}
