package main

import "fmt"

func main() {
	list := LinkedList{}
	list.Insert(10)
	list.Insert(20)
	list.InsertAtBegin(200)
	list.InsertNodeAfer(10, 15)
	fmt.Println("Print Array")
	res := list.ToArray()
	fmt.Println(res)
	fmt.Println("Print Reverse Array")
	res = list.ToReverseArray()
	fmt.Println(res)
}

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (list *LinkedList) Insert(data int) {
	node := &Node{data: data, next: nil, prev: nil}

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.prev = list.tail
		list.tail.next = node
		list.tail = node
	}
}

func (list *LinkedList) InsertAtBegin(data int) {
	node := &Node{data: data, next: list.head, prev: nil}
	head := list.head
	head.prev = node
	list.head = &Node{data: data, next: head, prev: nil}
}

func (list *LinkedList) InsertNodeAfer(value int, data int) {
	crr := list.head

	for crr != nil {
		if crr.data == value {
			newNode := &Node{data: data, next: crr.next, prev: crr}
			crr.next.prev = newNode
			crr.next = newNode
			break
		}
		crr = crr.next
	}
}

func (list *LinkedList) ToArray() []int {
	crr := list.head
	if crr == nil {
		return []int{}
	}
	arr := []int{}
	for crr != nil {
		arr = append(arr, crr.data)
		crr = crr.next
	}
	return arr
}

func (list *LinkedList) ToReverseArray() []int {
	crr := list.tail
	if crr == nil {
		return []int{}
	}
	arr := []int{}
	for crr != nil {
		arr = append(arr, crr.data)
		crr = crr.prev
	}
	return arr
}

func ArrayToLinkedList(arr []int) *LinkedList {
	var head, tail *Node

	for _, val := range arr {
		newNode := &Node{data: val, next: nil, prev: nil}

		if head == nil {
			head = newNode
			tail = newNode
		} else {
			newNode.prev = tail
			tail.next = newNode
			tail = newNode
		}
	}

	return &LinkedList{head: head, tail: tail}
}

func (list *LinkedList) DeleteNode(value int) *Node {
	if list.head != nil && list.head.data == value {
		return list.head
	}

	crr := list.head
	var prev *Node

	for crr != nil && crr.data != value {
		prev = crr
		crr = crr.next
	}

	if crr != nil {
		prev.next = crr.next
	}

	return list.head
}
