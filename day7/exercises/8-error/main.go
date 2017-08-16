package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path       string
	op         string
	createtime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path%s, op:%s, createtime:%s, message:%s", p.path,
		p.op, p.createtime, p.message)
}

func open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createtime: fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil

}

func main() {
	err := open("e:/dfddsfds/gffd/gfds/gs/gs/gsg/sg/sg")

	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error", v)
	}
}
