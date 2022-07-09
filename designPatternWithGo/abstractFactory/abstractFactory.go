package main

import "fmt"

type iCostmaticCompanyFactory interface {
	makeSkin() iSkin
	makeLotion() iLotion
}

func generateAmoreFactory() iCostmaticCompanyFactory {
	return &Amore{}
}

type Amore struct {
}

func (a *Amore) makeSkin() iSkin {
	return &AmoreSkin{
		Skin{
			price: 37000,
			logo:  "amore",
		},
	}
}

func (a Amore) makeLotion() iLotion {
	lb := NewLotionBuilder()
	return lb.setLogo("amore").
		setPrice(37100).
		getLotion()
}

type iSkin interface {
	getPrice() int
	setPrice(price int)
	getLogo() string
	setLogo(ingredient string)
}

type Skin struct {
	price int
	logo  string
}

func (s *Skin) getPrice() int {
	return s.price
}

func (s *Skin) setPrice(price int) {
	s.price = price
}

func (s *Skin) getLogo() string {
	return s.logo
}

func (s *Skin) setLogo(ingredient string) {
	s.logo = ingredient
}

type AmoreSkin struct {
	Skin
}

type iLotion interface {
	getPrice() int
	getLogo() string
}

type Lotion struct {
	price int
	logo  string
}

type LotionBuilder struct {
	lotion *Lotion
}

func NewLotionBuilder() *LotionBuilder {
	return &LotionBuilder{
		lotion: &Lotion{},
	}
}

func (lb *LotionBuilder) setPrice(price int) *LotionBuilder {
	lb.lotion.price = price
	return lb
}

func (lb *LotionBuilder) setLogo(logo string) *LotionBuilder {
	lb.lotion.logo = logo
	return lb
}

func (lb LotionBuilder) getLotion() *Lotion {
	return lb.lotion
}

func (l *Lotion) getPrice() int {
	return l.price
}

func (l *Lotion) getLogo() string {
	return l.logo
}

type AmoreLotion struct {
	Lotion
}

type Loreal struct {
}

func generateLorealFactory() *Loreal {
	return &Loreal{}
}

func (l *Loreal) makeSkin() iSkin {
	return &LorealSkin{Skin{
		price: 100,
		logo:  "LoReal",
	}}
}

type LorealSkin struct {
	Skin
}

func (l *Loreal) makeLotion() iLotion {
	lb := NewLotionBuilder()
	return lb.setLogo("LoReal").
		setPrice(200).
		getLotion()
}

type LorealLotion struct {
	Lotion
}

func main() {
	amoreFactory := generateAmoreFactory()
	fmt.Println(amoreFactory.makeSkin())
	amoreSkin := amoreFactory.makeSkin()
	fmt.Println("amore skin price:", amoreSkin.getPrice())
	amoreLotion := amoreFactory.makeLotion()
	fmt.Println("amore lotion price:", amoreLotion.getPrice())

	lorealFactory := generateLorealFactory()
	lorealSkin := lorealFactory.makeSkin()
	fmt.Println("loreal skin price:", lorealSkin.getPrice())
	lorealLotion := lorealFactory.makeLotion()
	fmt.Println("loreal lotion price:", lorealLotion.getPrice())
}
