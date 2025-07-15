package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
}

type linkedList struct {
	head *node
}

func (l linkedList) String() string {
	sb := strings.Builder{}
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%s ", iterator))
	}
	return sb.String()
}

func (l *linkedList) add(value int) {
	newNode := new(node)
	newNode.value = value

	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		for ; iterator.next != nil; iterator = iterator.next {
		}
		iterator.next = newNode
	}
}

func (l *linkedList) remove(value int) {
	var previousNode *node

	for current := l.head; current != nil; current = current.next {
		if current.value == value {
			if l.head == current {
				l.head = current.next
			} else {
				previousNode.next = current.next
			}
		}
		previousNode = current
	}
}

func (l linkedList) get(value int) (int, bool) {
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return value, true
		}
	}
	return value, false
}

func main() {
	l := new(linkedList)
	l.add(1)
	l.add(2)
	l.add(3)
	l.add(4)
	l.add(5)
	fmt.Println(l)
	fmt.Println(l.get(4))
	fmt.Println(l.get(22))
	l.remove(3)
	fmt.Println(l)

}
