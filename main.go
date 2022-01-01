package main

import "fmt"


type Student struct{
	Name string;
	Family string;
	Average float64;
}

func Hello(count int,s Student){
for i := 0; i < count; i++ {
	fmt.Printf("Hello %s %s \n",s.Name,s.Family)
}
}
func MultipleReturn(v int) (int ,int){
	return v-1 , v+1
}
func IncreaseAverage(i int,s Student) Student{
	s.Average+=float64(i);
	return s;
}

func main(){
	s :=Student{
		Name: "arsalan",
		Family: "Karimzad",
		Average: 15.4,
	}
	fmt.Println(s)
	Hello(5,s);
	s=IncreaseAverage(1,s)
	fmt.Println(s)

}