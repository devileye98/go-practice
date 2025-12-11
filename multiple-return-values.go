package main

import "fmt"

func vals() (int, int) {
	return 1,2
}

func main(){
	a, b:=vals()
	fmt.Println(a)
	fmt.Println(b)

	c,_ := vals()
	fmt.Println(c)
}