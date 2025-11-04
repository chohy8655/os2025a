package main

import (
	"fmt"
	"log"
	"week10/pkg/keyboard"
)

func main() {

	fmt.Print("화씨온도입력:")
	fahrenheit, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	var celcius float64

	celcius = (fahrenheit - 32) * 5 / 9
	fmt.Printf("화씨온도%.2f도는 섭씨온드 %.2f도입니다", fahrenheit, celcius)
}
