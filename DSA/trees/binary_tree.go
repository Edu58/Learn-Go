package trees

import (
	"fmt"
	"sync"
)

type Node struct {
	key       int
	leftNode  *Node
	rightNode *Node
}

type BinaryTree struct {
	rootNode *Node
	lock     sync.RWMutex
}

func (tree *BinaryTree) InsertElement(key int) *Node {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	newNode := Node{key: key}

	if tree.rootNode == nil {
		tree.rootNode = &newNode
		return &newNode
	}

	return tree.InsertNode(tree.rootNode, &newNode)
}

func (tree *BinaryTree) InsertNode(currNode *Node, newNode *Node) *Node {
	if newNode.key < currNode.key {
		if currNode.leftNode == nil {
			currNode.leftNode = newNode
			return newNode
		} else {
			tree.InsertNode(currNode.leftNode, newNode)
		}
	} else {
		if currNode.rightNode == nil {
			currNode.rightNode = newNode
			return newNode
		} else {
			tree.InsertNode(currNode.rightNode, newNode)
		}
	}

	return nil
}

// Left, Root, Right
func (tree *BinaryTree) InOrderTraversal() {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	tree.rootNode.inOrderTraversal()
}

func (node *Node) inOrderTraversal() {
	if node.leftNode != nil {
		node.leftNode.inOrderTraversal()
	}
	fmt.Println(node.key)
	if node.rightNode != nil {
		node.rightNode.inOrderTraversal()
	}
}

func (node *Node) preOrderTraversal() {
	fmt.Println(node.key)

	if node.leftNode != nil {
		node.leftNode.preOrderTraversal()
	}
	if node.rightNode != nil {
		node.rightNode.preOrderTraversal()
	}
}

func (node *Node) postOrderTraversal() {
	if node.leftNode != nil {
		node.leftNode.postOrderTraversal()
	}
	if node.rightNode != nil {
		node.rightNode.postOrderTraversal()
	}
	fmt.Println(node.key)
}

// func inOrderTraversal(root *Node) []int {
// 	if root == nil {
// 		return nil
// 	}

// 	left := inOrderTraversal(root.leftNode)
// 	right := inOrderTraversal(root.rightNode)

// 	output := make([]int, 0)

// 	fmt.Println("ASDASD")

// 	output = append(output, left...)
// 	output = append(output, root.key)
// 	output = append(output, right...)
// 	return output
// }
