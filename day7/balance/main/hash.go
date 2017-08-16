package main

import (
	"fmt"
	"go_dev/day7/balance/balance"
	"hash/crc32"
	"math/rand"
)

type HashBalance struct {
}

func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {

	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = fmt.Sprintf("%d", rand.Int())
	}
	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("No backend instance")
		return
	}

	crcTable := crc32.MakeTable(crc32.IEEE)
	// fmt.Println(crcTable)

	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	fmt.Println(hashVal)
	index := int(hashVal) % lens
	fmt.Println("index -->", index)
	inst = insts[index]
	return
}
