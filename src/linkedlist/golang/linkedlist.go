package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type SinglyLinkedList interface {
	AddToHead()
	AddTail()
	Interate()
	Tail() *Node
	IndexOf() int
	Contains() bool
	RemoveAt()
	Removes()
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	node.nextNode = linkedList.headNode
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

func (linkedList *LinkedList) AddTail(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil
	tailNode := linkedList.Tail()
	if tailNode != nil {
		tailNode.nextNode = node
	}
}

func (linkedList *LinkedList) IndexOf(property int) int {
	var node *Node
	index := 0
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			return index
		}
		index = index + 1
	}
	return -1
}

func (linkedList *LinkedList) Contains(property int) bool {
	return linkedList.IndexOf(property) != -1
}

func (linkedList *LinkedList) RemoveAt(index int) {

}

func (linkedList *LinkedList) Removes(property int) {

}

func main() {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToHead(2)
	linkedList.AddTail(5)
	linkedList.AddTail(6)
	linkedList.IterateList()
	fmt.Println("tail", linkedList.Tail())
	fmt.Println("index of ", linkedList.IndexOf(6))
	fmt.Println("Contain", linkedList.Contains(7))
}
