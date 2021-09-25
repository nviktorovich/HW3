package main

import (
	"fmt"
	"os"
)

func isSimple(num int) bool  {
	var cnt int
	for i := 1; i < num; i++ {
		if num % i == 0 {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return true
}

func main()  {
	var userNum int
	fmt.Println("Введите целое положительное число от 2 и более")
	_, err := fmt.Scan(&userNum)
	if err != nil || userNum < 2 {
		fmt.Println("Необходимо ввести целое положительное число число от 2 и более")
		os.Exit(1)
	}
	for i := 1; i < userNum + 1; i++ {
		if isSimple(i) {
			fmt.Printf("Число %v является простым\n", i)
		}
	}
}
