package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string
	Password string
	Age      int
	Sex      string
}

func testStrcut() {
	user1 := &User{
		Username: "user1",
		Password: "ddd",
		Age:      18,
		Sex:      "boy",
	}
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("Can't convert json")
		return
	}
	fmt.Printf("%s\n", string(data))
}

func testInt() {
	var age = 10
	data, err := json.Marshal(age)
	if err != nil {
		fmt.Println("Can't convert json")
		return
	}
	fmt.Printf("%s\n", string(data))

}
func testMap() {
	var m map[string]interface{}

	m = make(map[string]interface{})
	m["name"] = "user2"
	m["age"] = 12
	m["sex"] = "men"
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Can't convert json")
		return
	}
	fmt.Printf("%s\n", string(data))
}
func main() {
	testStrcut()
	// testInt()
	// testMap()
}
