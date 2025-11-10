package main

import "fmt"

func main() {

	subjects := []string{"Go", "javascript", "python", "linux"}
	subjectsSlice := subjects[1:3]
	for _, susubject := range subjects {
		fmt.Println(susubject)

	}
	fmt.Println("**************")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}

}
