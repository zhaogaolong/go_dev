package goroute

func Add(a int, b int, pip chan int){
	pip <- a + b
}
