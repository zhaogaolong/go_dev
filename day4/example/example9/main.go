package main

import "fmt"

func testMap() {
	var a map[string]string = map[string]string{
		"key": "value",
	}
	a["ddd"] = "ddd"
	a["ccc"] = "ccc"
	fmt.Println(a)
}

func testMap2() {
	a := make(map[string]map[string]string, 100)
	a["key1"] = make(map[string]string)
	a["key1"]["key1_1"] = "key1_1"
	a["key1"]["key1_2"] = "key1_1"
	a["key1"]["key1_3"] = "key1_1"
	a["key1"]["key1_4"] = "key1_1"
	fmt.Println(a)
}

func modify(a map[string]map[string]string) {
	_, ok := a["zhangsan"]
	if !ok {
		a["zhangsan"] = make(map[string]string)
	}

	a["zhangsan"]["password"] = "123456"
	a["zhangsan"]["age"] = "18"
	fmt.Println(a)
}

func main() {
	//testMap2()
	var a = make(map[string]map[string]string, 100)
	modify(a)
}
