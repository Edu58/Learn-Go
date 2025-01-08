package lists

import "fmt"

type Node struct {
	value    interface{}
	nextNode *Node
}

type LinkedList struct {
	head *Node
}

func (linkedList *LinkedList) AddToHead(value interface{}) {
	var node = Node{value: value}
	if node.nextNode != nil {
		node.nextNode = linkedList.head
	}
	linkedList.head = &node
	return
}

func (linkedList *LinkedList) Append(value interface{}) {
	for val := linkedList.head; val != nil; val = val.nextNode {
		if val.nextNode == nil {
			var nextNode = Node{value: value}
			val.nextNode = &nextNode
			fmt.Printf("Appended Node with value: %v \n", nextNode.value)
			return
		}
	}
	return
}

func (linkedList *LinkedList) LastNode() {
	for val := linkedList.head; val != nil; val = val.nextNode {
		if val.nextNode == nil {
			fmt.Printf("Last Node found: %v \n", val.value)
			return
		}
	}
	return
}

func (linkedList *LinkedList) FindNode(value interface{}) *Node {
	for node := linkedList.head; node != nil; node = node.nextNode {
		if node.value == value {
			fmt.Printf("Found Node with value: %v \n: %v", value, node)
			return node
		}
	}
	fmt.Printf("Node with value: %v NOT found\n", value)
	return nil
}

func (linkedList *LinkedList) AddAfter(value interface{}, after interface{}) {
	var node = linkedList.FindNode(after)
	if node != nil {
		fmt.Printf("Found Node with value: %v \n", after)
		fmt.Println("Adding node")
		var nextNode = Node{value: value, nextNode: node.nextNode}
		node.nextNode = &nextNode
		return
	}
	return
}

func (linkedList *LinkedList) Print() {
	for val := linkedList.head; val != nil; val = val.nextNode {
		fmt.Printf("Value: %v \n", val.value)
	}
	return
}
