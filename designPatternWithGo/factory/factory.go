package main

import "fmt"

// 피자를 만드는 Factory pattern

type iPizza interface {
	setName(name string)
	setPrice(price int)
	getName() string
	getPrice() int
}

type Pizza struct {
	name  string
	price int
}

func (c *Pizza) setName(name string) {
	c.name = name
}

func (c *Pizza) getName() string {
	return c.name
}

func (c *Pizza) setPrice(price int) {
	c.price = price
}

func (c *Pizza) getPrice() int {
	return c.price
}

type CheesePizza struct {
	Pizza
}

func newCheesePizza() iPizza {
	return &CheesePizza{
		Pizza{
			name:  "cheese",
			price: 100,
		},
	}
}

type MushroomPizza struct {
	Pizza
}

func newMushroomPizza() iPizza {
	return &MushroomPizza{
		Pizza{
			name:  "mushroom",
			price: 150,
		},
	}
}

func order(pizza string) (iPizza, error) {
	if pizza == "cheese" {
		return newCheesePizza(), nil
	}
	if pizza == "mushroom" {
		return newMushroomPizza(), nil
	}
	return nil, fmt.Errorf("There is no pizza like that")
}

func main() {
	cheesePizza, _ := order("cheese")
	printDetails(cheesePizza)
}

func printDetails(pizza iPizza) {
	fmt.Printf("Pizza: %s", pizza.getName())
}
