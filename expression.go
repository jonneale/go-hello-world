  
package main

import "regexp"
import "fmt"
import "os"
import "strconv"

func main(){
	expression := os.Args[1:]
	v1 := ""
	v2 := ""
	operator := ""
	for i := 0; i < len(expression); i++ {

		digit, _ := regexp.MatchString("[0-9]", expression[i])
		if digit && operator == "" {
			v1 += expression[i]
		} else if operator == "" {
			operator = expression[i]
		} else {
			v2 = expression[i]
		}
	}

	// fmt.Println("you typed: " + v1 + " " + operator + " " + v2)
	i1, _ := strconv.ParseInt(v1, 10, 8)
	i2, _ := strconv.ParseInt(v2, 10, 8)

	fmt.Println("Result: ", execute(int(i1), int(i2), operator))
}

func add(v1 int, v2 int) int{
	return v1 + v2
}

func multiply(v1 int, v2 int) int{
	return v1 * v2
}

func minus(v1 int, v2 int) int{
	return v1 - v2
}

func divide(v1 int, v2 int) int{
	return v1 / v2
}

func execute(v1 int, v2 int, expression string) int {
	operators := map[string]func(int,int)int{
    "+": add,
    "*": multiply,
    "-": minus,
    "/": divide,
	}
	operator := operators[expression]
	return operator(v1,v2)
}