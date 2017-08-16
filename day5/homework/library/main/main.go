package main

import (
	"bufio"
	"fmt"
	"go_dev/day5/homework/library/model"
	"os"
)

func main() {
	var quit bool
	// model.InitBook()

	// if status {
	// 	printLn(&headStudent, &headBook)

	// }
	for !quit {

		fmt.Println("-----------Welcome Library--------------")
		fmt.Println("T(top10), B(books), R(register), L(login), E(exit)")

		reader := bufio.NewReader(os.Stdin)
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Print("input error", err)
		}
		data := string(input)

		switch data {
		case "Books", "books", "B", "b":
			model.BookList()
		case "Top", "top", "T", "t":
			model.TopTen()
		case "Register", "register", "R", "r":
			model.Register()
			// printLn(&headStudent, &headBook)
		case "Login", "login", "L", "l":
			user, status := model.Login()
			if status {
				model.ComeLibrary(user)
			}
		case "Exit", "exit", "E", "e", "q", "Q":
			fmt.Println("Exit Library")
			quit = true
		default:
			fmt.Println("yous input Wrong!!")
		}
	}

	// close logfile
	if model.LogFileOpened {
		model.LogFile.Close()
	}
}
