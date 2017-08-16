package main

func main() {

	var link Link
	for i := 0; i < 10; i++ {
		link.insertHead(i)
	}

	link.Trans()
}
