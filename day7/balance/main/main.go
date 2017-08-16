package main

import (
	"fmt"
	"go_dev/day7/balance/balance"
	"os"
	"time"
)

func main() {
	var insts []*balance.Instance

	for i := 0; i < 10; i++ {
		host := fmt.Sprintf("192.168.1.%d", i+10)
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
		fmt.Println(one)
	}
	// blancer := balance.RandomBalance{}
	// blancer := balance.RoundrobinBalance{}
	fmt.Println()

	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Println("do balance err:", err)
			continue
		}

		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
