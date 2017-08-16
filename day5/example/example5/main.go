package main


type Car struct{
	weight int
	name string
}


type Bike struct{
	Car
	lunzi int
}

type Train struct{
	c Car
}


func main(){
	var a Bike
	a.Car.name = "bike"
	a.Car.weight = 100
}








k