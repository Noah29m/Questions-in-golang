package main

import "fmt"

func main() {
	fmt.Println("Quiz Game")
	fmt.Printf("Enter your name")
	var name string

	fmt.Scan(&name)
	fmt.Printf("Hello  %v, welcome to the game \n", name)
	fmt.Printf("Enter your age")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay get in and play! ")
	} else {
		fmt.Println("Gerrarhere!")
		return
	}
	score := 0
	num_questions := 2

	fmt.Printf("what is better, audi or BMW? ")

	var answer string
	fmt.Scan(&answer)

	if answer == "audi" {
		fmt.Println("Correct")
		score++
	} else if answer == "Audi" {
		fmt.Println("correct")
		score++
	} else {
		fmt.Println("oolskia wapi")
	}
	fmt.Printf("which is better, cores or Kotlin?")
	var cores string
	fmt.Scan(&cores)
	if answer == cores {
		fmt.Println("Very true")
		score++
	}
	fmt.Printf("You scored %v out of %v", score, num_questions)
	percent := (float64(score)/float64(num_questions)) *100
	fmt.Printf("You scored: %v%%",percent)
	

}
