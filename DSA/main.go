package main

import (
	"github.com/Edu58/GoogleDev/trees"
)

func main() {
	// var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	// var addedSum int = 18
	// var combinations [19]int

	// lists.FindElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)

	// var linkedList = lists.LinkedList{}
	// linkedList.AddToHead(12)
	// linkedList.Append(13)
	// linkedList.Append(20)
	// linkedList.Append(89)
	// linkedList.Append(9)
	// linkedList.Append(10)
	// linkedList.AddAfter(2, 203)

	// linkedList.Print()

	// var set = &sets.Set{}
	// var anotherSet = &sets.Set{}
	// set.New()
	// anotherSet.New()
	// set.AddEelement(23, true)
	// set.AddEelement(3, false)
	// set.AddEelement(87, true)
	// set.AddEelement(54, true)
	// anotherSet.AddEelement(23, true)
	// anotherSet.AddEelement(30, false)
	// anotherSet.AddEelement(87, true)
	// fmt.Println(set)
	// fmt.Println(anotherSet)
	// fmt.Println(set.Intersection(anotherSet))

	var tree = trees.BinaryTree{}
	tree.InsertElement(4)
	tree.InsertElement(2)
	tree.InsertElement(1)
	tree.InsertElement(6)
	tree.InsertElement(3)
	tree.InsertElement(9)
	tree.InOrderTraversal()
	tree.MinNode()
	tree.MaxNode()
	// search := tree.SearchNode(4)

	// fmt.Printf("SEARCH RESULT: %t\n", inorder)
}
