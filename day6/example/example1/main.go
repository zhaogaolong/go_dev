package main

import "fmt"

type Car interface {
	GetName() string
	Run()
	DiDi()
}

type BMW struct {
	Name string
}

func (p *BMW) GetName() string {
	return p.Name
}

func (p *BMW) Run() {
	fmt.Printf("car %s is run...\n", p.Name)
}

func (p *BMW) DiDi() {
	fmt.Printf("car %d is start didi", p.Name)
}

type BYD struct {
	Name string
}

func (p *BYD) GetName() string {
	return p.Name
}

func (p *BYD) Run() {
	fmt.Printf("car %s is run...\n", p.Name)
}

func (p *BYD) DiDi() {
	fmt.Printf("car %d is start didi", p.Name)
}

// 初识接口
func main() {
	var car Car
	fmt.Println(car)
	bmw := &BMW{
		Name: "BMW",
	}
	car = bmw
	car.Run()

	byd := &BYD{

		Name: "BYD",
	}

	car = byd
	car.Run()

}
