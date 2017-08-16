package model

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type Book struct {
	Name            string
	Total           int
	Surplus         int
	Author          string
	PublicationTime string
	Info            map[string]int
	// Next            *Book
	BorrowCount int
}

type Books []*Book

var BooksSlice Books

func (b Books) Len() int {
	return len(b)
}

func (b Books) Less(i, j int) bool {
	return b[i].BorrowCount > b[j].BorrowCount
}

func (b Books) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b Book) save() {

	file, err := db("book", "write")
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	defer file.Close()
	for _, book := range BooksSlice {
		data, _ := json.Marshal(book)
		// fmt.Println(book)

		file.WriteString(string(data))
		file.WriteString("\n")
	}

}

func init() {

	// initBook()
	file, err := db("book", "read")
	if err != nil {
		err = fmt.Errorf("Cant found book.db file")
		return
	}
	defer file.Close()
	buff := bufio.NewReader(file)
	for {
		data, err := buff.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			err = fmt.Errorf("Cant found book.db file")
			return
		}
		loadBook(data)
	}
	// fmt.Println(BooksSlice)

}

// 字符串json 加载到切片中
func loadBook(str string) {
	var book Book
	err := json.Unmarshal([]byte(str), &book)
	if err != nil {
		// logging
		return
	}

	if book.Info == nil {
		book.Info = make(map[string]int)
	}

	BooksSlice = append(BooksSlice, &book)
}

func initBook() {
	file, _ := os.OpenFile("book.db", os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	for i := 0; i < 10; i++ {
		var book = Book{
			Name:            fmt.Sprintf("book%d", i),
			Total:           10,
			Surplus:         10,
			Author:          "testName",
			PublicationTime: fmt.Sprintf("201%d-0%d-01", i, i),
		}

		data, err := json.Marshal(book)
		if err != nil {
			fmt.Println("json format failed")
			return
		}
		// fmt.Printf("%s\n", string(data))
		// n, _ := file.Seek(0, os.SEEK_END)

		file.WriteString(string(data))
		file.WriteString("\n")

	}
}

// 书的详细信息
func ShowBook() {

}

func hasBook(name string) bool {
	for _, v := range BooksSlice {
		if v.Name == name {
			return true
		}
	}
	return false
}

// 返回该书的地址
func BookPrt(name string, headBook *Book) *Book {
	tmp := headBook

	// 	for tmp != nil {
	// 		if tmp.Name == name {
	// 			return tmp
	// 		}
	// 		tmp = tmp.Next
	// 	}

	return tmp
}

//是否能借书
func CanBorrow(name string, count int) bool {
	for _, v := range BooksSlice {
		if v.Name == name {
			return v.Surplus >= count
		}
	}
	return false
}

func BookList() {
	fmt.Println("Name\tSurplusCount")

	for i, book := range BooksSlice {
		fmt.Printf("%d\t%s\t%d\n", i, book.Name, book.Surplus)
	}
}

func TopTen() {
	var tmp Books
	for _, book := range BooksSlice {
		tmp = append(tmp, book)
	}
	sort.Sort(tmp)
	fmt.Println("ID\t Name\t BorrowCount")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d \t%s \t%d\n", i, tmp[(i-1)].Name, tmp[(i-1)].BorrowCount)
	}

}

func getBook(name string) *Book {
	var tmpBook *Book
	for _, v := range BooksSlice {
		if v.Name == name {
			tmpBook = v
			break
		}
	}
	return tmpBook
}

func libraryBorrowBook(userName, bookName string, count int) {
	book := getBook(bookName)
	book.BorrowCount += count
	book.Info[userName] = count
	book.Surplus -= count
	book.save()

}

func libraryReturnBook(userName, bookName string, count int) {
	book := getBook(bookName)

	if book.Info[userName] > count {
		book.Info[userName] -= count
		book.Surplus += count
	} else if book.Info[userName] == count {
		delete(book.Info, userName)
		book.Surplus += count
	}
	book.save()
	log.Debugf("return book:%s, count:%d", bookName, count)

}
