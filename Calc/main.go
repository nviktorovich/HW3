package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	Sum = "+"
	Subtraction = "-"
	Multiplication = "*"
	Division = "/"
)

func getOperands() (operand1, operand2 float32, err error)  {

	fmt.Println("Введите первый операнд")

	if _, err = fmt.Scanf("%v", &operand1); err != nil {
		err = errors.New("недопустимое значение первого операнда")
		return
	}

	fmt.Println("Введите второй операнд")
	if _, err = fmt.Scanf("%v", &operand2); err != nil {
		err = errors.New("недопустимое значение второго операнда")
		return
	}

	return
}

func getOperator() (operator string, err error)  {
	fmt.Println("Введите математическую операцию")
	fmt.Scan(&operator) // не проверяю здесь ошибку, т.к любой ввод строка, проверю конкретное содержимое в блоке switch
	switch operator {
	case Sum, Subtraction, Multiplication, Division:
		return
	default:
		err = errors.New("only +, -, *, / may be used")
		return
	}
}

func getResult(operand1, operand2 float32, operator string) (result float32, err error) {
	switch operator {
	case Sum:
		result = operand1 + operand2
	case Subtraction:
		result = operand1 - operand2
	case Multiplication:
		result = operand1 * operand2
	case Division:
		if operand2 == 0 { // проверка деления на ноль
			err = errors.New("division by zero")
			return
		}
		result = operand1 / operand2
	}
	return
}

func main() {

	operand1, operand2, err := getOperands()
	if err != nil {
		log.Printf("operand error, %s\n", err)
		os.Exit(1)
	}

	operator, err := getOperator()
	if err != nil {
		log.Printf("operator error, %s\n", err)
		os.Exit(1)
	}

	res, err := getResult(operand1, operand2, operator)
	if err != nil {
		log.Printf("result error, %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v %v %v = %v", operand1, operator, operand2, res)
}
