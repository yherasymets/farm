package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

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

var (
	minCatWeight float32 = 1
	minDogWeight float32 = 2
	minCowWeight float32 = 100
	maxCatWeight float32 = 7
	maxDogWeight float32 = 40
	maxCowWeight float32 = 800
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

func getInfo(animals []Animal) (float32, error) {
	var result float32
	for _, animal := range animals {
		err := validate(animal)
		if errors.Is(err, errors.New("can't eat")) {
			fmt.Printf("for %s: validation error: %v", animal.getName(), err)
		} else if err != nil {
			return 0, fmt.Errorf("for %s: fatal validation error: %v", animal.getName(), err)
		}
		result += animal.getAvgFood()
		fmt.Printf("%s weight %.0f kg needs %.0f kg food per month\n", animal.getName(), animal.getWeight(), animal.getAvgFood())
	}
	return result, nil
}

func validate(animal Animal) error {
	switch animal.(type) {
	case Cow:
		if !strings.Contains(strings.ToLower(animal.getName()), "cow") {
			return errors.New("not a cow")
		}
		if animal.getWeight() < minCowWeight || animal.getWeight() > maxCowWeight {
			return errors.New("cows weight too light")
		}
		return nil
	case Dog:
		if !strings.Contains(strings.ToLower(animal.getName()), "dog") {
			return errors.New("not a dog")
		}
		if animal.getWeight() < minDogWeight || animal.getWeight() > maxDogWeight {
			return errors.New("dogs weight too light")
		}
		return errors.New("can't eat")
	case Cat:
		if !strings.Contains(strings.ToLower(animal.getName()), "cat") {
			return errors.New("not a cat")
		}
		if animal.getWeight() < minCatWeight || animal.getWeight() > maxCatWeight {
			return errors.New("cats weight too light")
		}
		return errors.New("can't eat")
	default:
		return errors.New("unknown zvir")
	}
}

func main() {
	farm := []Animal{
		// Cat{
		// 	name:   "Cat Boris",
		// 	weight: 4.5,
		// },
		// Cat{
		// 	name:   "Cat Venya",
		// 	weight: 6,
		// },
		// Cat{
		// 	name:   "Cat Brunya",
		// 	weight: 2.2,
		// },
		Dog{
			name:   "Dog Barbos",
			weight: 15,
		},
		Dog{
			name:   "Dog Matros",
			weight: 12,
		},
		Cow{
			name:   "Cow Milka",
			weight: 400,
		},
		Cow{
			name:   "Cow Roshen",
			weight: 390,
		},
		Cow{
			name:   "Cow Snickers",
			weight: 405,
		},
	}
	sum, err := getInfo(farm)
	if err != nil {
		log.Fatalf("get info failed: %v", err)
	}
	fmt.Println("Total foods for all animals:", sum, "kg")
}
