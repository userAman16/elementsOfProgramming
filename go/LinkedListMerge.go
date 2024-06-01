package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *LinkedList) Add(val int) {
	newNode := &Node{val: val, next: nil}
	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

func (l *LinkedList) Remove(val int) {
	if l.size == 0 {
		return
	}
	if l.head.val == val {
		l.head = l.head.next
	}
	current := l.head
	for current.next != nil {
		if current.next.val == val {
			current.next = current.next.next
			if current.next == nil {
				l.tail = current
			}
			l.size--
			return
		}
		current = current.next
	}
}

func (l *LinkedList) Print() {
	if l.size == 0 {
		return
	}
	current := l.head

	for current != nil {
		fmt.Println(current.val, "->")
		current = current.next
	}
}

func Print(head *Node) {
	for head != nil {
		fmt.Println(head.val, "->")
		head = head.next
	}
}

func LinkedListMerge(ll1 *LinkedList, ll2 *LinkedList) {
	ll3 := &LinkedList{}
	curr1 := ll1.head
	curr2 := ll2.head

	for curr1 != nil && curr2 != nil {
		if curr1.val <= curr2.val {
			ll3.Add(curr1.val)
			curr1 = curr1.next
		} else {
			ll3.Add(curr2.val)
			curr2 = curr2.next
		}
	}

	if curr1 == nil {
		for curr2 != nil {
			ll3.Add(curr2.val)
			curr2 = curr2.next
		}
	}
	if curr2 == nil {
		for curr1 != nil {
			ll3.Add(curr1.val)
			curr1 = curr1.next
		}
	}
	ll3.Print()
}

func LinkedListMergeInPlace(ll1 *LinkedList, ll2 *LinkedList) {
	var head, curr, other *Node
	if ll1.head.val < ll2.head.val {
		head = ll1.head
		curr = ll1.head
		other = ll2.head

	} else {
		head = ll2.head
		curr = ll2.head
		other = ll1.head
	}

	for curr.next != nil && other != nil {
		if curr.next.val >= other.val {
			temp := other
			other = other.next
			temp.next = curr.next
			curr.next = temp
			curr = temp.next
		}
	}

	Print(head)

}

func LinkedListReverse(ll *LinkedList) {
	curr := ll.head
	a, c := curr, curr
	b := curr.next
	for b != nil {
		print(a.val, b.val, c.val)
		temp := b.next
		b.next = a
		c.next = temp

		a = b
		b = temp
		Print(a)
	}

}

func LinkedListFindLoop(ll *LinkedList) {
	loop := true
	pl1, pl2 := ll.head, ll.head
	for {
		if pl1.next == nil {
			loop = false
			break

		}
		pl1 = pl1.next
		pl2 = pl2.next.next
		if pl1 == pl2 {
			loop = true
			break
		}
	}
	if loop {
		cycleLen := 0
		for {
			cycleLen += 1
			pl2 = pl2.next
			if pl1 == pl2 {
				break
			}
		}

		pl1, pl2 = ll.head, ll.head

		for x := 0; x < cycleLen; x++ {
			pl2 = pl2.next
		}

		for pl1 != pl2 {
			pl1 = pl1.next
			pl2 = pl2.next

		}
		fmt.Println("Cycle Start : ", pl1.val)

	} else {
		fmt.Println("no loop")
	}

}

func LinkedListMergeMain() {
	ll1 := &LinkedList{}
	(*ll1).Add(2)
	(*ll1).Add(3)
	(*ll1).Add(5)
	(*ll1).Add(7)
	(*ll1).Add(11)
	ll2 := &LinkedList{}
	(*ll2).Add(3)
	(*ll2).Add(4)
	(*ll2).Add(6)
	(*ll2).Add(9)
	//LinkedListMerge(ll1, ll2)
	//LinkedListMergeInPlace(ll1, ll2)
	//LinkedListReverse(ll1)
	ll3 := &LinkedList{}
	(*ll3).Add(2)
	(*ll3).Add(3)
	(*ll3).Add(5)
	(*ll3).Add(7)
	(*ll3).Add(11)
	(*ll3).Add(112)
	(*ll3).Add(114)
	ll3.tail.next = ll3.head.next.next.next.next
	LinkedListFindLoop(ll3)

}
