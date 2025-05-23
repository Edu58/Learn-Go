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

func (tree *BinaryTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	treeNode := tree.rootNode

	if treeNode == nil {
		return nil
	}

	for {
		if treeNode.leftNode == nil {
			return &treeNode.key
		}
		treeNode = treeNode.leftNode
	}
}

func (tree *BinaryTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	treeNode := tree.rootNode

	if treeNode == nil {
		return nil
	}

	for {
		if treeNode.rightNode == nil {
			return &treeNode.key
		}
		treeNode = treeNode.rightNode
	}
}

func (tree *BinaryTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	node := tree.rootNode

	if node == nil {
		return false
	}

	for {
		if node.key == key {
			return true
		}

		fmt.Printf("CHECKING %d < %d\n", node.key, key)

		if key < node.key {
			if node.leftNode != nil {
				node = node.leftNode
			}
		} else {
			if node.rightNode != nil {
				node = node.rightNode
			}
		}

		return false
	}
}

func (tree *BinaryTree) RemoveNode(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	removeNode(tree.rootNode, key)
	tree.rootNode.inOrderTraversal()
}

func removeNode(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.leftNode = removeNode(node.leftNode, key)
		return node
	}

	if key > node.key {
		node.rightNode = removeNode(node.rightNode, key)
		return node
	}

	if node.leftNode == nil && node.rightNode == nil {
		node = nil
		return nil
	}

	if node.leftNode == nil {
		node = node.rightNode
		return node
	}

	if node.rightNode == nil {
		node = node.leftNode
		return node
	}

	// Find a replacement since node has left and right child nodes
	leftMostRightNode := node.rightNode
	for {
		//find smallest value on the right side
		if leftMostRightNode != nil && leftMostRightNode.leftNode != nil {
			leftMostRightNode = leftMostRightNode.leftNode
		} else {
			break
		}
	}

	node.rightNode = removeNode(node.rightNode, leftMostRightNode.key)
	return node
}
