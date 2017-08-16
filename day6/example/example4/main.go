package main

type student struct {
	name string
	next *student
}

type test interface {
	Link()
}

func (p *student) Link(a, b *student) {
	a.next = b
}

func main() {

}
