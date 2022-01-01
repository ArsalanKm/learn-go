package main

import "fmt"


func main ()  {
	var a []int;
	a=make([]int, 10)
	for index, value :=range a{
		fmt.Printf("[%d]:%d\n",index,value)
	}
	

}