package main

import "fmt"

type animal string

func newAnimal() animal         { return animal("animal") }
func (a animal) String() string { return string(a) }

type mouse struct {
	parent animal
}

func newMouse(animal animal) mouse { return mouse{animal} }
func (m mouse) String() string     { return fmt.Sprintf("<mouse:%s>", m.parent) }

type flyer struct {
	parent animal
}

func newFlyer(animal animal) flyer { return flyer{animal} }
func (f flyer) String() string     { return fmt.Sprintf("<flyer:%s>", f.parent) }

type bat struct {
	mouseParent mouse
	flyerParent flyer
}

func newBat(mouse mouse, flyer flyer) bat { return bat{mouse, flyer} }
func (b bat) String() string              { return fmt.Sprintf("%s %s", b.mouseParent, b.flyerParent) }

func main() {
	fmt.Println(InitializeBat())
}
