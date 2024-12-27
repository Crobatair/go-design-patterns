package main

import "fmt"

type AgedAnimal interface {
	Age() int
	SetAge(l int)
}

type Bird struct {
	age int
}

type Lizard struct {
	age int
}

func (b *Bird) Age() int {
	return b.age
}

func (b *Bird) SetAge(age int) {
	b.age = age
}

func (b *Bird) Fly() {
	if b.age >= 3 {
		fmt.Println("Flying!")
	}
}

func (l *Lizard) Crawl() {
	switch {
	case l.Age() < 10:
		fmt.Println("Mini Crawl!")
	case l.Age() >= 10:
		fmt.Println("Crawling!")
	default:
		fmt.Println("Dead!")
	}
}

func (l *Lizard) Age() int {
	return l.age
}

func (l *Lizard) SetAge(age int) {
	l.age = age
}

type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (d *Dragon) Age() int {
	return d.bird.Age()
}

func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

func NewDragon(age int) *Dragon {
	return &Dragon{Bird{age}, Lizard{age}}
}

func main() {
	dragon := NewDragon(9)
	dragon.Fly()
	dragon.Crawl()
}
