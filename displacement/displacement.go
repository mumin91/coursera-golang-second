package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn generates displacement
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5*a*math.Pow(t, 2) + v0*t + s0)
	}
}

func main() {
	var acceleration, velocity, displacement, time float64

	fmt.Println("Enter acceleration:")
	fmt.Scanln(&acceleration)

	fmt.Println("Enter the initial velocity:")
	fmt.Scanln(&velocity)

	fmt.Println("Enter initial displacement:")
	fmt.Scanln(&displacement)

	fmt.Println("Enter time:")
	fmt.Scanln(&time)

	fn := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Printf("the displacement after %f seconds is %f", time, fn(time))

}
