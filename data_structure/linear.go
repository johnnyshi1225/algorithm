//#########################################################################
// Author: Johnny Shi
// Created Time: 2015-06-24 00:38:58
// File Name: linear.go
// Description: 
//#########################################################################
package main

import (
	"fmt"
)

// ========== array ===========
func array() {
	fmt.Println("array")
	var a [10]int
	fmt.Println(a)

	b := []int{1, 3, 5, 7}
	fmt.Println(b)
}
// ========== array end ===========

// ========== link list ===========
type Node struct {
	value int
	next *Node
}

func initList(arr []int) (list *Node){
	header := &Node{}
	// 采用哨兵表头，即固定的空表头，value存放连表长度
	// 创建采用头插法，省去记录尾指针
	for _, item := range arr {
		node := &Node{value:item}
		if header.next == nil {
			header.next = node
		} else {
			node.next = header.next
			header.next = node
		}
		header.value++
	}

	return header
}

func listLen(list *Node) int {
	return list.value
}

func printList(list *Node) {
	var current *Node
	current = list.next
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func insertIntoList(list *Node, value int, index int) (ok bool) {
	ok = true
	i := 1
	cur := list

	for cur != nil && i < index {
		cur = cur.next
		i++
	}

	fmt.Println("i", i)

	if cur == nil {
		ok = false
	} else {
		node := &Node{value:value}
		node.next = cur.next
		cur.next = node
	}

	return ok
}

func linkList() {
	fmt.Println("link list")
	list := initList([]int{1, 3, 5, 7})
	fmt.Println("len of list:", listLen(list))
	printList(list)
	insertIntoList(list, 9, 3)
	printList(list)
}
// ========== link list end ===========

func main() {
	//array()
	linkList()
}


// vim: set noexpandtab ts=4 sts=4 sw=4 :
