package main

import "fmt"

func main() {
	l := LinkedList{}

	l.Insert("jitendra")
	l.Insert("yashwantrao")
	l.Insert(36)
	l.Print()
	fmt.Println(l.GetAt(2))
	l.InsertAt(1, "inserted at perticular positiion")
	l.Print()
}

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head *Node
	len  int
}

//Insert element
func (l *LinkedList) Insert(val interface{}) {

	n := Node{}
	n.value = val

	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}

	ptr := l.head

	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

//Print list
func (l *LinkedList) Print() {

	if l.len == 0 {
		fmt.Println("No elements are present in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}

//GetAt position
func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}

	if pos > (l.len - 1) {
		return nil
	}

	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

//InsertAt position
func (l *LinkedList) InsertAt(pos int, val interface{}) {

	n := Node{}
	n.value = val

	if pos < 0 {
		return
	}

	if pos == 0 {

		n.next = l.head
		l.head = &n

		return
	}

	if pos > l.len {
		return
	}

	nodeat := l.GetAt(pos)
	nodeatprev := l.GetAt(pos - 1)
	n.next = nodeat
	nodeatprev.next = &n
	l.len++
	return

}
