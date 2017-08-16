package main

import "fmt"

func test(number *int) {
	fmt.Println("number内存地址", number)     //打印number的值（number是一个指针，也是一个变量，存储的是内存地址）
	fmt.Println("number 内存地址的值", *number) //打印number存储的内存地址的值
	fmt.Println("number自己的内存地址", &number) //打印number自己的内存地址
	*number = 100
	return

}

func main() {
	num := 9
	test(&num)
	fmt.Println("num 最后的值：", num)
}
