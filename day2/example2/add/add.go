package add

import (
	"fmt"
	"go_dev/day2/example2/sub"
)


var Name string
var Age int


func init (){
	fmt.Println("Package add initialied")
}


func Test(){
	Name = "Conny"
	Age = 18
}
