package main

import (
	"fmt"
)

//=======MONEY=======
type Money int64

func (m Money) String() string {
	return fmt.Sprintf("%+6d.%.2d", m/100, (m % 100).Abs())
}

func (m Money) Abs() Money {
	if m < 0 {
		return m * -1
	}
	return m
}

//======ACCOUNT=====
type Account struct {
	Name      string              `json:"name"`
	Envelopes map[string]Envelope `json:"envelopes"`
}

func NewAccount(name string) Account {
	return Account{Name: name, Envelopes: make(map[string]Envelope)}
}

func (a Account) GetTotal() Money {
	sum := Money(0)
	for _, e := range a.Envelopes {
		sum += e.Total
	}
	return sum
}

func (a Account) Print() {
	for _, e := range a.Envelopes {
		e.Print()
	}
	fmt.Println("----------------------")
	fmt.Printf("|%9s: %-9s|\n", "total", a.GetTotal())
	fmt.Println("----------------------")
}

//======ENVELOPE=====
type Envelope struct {
	Name  string `json:"name"`
	Total Money  `json:"total"`
}

func (e *Envelope) Print() {
	fmt.Printf("|%9s: %-9v|\n", e.Name, e.Total)
}

func (e *Envelope) Add(m Money) {
	e.Total = e.Total + m
}
