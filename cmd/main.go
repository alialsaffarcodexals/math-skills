package main

import (
	"fmt"
	ms "ms/funcs"
	"os"
)

func main() {

args := os.Args
if len(args) != 2 {
	fmt.Println("Invalid args number!")
	return
}

	a := os.Args
	f := a[1]
	fs := string(f)
	ai := ms.DtoArr(fs)
	if len(ai) == 0 {
		fmt.Println("Invalid file content")
		return
	}
	fmt.Println("Average:",int(ms.Avg(ai)))
	fmt.Println("Median:",ms.Med(ai))
	fmt.Println("Variance:",ms.Var(ms.Avg(ai), ai))
	fmt.Println("Standard Deviation:",ms.SD(ms.Var(ms.Avg(ai), ai)))

}
