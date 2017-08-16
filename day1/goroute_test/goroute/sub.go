package goroute

func Sub(a int, b int, pip chan int){
	pip <- a -b
}