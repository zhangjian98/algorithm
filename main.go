package main

import (
	"algorithm/utils"
	"fmt"
	)

func main()  {
	n := 10
	arr := utils.GenerateRandomArray(n,0,n)
	fmt.Print(arr)
}
