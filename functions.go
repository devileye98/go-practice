package main

import "fmt"

func plus(a int, b int) int {
	return a+b
}

func sumOfThree(a, b, c int) int {
	return a+b+c
}

func main(){
	res:=plus(1,2)
	fmt.Println(res)
	res = sumOfThree(1,2,3)
	fmt.Println(sumOfThree)	
}