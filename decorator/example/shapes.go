package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

type Square struct {
	Side float32
}

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *Circle) Render() string {
	return fmt.Sprintf("A circle of radius %f", c.Radius)
}

func (s *Square) Render() string {
	return fmt.Sprintf("A square with side %f", s.Side)
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

func Resize(c *Circle, factor float32) {
	c.Radius *= factor
}

func main() {
	circle := Circle{2}
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "red"}
	fmt.Println(redCircle.Render())
}
