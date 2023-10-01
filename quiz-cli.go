package main

import "fmt"

func main() {

	fmt.Println("Welcome to Quiz CLI!")
	fmt.Println("To play this quiz you must be older than 15 Years")
	fmt.Println("There will be two questions which you have to answer")

	// variables
	var name, answer1, answer2 string
	var age, answer, score int

	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Please wait while we check if you are eligble.")

	if age < 15 {
		fmt.Printf("Sorry %v, you are not eligble to play.\n", name)
		return
	}

	fmt.Println("Congraluations, you are eligble. Lets play...")

	fmt.Println("")

	// Question 1
	fmt.Println("Q1. What is greater: 10 or 9")
	fmt.Printf("Answer: ")
	fmt.Scan(&answer)

	// Calucatin answers
	if answer == 10 {
		score++
	}

	// Question 2
	fmt.Println("Q2. Who is known as father of Computers ?")
	fmt.Printf("Answer (use full name): ")
	fmt.Scan(&answer1, &answer2)

	// Calucatin answers
	if answer1+" "+answer2 == "charles babbage" || answer1+" "+answer2 == "Charles Babbage" || answer1+" "+answer2 == "CHARLES BABBAGE" {
		score++
	}

	// Results
	fmt.Printf("Hey %v, you scored %v out of 2\n", name, score)

	fmt.Printf("%v, you answered %v%% Questions \n", name, ((float64)(score) / 2 * 100))

	fmt.Println("Thank you for playing! Quiz ClI")

}
