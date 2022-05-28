package main

import "fmt"

type ProgrammingLanguage interface {
	leanHard()
}

type Go struct {
}

func (g *Go) leanHard() {
	fmt.Println("Working hard with Golang")
}

type PHP struct {
}

func (p *PHP) leanHard() {
	fmt.Println("Working hard with PHP")
}

type Java struct {
}

func (j *Java) leanHard() {
	fmt.Println("Working hard with Java")
}

type Developer struct {
	name                string
	programmingLanguage ProgrammingLanguage
}

func (d *Developer) currentProgrammingLeanHard() {
	d.programmingLanguage.leanHard()
}

func makeDeveloper(name string, pl ProgrammingLanguage) *Developer {
	return &Developer{
		name:                name,
		programmingLanguage: pl,
	}
}

func (d *Developer) decideNewProgrammingLanguage(pl ProgrammingLanguage) {
	d.programmingLanguage = pl
}

func main() {
	java := &Java{}
	developer := makeDeveloper("ninehundred", java)
	developer.currentProgrammingLeanHard()

	php := &PHP{}
	developer.decideNewProgrammingLanguage(php)
	developer.currentProgrammingLeanHard()

	golang := &Go{}
	developer.decideNewProgrammingLanguage(golang)
	developer.currentProgrammingLeanHard()
}
