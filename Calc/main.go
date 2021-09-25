package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	SUM = "+"
	SUBTRACTION = "-"
	MULTIPLICATION = "*"
	DIVISION = "/"
)

func getOperands() (float32, float32, error)  {

	var operand1, operand2 float32

	fmt.Println("Введите через пробел два числа")
	_, err := fmt.Scanf("%v %v", &operand1, &operand2)

	if err != nil { // проверка на тип вводимых данных, должны быть числа
		err := errors.New("operand error. Only digits may be used")
		return 0, 0, err
	}
	return operand1, operand2, nil
}

func getOperator() (string, error)  {
	var operator string
	fmt.Println("Введите математическую операцию")
	fmt.Scan(&operator) // не проверяю здесь ошибку, т.к любой ввод строка, проверю конкретное содержимое в блоке switch
	switch operator {
	case "+":
		return SUM, nil
	case "-":
		return SUBTRACTION, nil
	case "*":
		return MULTIPLICATION, nil
	case "/":
		return DIVISION, nil
	default:
		err := errors.New("operator error. Only +, -, *, / may be used")
		return "0", err
	}
}

func getResult(operand1, operand2 float32, operator string) (float32, error) {
	var result float32
	switch operator {
	case SUM:
		result = operand1 + operand2
	case SUBTRACTION:
		result = operand1 - operand2
	case MULTIPLICATION:
		result = operand1 * operand2
	case DIVISION:
		if operand2 == 0 { // проверка деления на ноль
			err := errors.New("error. Division by zero")
			return 0, err
		}
		result = operand1 / operand2
	}
	return result, nil
}

func main() {

	operand1, operand2, err := getOperands()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	operator, err := getOperator()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := getResult(operand1, operand2, operator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%v %v %v = %v", operand1, operator, operand2, res)
}
