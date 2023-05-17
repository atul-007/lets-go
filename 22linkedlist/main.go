package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
func (l linkedlist) toprintdata() {
	toprint := l.head
	for l.length != 0 {
		fmt.Printf("%d", toprint.data)
		toprint = toprint.next
		l.length--
		fmt.Print(" ")
	}
	fmt.Println("")
}
func (l *linkedlist) todeletedata(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previoustodelete := l.head
	for previoustodelete.next.data != value {
		if previoustodelete.next.next == nil {
			return
		}
		previoustodelete = previoustodelete.next
	}
	previoustodelete.next = previoustodelete.next.next
	l.length--
}
func main() {
	mylist := linkedlist{}
	node1 := &node{data: 48}
	node2 := &node{data: 834}
	node3 := &node{data: 999}
	node4 := &node{data: 78}
	node5 := &node{data: 9999}
	node6 := &node{data: 124}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.toprintdata()
	mylist.todeletedata(78)
	mylist.todeletedata(100)
	emptylist := linkedlist{}
	emptylist.todeletedata(10)
	mylist.todeletedata(124)

	mylist.toprintdata()
}
