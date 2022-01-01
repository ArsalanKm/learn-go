package main

import "fmt"


type Helloer interface{
	Hello(count int)
}

type H1 interface{
	Hello(count int)
}

type Student struct{
	Name string;
	Family string;
	Average float64;
}

func (s Student) Hello(count int){
for i := 0; i < count; i++ {
	fmt.Printf("Hello %s %s \n",s.Name,s.Family)
}
}
func MultipleReturn(v int) (int ,int){
	return v-1 , v+1
}
func (s Student)IncreaseAverage(i int,) Student{
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
	s.Hello(5);
	s=s.IncreaseAverage(1)
	fmt.Println(s)
	var h Helloer = s
	h.Hello(10)

}