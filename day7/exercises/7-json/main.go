package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:nickname`
	Age      int
	Birthday string
	Sex      string
	Email    string
}

func jsonStruct() (res string, err error) {
	var user1 = User{
		UserName: "user1",
		NickName: "Conny",
		Age:      18,
		Birthday: "other",
		Sex:      "man",
		Email:    "test@mail.com",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("json format failed")
		return
	}
	// fmt.Printf("%s\n", string(data))
	res = string(data)
	return

}

func jsonInt() {
	var age = 100
	data, err := json.Marshal(age)
	if err != nil {
		fmt.Println("json format failed")
		return
	}
	fmt.Printf("%s\n", string(data))
}

func jsonMap() (res string, err error) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "stud01"
	m["age"] = 110
	m["home"] = "London"
	m["from"] = "Bijing"

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json format failed")
		return
	}
	// fmt.Printf("%s\n", string(data))
	res = string(data)
	return
}

func jsonSlice() (res string, err error) {
	var s []map[string]interface{}
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "stud01"
	m["age"] = 110
	m["home"] = "London"
	m["from"] = "Bijing"

	s = append(s, m)

	var n map[string]interface{}
	n = make(map[string]interface{})
	n["username"] = "stud02"
	n["age"] = 18
	n["home"] = "London"
	n["from"] = "Bijing"
	s = append(s, n)

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json format failed")
		return
	}
	//fmt.Printf("%s\n", string(data))
	res = string(data)
	return

}
func testSlice() {

	data, err := jsonSlice()
	if err != nil {
		fmt.Printf("test json fail")
		return
	}
	var m []map[string]interface{}
	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Printf("test json fail")
	}
	fmt.Println(m)
}
func testMap() {
	data, err := jsonMap()
	if err != nil {
		fmt.Printf("test json fail")
		return
	}
	var m map[string]interface{}
	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Printf("test json fail")
		return
	}
	fmt.Println(m)

}

func testStruct() {
	// json转换为结构体
	data, err := jsonStruct()
	if err != nil {
		fmt.Printf("test json fail")
		return
	}
	var user1 User
	err = json.Unmarshal([]byte(data), &user1)
	if err != nil {
		fmt.Printf("test json fail")
		return
	}
	fmt.Println(user1)

}

func main() {

	//jsonInt()
	// jsonMap()
	// jsonSlice()
	// testStruct()
	// testMap()
	testSlice()
}
