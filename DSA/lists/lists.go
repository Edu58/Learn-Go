package lists

import (
	"container/list"
	"fmt"
)

func List() {
	var intList list.List

	intList.PushBack(1)
	intList.PushBack(23)
	intList.PushBack(12)

	for el := intList.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}
}
