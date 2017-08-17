package main

import "fmt"

func logic() {
	enterMenu()
}

func enterMenu() {
	fmt.Println("1. list online users")
	fmt.Println("2. talk")
	fmt.Println("3. exit")
	var sel int
	fmt.Scanf("%d\n", &sel)

	switch sel {
	case 1:
		outputUserOnline()
	case 2:
		// enterTalk()
		fmt.Println("start talk...")
	case 3:
		return
	}
}
