package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal is an interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow is a type of cow
type Cow struct {
	food, locomotion, noise string
}

// Eat prints animal's food
func (animal Cow) Eat() {
	fmt.Println(animal.food)
}

// Speak prints animal's noise
func (animal Cow) Speak() {
	fmt.Println(animal.noise)
}

// Move prints animal's locomotion
func (animal Cow) Move() {
	fmt.Println(animal.locomotion)
}

// Bird is a type of bird
type Bird struct {
	food, locomotion, noise string
}

// Eat prints animal's food
func (animal Bird) Eat() {
	fmt.Println(animal.food)
}

// Speak prints animal's noise
func (animal Bird) Speak() {
	fmt.Println(animal.noise)
}

// Move prints animal's locomotion
func (animal Bird) Move() {
	fmt.Println(animal.locomotion)
}

// Snake is a type of snake
type Snake struct {
	food, locomotion, noise string
}

// Eat prints animal's food
func (animal Snake) Eat() {
	fmt.Println(animal.food)
}

// Speak prints animal's noise
func (animal Snake) Speak() {
	fmt.Println(animal.noise)
}

// Move prints animal's locomotion
func (animal Snake) Move() {
	fmt.Println(animal.locomotion)
}

func main() {

	data := make(map[string]Animal)

	var userInput string

	for {

		fmt.Println(">")

		reader := bufio.NewReader(os.Stdin)

		userInput, _ = reader.ReadString('\n')

		userInput = strings.TrimSpace(userInput)

		if userInput == "X" {
			break
		}

		req := strings.Split(userInput, " ")

		if req[0] == "newanimal" {

			var m Animal

			switch req[2] {
			case "cow":

				m = Cow{"Grass", "Walk", "moo"}

			case "snake":

				m = Snake{"mice", "slither", "hsss"}

			case "bird":

				m = Bird{"worm", "fly", "peep"}

			default:
				fmt.Println("Invalid Input of animal type")
				continue

			}

			data[req[1]] = m
			fmt.Println("Created it!")

		} else if req[0] == "query" {

			switch req[2] {

			case "eat":

				data[req[1]].Eat()

			case "move":

				data[req[1]].Move()

			case "speak":

				data[req[1]].Speak()

			default:
				fmt.Println("Invalid data for animal")

			}

		}

	}
}
