package main

import "fmt"

type Champion struct {
	Name            string
	Attack, Defense int
}

func NewChampion(name string, attack, defense int) *Champion {
	return &Champion{Name: name, Attack: attack, Defense: defense}
}

func (c *Champion) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type ChampionModifier struct {
	champion *Champion
	next     Modifier
}

func (c *ChampionModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *ChampionModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

type IncreaseDefenseModifier struct {
	ChampionModifier
	Percentage float32
}

func NewIncreaseDefenseModifier(c *Champion, percentage float32) *IncreaseDefenseModifier {
	return &IncreaseDefenseModifier{ChampionModifier{
		champion: c,
	}, percentage}
}

func (i *IncreaseDefenseModifier) Handle() {
	fmt.Println("Increasing", i.champion.Name, "defense")
	i.champion.Defense *= int(i.Percentage)
	i.ChampionModifier.Handle()
}

func NewChampionModifier(c *Champion) *ChampionModifier {
	return &ChampionModifier{champion: c}
}

func main() {
	t := NewChampion("Teemo", 10, 3)
	fmt.Println(t.String())

	root := NewChampionModifier(t)
	root.Add(NewIncreaseDefenseModifier(t, 2.25))

	root.Handle()
	fmt.Println(t.String())
}
