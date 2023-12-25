package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

func (linkedList *LinkedList) IterateList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (linkedList *LinkedList) Tail() *Node {
	var tail *Node
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			tail = node
		}
	}
	return tail
}

func (linkedList LinkedList) AddTail(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil
	tailNode := linkedList.Tail()
	if tailNode != nil {
		tailNode.nextNode = node
	}
}

func main() {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	fmt.Print(linkedList.Tail())
	fmt.Println(linkedList.headNode.property)
}
