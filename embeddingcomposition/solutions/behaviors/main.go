package main

import "fmt"

// Storer is an interface for things that can store
// cargo
type Storer interface {
	Stow(string)
	Inventory() []string
}
type Bicycle struct {
	Basket []string
}

func (b *Bicycle) Stow(item string) {
	b.Basket = append(b.Basket, item)
}

func (b *Bicycle) Inventory() []string {
	return b.Basket
}

type Subaru struct {
	Trunk []string
}

func (s *Subaru) Stow(item string) {
	s.Trunk = append(s.Trunk, item)
}

func (s *Subaru) Inventory() []string {
	return s.Trunk
}

type Boeing struct {
	CargoHold []string
}

func (b *Boeing) Stow(item string) {
	b.CargoHold = append(b.CargoHold, item)
}

func (b *Boeing) Inventory() []string {
	return b.CargoHold
}

func main() {
	s := &Subaru{}
	b := &Boeing{}
	c := &Bicycle{}

	PutCargo(s, "spare tire")
	PutCargo(b, "suitcase")
	PutCargo(c, "bread")

	GetInventory(s)
	GetInventory(b)
	GetInventory(c)

}

func PutCargo(s Storer, item string) {
	s.Stow(item)
}

func GetInventory(s Storer) {
	fmt.Println(s.Inventory())
}
