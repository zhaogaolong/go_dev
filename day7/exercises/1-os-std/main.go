package main

import (
	"fmt"
)

type student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	// testFile, err := os.OpenFile("os-std.txt", os.O_CREATE|os.O_WRONLY, 0644)
	// defer testFile.Close()
	// st, err := fmt.Fprintf(testFile, "this is a fmt.Fprintf() insterface\n")
	// fmt.Println(st, err)

	var str1 = "stu01 18 90.54"
	var stu student
	fmt.Sscanf(str1, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)
}
