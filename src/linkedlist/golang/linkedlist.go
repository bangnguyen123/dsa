package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) addToHead(property int) {
	var node = Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

func main() {
	linkedList := LinkedList{}
	linkedList.addToHead(1)
	linkedList.addToHead(3)
	fmt.Println(linkedList.headNode.property)
}
