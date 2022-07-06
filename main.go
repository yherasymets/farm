package main

import "fmt"

type Animal interface {
	getName() string
	getWeight() float32
}

const (
	dogsEatsPerKg float32 = 2 // 10 / 5 = 2
	catEatsPerKg  float32 = 7
	cowEatsPerKg  float32 = 25
)

type Cat struct {
	name   string
	weight float32
}

func (c *Cat) getName() string {
	return c.name
}

func (c *Cat) getWeight() float32 {
	return c.weight
}

type Dog struct {
	name   string
	weight float32
}

func (d *Dog) getName() string {
	return d.name
}

func (d *Dog) getWeight() float32 {
	return d.weight
}

type Cow struct {
	name   string
	weight float32
}

func (c *Cow) getName() string {
	return c.name
}

func (c *Cow) getWeight() float32 {
	return c.weight
}

func getInfo(animals []Animal) float32 {
	var result float32
	for _, animal := range animals {
		switch animal.(type) {
		case *Dog:
			result += animal.getWeight() * dogsEatsPerKg
			fmt.Printf("Dog %s weight %.0f kg needs %.0f kg food per month\n", animal.getName(), animal.getWeight(), animal.getWeight()*dogsEatsPerKg)
		case *Cat:
			result += animal.getWeight() * catEatsPerKg
			fmt.Printf("Cat %s weight %.0f kg needs %.0f kg food per month\n", animal.getName(), animal.getWeight(), animal.getWeight()*catEatsPerKg)
		case *Cow:
			result += animal.getWeight() * cowEatsPerKg
			fmt.Printf("Cow %s weight %.0f kg needs %.0f kg food per month\n", animal.getName(), animal.getWeight(), animal.getWeight()*cowEatsPerKg)
		default:
			fmt.Printf("Unknown animal type %t\n", animal)
		}
	}
	return result
}

func main() {
	farm := []Animal{
		&Cat{
			name:   "Boris",
			weight: 4.5},
		&Cat{
			name:   "Venya",
			weight: 6,
		},
		&Cat{
			name:   "Brunya",
			weight: 2.2,
		},
		&Dog{
			name:   "Barbos",
			weight: 15,
		},
		&Dog{
			name:   "Matros",
			weight: 12,
		},
		&Cow{
			name:   "Milka",
			weight: 400,
		},
		&Cow{
			name:   "Roshen",
			weight: 390,
		},
		&Cow{
			name:   "Snickers",
			weight: 405,
		},
		&Cow{
			name:   "Bounty",
			weight: 350,
		},
		&Cow{
			name:   "Nuts",
			weight: 380,
		},
		&Cow{
			name:   "Kittkat",
			weight: 410,
		},
	}

	sum := getInfo(farm)
	fmt.Println("Total foods for all animals:", sum, "kg")
}
