package main

import (
	linkedlist "algorithms/dataStructures/LinkedList"
	"fmt"
)

func main() {
	var list *linkedlist.ListNode
	for i := 1; i < 30; i++ {
		list = list.Append(i)
	}
	fmt.Println("Список после добавления элементов")
	list.Print()
	fmt.Println()
	list = list.Delete(5)
	fmt.Println("После удаления:")
	list.Print()
	fmt.Println()
	list = list.Prepend(100)
	list.Print()
	fmt.Println()
	node := list.Search(15)
	if node != nil {
		fmt.Println("Найдено: ", node.Val)
	} else {
		fmt.Println("Не найдено")
	}

}
