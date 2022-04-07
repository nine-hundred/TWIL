package main

import "fmt"

// 고객이 웨이터에게 음식을 주문하는 command pattern

type FoodCommand interface {
	order()
}

type RisottoCommand struct {
	employee Employee
}

func (risotto *RisottoCommand) order() {
	risotto.employee.
		listen("risotto")
	fmt.Println("the risotto is ordered")
}

type PastaCommand struct {
	employee Employee
}

func (pasta *PastaCommand) order() {
	pasta.employee.
		listen("pasta")
	fmt.Println("the pasta is ordered")
}

type Customer struct {
	command []FoodCommand
	bills   int
}

func (c *Customer) call() {
	for _, command := range c.command {
		command.order()
	}
}

type Employee interface {
	listen(food string)
	stick()
}

type Waiter struct {
	note []string
}

func (w *Waiter) listen(food string) {
	w.note = append(w.note, food)
	fmt.Println("listen and write note :", food)
}

func (w *Waiter) stick() {
	fmt.Println("pass to chef :", w.note)
}

func main() {
	waiter := &Waiter{}

	risotto := &RisottoCommand{
		employee: waiter,
	}

	pasta := &PastaCommand{
		employee: waiter,
	}

	foods := []FoodCommand{risotto, pasta}

	customer := &Customer{
		command: foods,
	}

	customer.call()
	waiter.stick()
}
