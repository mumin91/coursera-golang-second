package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal type
type Animal struct {
	food, locomotion, noise string
}

// Eat prints the animal food
func (a *Animal) Eat() {
	fmt.Println(a.food)
}

// Move prints the animal locomotion
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak prints the animal noise
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

// makeOutput makes output based on user input
func makeOutput(inputString string) {

}

func main() {

	var userInput string
	reader := bufio.NewReader(os.Stdin)
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {
		fmt.Println(">")

		userInput, _ = reader.ReadString('\n')

		userInput = strings.TrimSpace(userInput)

		if userInput == "x" {
			break
		}

		nameArray := strings.Split(userInput, " ")

		switch nameArray[0] {
		case "cow":
			if nameArray[1] == "eat" {
				cow.Eat()
			} else if nameArray[1] == "move" {
				cow.Move()
			} else if nameArray[1] == "speak" {
				cow.Speak()
			} else {
				fmt.Println("Wrong Input for eat, move or speak")
			}

		case "bird":
			if nameArray[1] == "eat" {
				bird.Eat()
			} else if nameArray[1] == "move" {
				bird.Move()
			} else if nameArray[1] == "speak" {
				bird.Speak()
			} else {
				fmt.Println("Wrong Input for eat, move or speak")
			}

		case "snake":
			if nameArray[1] == "eat" {
				snake.Eat()
			} else if nameArray[1] == "move" {
				snake.Move()
			} else if nameArray[1] == "speak" {
				snake.Speak()
			} else {
				fmt.Println("Wrong Input for eat, move or speak")
			}
		default:
			fmt.Println("Wrong Input for cow, snake or bird")

		}
	}
}
