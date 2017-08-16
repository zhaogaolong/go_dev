package main

import "fmt"

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (p *Link) insertHead(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}

	if p.tail == nil && p.head == nil {
		p.head = node
		p.tail = node
		return
	}

	node.next = p.head
	p.head = node
}

func (p *Link) insertTail(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}

	if p.tail == nil && p.head == nil {
		p.head = node
		p.tail = node
		return
	}

	p.tail.next = node
	p.tail = node

}

func (p *Link) Trans() {
	q := p.head
	for q != nil {
		fmt.Println(q)
		q = q.next
	}
}
