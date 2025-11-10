package main

import "fmt"

func main() {

	var subjects []string
	subjects = make([]string, 3)
	subjects[0] = "go"
	subjects[2] = "python"

	for _, susubject := range subjects {
		fmt.Println(susubject)

	}

}
