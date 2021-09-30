package main

import (
	"fmt"
	"os"
)

func isPrime(num int) bool  {
	for i := 2; i < num / 2 + 1; i++ {
		if num % i == 0 {
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
	for i := 2; i <= userNum; i++ {
		if isPrime(i) {
			fmt.Printf("Число %v является простым\n", i)
		}
	}
}
