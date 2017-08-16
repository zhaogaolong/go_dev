package model

import "os"
import "fmt"

func db(role, mode string) (file *os.File, err error) {
	if mode == "read" {
		if role == "student" {
			file, err = os.Open("student.db")
			return
		} else if role == "book" {

			file, err = os.Open("book.db")
			return
		}
	} else if mode == "write" {
		if role == "student" {
			file, err = os.OpenFile("student.db", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			return
		} else if role == "book" {
			file, err = os.OpenFile("book.db", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			return
		}
	}
	err = fmt.Errorf("you input pass error")
	return
}
