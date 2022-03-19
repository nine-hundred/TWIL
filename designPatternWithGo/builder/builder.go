package main

import "fmt"

// 음료를 만드는 Builder pattern

type Beverage struct {
	Brand       string
	Category    string
	Name        string
	ingredients []string
	Bottle      string
}

type BeverageBuilder struct {
	beverage *Beverage
}

func NewBeverageBuilder() *BeverageBuilder {
	return &BeverageBuilder{
		beverage: &Beverage{},
	}
}

func (b *BeverageBuilder) Category(category string) *BeverageBuilder {
	b.beverage.Category = category
	return b
}

func (b *BeverageBuilder) Name(name string) *BeverageBuilder {
	b.beverage.Name = name
	return b
}

func (b *BeverageBuilder) Ingredient(ingredient string) *BeverageBuilder {
	b.beverage.ingredients = append(b.beverage.ingredients, ingredient)
	return b
}

func (b *BeverageBuilder) Ingredients(ingredients ...string) *BeverageBuilder {
	b.beverage.ingredients = ingredients
	return b
}

type BeverageBrandBuilder struct {
	BeverageBuilder
}

func (b *BeverageBuilder) Brand() *BeverageBrandBuilder {
	return &BeverageBrandBuilder{*b}
}

func (b *BeverageBrandBuilder) StarBucks() *BeverageBrandBuilder {
	b.beverage.Brand = "StarBucks"
	return b
}

func (b *BeverageBrandBuilder) Ediya() *BeverageBrandBuilder {
	b.beverage.Brand = "Ediya"
	return b
}

type BeverageBottleBuilder struct {
	BeverageBuilder
}

func (b *BeverageBuilder) Bottle() *BeverageBottleBuilder {
	return &BeverageBottleBuilder{
		*b,
	}
}

func (b *BeverageBottleBuilder) TakeOut() *BeverageBottleBuilder {
	b.beverage.Bottle = "plastic-cup"
	return b
}

func (b *BeverageBottleBuilder) InHere() *BeverageBottleBuilder {
	b.beverage.Bottle = "mug-cup"
	return b
}

func (b *BeverageBottleBuilder) BringTumbler() *BeverageBottleBuilder {
	b.beverage.Bottle = "tumbler"
	return b
}

func (b *BeverageBuilder) Beverage() *Beverage {
	return b.beverage
}

func main() {
	beverageBuilder := NewBeverageBuilder()
	starBucksHotAmericano := beverageBuilder.
		Name("Americano").
		Category("coffe").
		Brand().
		StarBucks().
		Ingredients("hot-water", "espresso").
		Bottle().
		TakeOut().
		Beverage()

	beverageBuilder = NewBeverageBuilder()
	ediyaAppleJuicy := beverageBuilder.
		Brand().
		Ediya().
		Name("appleJuicy").
		Category("juice").
		Ingredients("apple", "ice", "water", "syrup").
		Bottle().
		InHere().
		Beverage()

	fmt.Println(ediyaAppleJuicy)
	fmt.Println(starBucksHotAmericano)
}
