package main

import "fmt"

type Operation func(float64, float64) float64

//operator
func calculator(op string) Operation {
	switch op {
	case "+": //addition
		return func(x, y float64) float64 { return x + y }
	case "-": //substraction
		return func(x, y float64) float64 { return x - y }
	case "*": //multiplication
		return func(x, y float64) float64 { return x * y }
	case "/": //divison
		return func(x, y float64) float64 {
			if y == 0 {
				panic("mistake!")
			}
			return x / y
		}
	default:
		panic("error,不支持的字符")
	}

}

func main() {
	var x, y float64
	var op string
	fmt.Println("Please input two numbers") //if 除数为0怎么办？
	fmt.Scan(&x, &y)
	fmt.Println("Please input the operator")
	fmt.Scan(&op)
	result := calculator(op)(x, y)
	fmt.Printf("%.3f %s %.3f = %.3f", x, op, y, result)

}
