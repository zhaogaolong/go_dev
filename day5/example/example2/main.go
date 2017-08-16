package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	Next  *Student
}


// 单链表
func singleLinkedList(){
	var head Student

	head.Name = "haha"
	head.Age = 20
	head.Score = 80

    var stu1 Student
    stu1.Name = "stu1"
    stu1.Age = 18
    stu1.Score = 20

    head.Next = &stu1

    var stu2 Student
    stu2.Name = "stu2"
    stu2.Age = 13
    stu2.Score = 24
    stu1.Next = &stu2

    var stu3 Student
    stu3.Name = "stu3"
    stu3.Age = 13
    stu3.Score = 24
    stu2.Next = &stu3

    trans(&head)
}


func trans(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.Next
	}
    fmt.Println()

}



// 循环链表
func forInstert(p *Student){
    var tail = p
    for i := 0; i < 10; i++{
        var stu = Student{
            Name: fmt.Sprintf("stud%d", i),
            Age: rand.Intn(100),
            Score: rand.Float32() * 100,
        }
        tail.Next = &stu
        tail = &stu
    }
    // trans(p)

}


func insertHead(p *Student){

    for i := 0; i < 10; i++{
        var stu = Student{
            Name: fmt.Sprintf("stud%d", i),
            Age: rand.Intn(100),
            Score: rand.Float32() * 100,
        }
        stu.Next = p
        p = &stu
    }
    // trans(p)
}

func insertTail(p *Student){
    tail := p
    for i := 0; i < 10; i++{
        var stu = Student{
            Name: fmt.Sprintf("stut%d", i),
            Age: rand.Intn(100),
            Score: rand.Float32() * 100,
        }
        stu.Next = tail
        tail = &stu
    }
}

func delNode(p *Student){
    node := p
    for p.Next != nil{
        if p.Name == "stud6"{
            node.Next = p.Next
            break
        }
        node = p
        p = p.Next
    }

}

func addNode(p *Student, newNode Student){

    for p.Next != nil{
        if p.Name == "stud6"{
            newNode.Next = p.Next
            p.Next = &newNode
            break
        }
        p = p.Next
    }

}

func main(){
    // singleLinkedList()
	var head *Student = new(Student)
	head.Name = "haha"
	head.Age = 20
	head.Score = 80

    forInstert(head)
    // insertHead(head)

    trans(head) // 先打印一下生成的链表

    var newNode = Student{
        Name: "NewNode",
        Age: 18,
        Score: 50,
    }


    addNode(head, newNode)
    trans(head) // 再打印一下删除完的链表
}

