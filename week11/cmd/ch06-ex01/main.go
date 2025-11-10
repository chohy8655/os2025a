package main

import "fmt"

func main() {

	subjects := [4]string{"Go", "javascript", "python", "linux"}
	subjectsSlice := subjects[:3]
	// subjects[0] = "java"
	subjectsSlice[0] = "java"
	subjectsSlice = append(subjectsSlice, "Go", "kotln", "database")
	for _, susubject := range subjects {
		fmt.Println(susubject) //

	}
	fmt.Println("**************")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}

}
