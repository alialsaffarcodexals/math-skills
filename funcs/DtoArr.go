package ms

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func DtoArr(s string)[]int{
	b, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
		return []int{}
	}
	st := string(b)
	as := strings.Split(st, "\n")
	var ai []int
	for _, v := range as {
		if v == "" {
			continue
		}
		n ,err:= strconv.Atoi(v)
		if err != nil {
			
			return []int{}
		}
		ai = append(ai, n) 
		
	}

	return ai


}