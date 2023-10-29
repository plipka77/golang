package main

import (
	"errors"
	"fmt"
)

type PersonNode struct {
	name string
	id   string
	next *PersonNode
}

type PersonLinkedList struct {
	head *PersonNode
	tail *PersonNode
	size int
}

func (list *PersonLinkedList) isEmpty() bool {
	return list.size == 0
}

func (list *PersonLinkedList) getSize() int {
	return list.size
}

func (list *PersonLinkedList) append(n PersonNode) {
	if list.head == nil {
		list.head = &n
		list.tail = &n
	} else {
		list.tail.next = &n
		list.tail = &n
	}
	list.size++
}

func (list *PersonLinkedList) insert(n PersonNode, index int) (bool, error) {
	if index < 0 || index >= list.size {
		if index == list.size {
			if list.size == 0 {
				list.head = &n
				list.tail = &n
			} else {
				list.tail.next = &n
				list.tail = &n
			}
		}
		fmt.Println("Index out of bounds")
		return false, errors.New("Index Out of Bounds")
	}

	if list.isEmpty() {
		list.head = &n
		list.tail = &n
		list.size++
		return true, nil
	}

	var prev *PersonNode
	temp := list.head
	for i := 0; i < index; i++ {
		prev = temp
		temp = temp.next
	}

	if prev == nil && temp.next == nil {
		list.head = &n
		list.tail = temp
	} else if prev == nil {
		list.head = &n
	} else if temp.next == nil {
		list.tail = &n
	}

	n.next = temp
	if prev != nil {
		prev.next = &n
	}
	list.size++
	return true, nil
}

func (list *PersonLinkedList) delete(index int) (bool, error) {
	if index < 0 || index >= list.size {
		fmt.Println("Invalid Index")
		return false, errors.New("Index Out of Bounds")
	}

	var prev *PersonNode
	temp := list.head
	for i := 0; i < index; i++ {
		prev = temp
		temp = temp.next
	}

	if temp.next == nil && prev == nil {
		list.tail = nil
		list.head = nil
	} else if temp.next == nil {
		fmt.Println("here")
		list.tail = prev
		if list.size == 2 {
			list.head = prev
		}
		prev.next = temp.next
	} else if prev == nil {
		list.head = temp.next
		if list.size == 2 {
			list.tail = temp.next
		}
	} else {
		prev.next = temp.next
	}

	list.size--
	return true, nil
}

func (list *PersonLinkedList) clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *PersonLinkedList) print() {
	if list.isEmpty() {
		fmt.Println("There are no items in the list")
		return
	}

	result := ""

	temp := list.head
	for temp.next != nil {
		result += ("[name: " + temp.name + ", id: " + temp.id + "] > ")
		temp = temp.next
	}
	result += ("[name: " + temp.name + ", id: " + temp.id + "]")
	fmt.Println(result)
}
