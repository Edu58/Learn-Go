package lists

import "fmt"

// FindElementsWithSum of k from arr of size
func FindElementsWithSum(arr [10]int, combinations [19]int, size int, k int,
	addValue int, l int, m int) int {
	var num int = 0
	if addValue > k {
		return -1
	}
	if addValue == k {
		num = num + 1
		var p int = 0
		for p = 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}
	var i int
	for i = l; i < size; i++ {
		//fmt.Println(" m", m)
		combinations[m] = l
		FindElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l = l + 1
	}
	return num
}
