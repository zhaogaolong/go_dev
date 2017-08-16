package main

import (
	"fmt"
	"os"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	var str = "stu01 18 19.23"

	var stu Student
	fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)
	fmt.Fprintf(os.Stdout, "ddddddddddddd\n")
}
