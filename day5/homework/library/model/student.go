package model

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Student struct {
	Name        string
	Password    string
	Age         int
	Class       string
	Sex         string
	ID          int
	Books       map[string]int
	BorrowCount int
}

// 结构体
type students []*Student

//  全局变量
var Students students

// save student status
func (s Student) save() {
	file, err := db("student", "write")
	if err != nil {
		return
	}
	defer file.Close()
	for _, stu := range Students {
		data, _ := json.Marshal(stu)
		file.WriteString(string(data))
		file.WriteString("\n")
	}
}

func init() {
	file, err := db("student", "read")
	if err != nil {
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
			fmt.Printf("Cant found students.db file err:%v\n", err)
			return
		}
		loadStudent(data)
	}
	// fmt.Println(Students)
	// for _, stu := range Students {
	// 	fmt.Printf("student %s address:%p\n", stu.Name, &stu)
	// }

	// fmt.Println(Students)
}

func loadStudent(str string) {
	var stu Student
	err := json.Unmarshal([]byte(str), &stu)
	if err != nil {
		// logging
		return
	}

	// init stu.Books map
	if stu.Books == nil {
		stu.Books = make(map[string]int)
	}

	// fmt.Printf("Init student struct %v\n", stu)
	Students = append(Students, &stu)

}

// 添加新的用户到文件中
// func writeNewUser(stu Student) {
// 	file, err := db("student", "write")
// 	if err != nil {
// 		return
// 	}
// 	defer file.Close()
// 	data, err := json.Marshal(stu)
// 	file.WriteString(string(data))
// 	file.WriteString("\n")
// }
// 学生还书
func returnBookFlow(stu *Student) {
	var (
		bookName string
		count    int
	)

	fmt.Println("Please input book name")
	fmt.Scanf("%s\n", &bookName)
	status := studentHasBook(stu, bookName)
	if status {
		fmt.Println("Please input return books count:")
		fmt.Scanf("%d\n", &count)
		status = canReutrnBook(stu, bookName, count)
		if status {
			//success
			studentReturnBook(stu, bookName, count)
		} else {
			fmt.Println("failed, please check book count")
			// failed
		}
		studentReturnBook(stu, bookName, count)
		fmt.Println("Borrow sucess")

	} else {
		fmt.Println("You input book's name not found.")
	}
}

// 学生借书
func BorrowBookFlow(stu *Student) {
	var (
		bookName string
		count    int
	)

	fmt.Println("Please input book name")
	fmt.Scanf("%s\n", &bookName)
	status := hasBook(bookName)
	if status {
		fmt.Println("Please input number:")
		fmt.Scanf("%d\n", &count)
		status = CanBorrow(bookName, count)
		if status {
			studentBorrowBook(stu, bookName, count)
			fmt.Println("Borrow sucess")

		} else {
			fmt.Println("you input count error")
		}

	} else {
		fmt.Println("You input book's name not found.")
	}
}

func Login() (username string, status bool) {
	fmt.Println("-----Welcome Login ----------")
	var (
		password string
	)
	fmt.Println("Please input your name.")
	fmt.Scanf("%s\n", &username)
	status = hasStudent(username)
	if status {
		fmt.Println("Please input your password.")
		fmt.Scanf("%s\n", &password)

		status := checkPassword(username, password)
		if status == true {
			fmt.Printf("---User %s logined\n", username)
			log.Debugf("user %s is logined", username)
			return username, status
		} else {
			fmt.Println("you input password wrong.")
		}

	} else {
		fmt.Printf("Can't Find %s user\n", username)
	}
	return

}

func Register() {
	fmt.Println("-----Welcome Register User -----")
	var (
		username string
		password string
		id       int
		status   bool
		quit     string
	)

	for status != true {
		fmt.Println("Please input your name")
		fmt.Scanf("%s \n", &username)
		check := hasStudent(username)

		if !check {
			fmt.Println("Input your password")
			fmt.Scanf("%s\n", &password)

			fmt.Println("Input your ID")
			fmt.Scanf("%d\n", &id)

			fmt.Println("Please wait a moment, creating user")
			time.Sleep(time.Second)
			createUser(username, password, id)
			fmt.Println("User register success, please login")
			break

		} else {
			fmt.Println("Sorry you input this username is used. please input other username.")
			fmt.Println("if you quit register, please input quit ")
			fmt.Scanf("%s", quit)
			// if quit == "quit" {
			// 	break
			// }
			continue
		}

	}

}

func hasStudent(username string) bool {
	// fmt.Println("input username :", username)
	for _, val := range Students {
		if val.Name == username {
			return true
		}
	}
	return false
}

func createUser(username, password string, id int) {
	newUser := Student{
		Name:     username,
		Password: password,
		ID:       id,
		Books:    make(map[string]int),
	}
	Students = append(Students, &newUser)
	log.Infof("Register user:%s success\n", username)
	newUser.save()
}

func checkPassword(username, password string) (status bool) {
	for _, val := range Students {
		if val.Name == username && val.Password == password {
			return true
		}
	}
	return false
}

func getStudent(username string) *Student {
	var stu *Student
	for _, v := range Students {
		if v.Name == username {
			return v
		}
	}
	fmt.Println("Cant found student struct!!")
	return stu
}

func ComeLibrary(username string) {
	var (
		logout = true
	)
	// reader := bufio.NewReader(os.Stdin)
	// input, _, err := reader.ReadLine()
	// if err != nil {
	// 	fmt.Print("input error", err)
	// }
	// data := string(input)
	userStruct := getStudent(username)
	// fmt.Printf("student %p\n", userStruct)

	for logout {
		fmt.Printf("--------Welcome %s ---------\n", username)
		fmt.Println("T(top10), B(borrow 借书), R(return,还书), L(books list,书列表) C(Book Car), O(logout),")
		var data string
		fmt.Scanf("%s\n", &data)
		// fmt.Println("input ", data)
		switch data {
		case "O", "o", "logout":
			logout = false
		case "B", "b", "borrow":
			BorrowBookFlow(userStruct)
		case "T", "t", "top10":
			TopTen()
		case "R", "r", "return":
			returnBookFlow(userStruct)
		case "L", "l", "list":
			BookList()
		case "c", "C":
			StudentBooksCar(userStruct)
		}
	}
}
func StudentBooksCar(stu *Student) {
	fmt.Printf("----%s Borrow Car---\n", stu.Name)
	fmt.Println("Book\t Count")
	for k, v := range stu.Books {
		fmt.Println(k, "\t", v)
	}

}

func studentBorrowBook(stu *Student, bookName string, count int) {

	if _, ok := stu.Books[bookName]; ok {
		stu.Books[bookName] += count
	} else {
		stu.Books[bookName] = count
	}
	stu.BorrowCount += count

	stu.save()
	log.Infof("user:%s borrow:%s count:%d", stu.Name, bookName, count)
	// fmt.Println(stu)

	libraryBorrowBook(stu.Name, bookName, count)
}

func studentReturnBook(stu *Student, bookName string, count int) {

	if stu.Books[bookName] > count {
		stu.Books[bookName] -= count
	} else if stu.Books[bookName] == count {
		delete(stu.Books, bookName)
	}

	// fmt.Println(stu)
	stu.save()
	log.Infof("user:%s return:%s count:%d", stu.Name, bookName, count)
	// fmt.Println(stu)
	libraryReturnBook(stu.Name, bookName, count)
}

func studentHasBook(stu *Student, bookName string) bool {
	for book := range stu.Books {
		if book == bookName {
			return true
		}
	}
	return false
}

func canReutrnBook(stu *Student, bookName string, count int) bool {
	if stu.Books[bookName] >= count {
		fmt.Println()
		return true
	} else {
		fmt.Println("You can't has %d of %s", count, bookName)
		return false
	}
}
