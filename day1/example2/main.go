package main

import "fmt"

func main() {
	arr := [8]string{"aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh"}

	sli := arr[:4]

	fmt.Println("sli len:", len(sli), cap(sli), cap(arr))
	fmt.Printf("sli type:%T\n", sli)
	fmt.Printf("arr data:%v\n", arr)
	fmt.Printf("sli data:%v\n", sli)
	sli = append(sli, "aa", "bb", "cc", "dd", "aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj", "kk")
	sli[0] = "11"

	fmt.Println("sli len:", len(sli), cap(sli), cap(arr))
	fmt.Printf("arr data:%v\n", arr)

}
