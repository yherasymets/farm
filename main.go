package main

import "fmt"

type Animal interface {
	getName() string
	getWeight() float32
	getAvgFood() float32
}

const (
	dogsEatsPerKg float32 = 2
	catEatsPerKg  float32 = 7
	cowEatsPerKg  float32 = 25
)

type Cat struct {
	name   string
	weight float32
}

func (c Cat) getName() string {
	return c.name
}

func (c Cat) getWeight() float32 {
	return c.weight
}

func (c Cat) getAvgFood() float32 {
	return c.weight * catEatsPerKg
}

type Dog struct {
	name   string
	weight float32
}

func (d Dog) getName() string {
	return d.name
}

func (d Dog) getWeight() float32 {
	return d.weight
}

func (d Dog) getAvgFood() float32 {
	return d.weight * catEatsPerKg
}

type Cow struct {
	name   string
	weight float32
}

func (c Cow) getName() string {
	return c.name
}

func (c Cow) getWeight() float32 {
	return c.weight
}

func (c Cow) getAvgFood() float32 {
	return c.weight * catEatsPerKg
}

func getInfo(animals []Animal) float32 {
	var result float32
	for _, animal := range animals {
		result += animal.getAvgFood()
		fmt.Printf("%s weight %.0f kg needs %.0f kg food per month\n", animal.getName(), animal.getWeight(), animal.getAvgFood())
	}
	return result
}

func main() {
	farm := []Animal{
		&Cat{
			name:   "Cat Boris",
			weight: 4.5},
		&Cat{
			name:   "Cat Venya",
			weight: 6,
		},
		&Cat{
			name:   "Cat Brunya",
			weight: 2.2,
		},
		&Dog{
			name:   "Dog Barbos",
			weight: 15,
		},
		&Dog{
			name:   "Dog Matros",
			weight: 12,
		},
		&Cow{
			name:   "Cow Milka",
			weight: 400,
		},
		&Cow{
			name:   "Cow Roshen",
			weight: 390,
		},
		&Cow{
			name:   "Cow Snickers",
			weight: 405,
		},
	}
	sum := getInfo(farm)
	fmt.Println("Total foods for all animals:", sum, "kg")
}
