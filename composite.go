package main

import "fmt"

type Graphic interface {
	draw()
}

type dot struct {
	x, y int
}

func (d *dot) draw() {
	fmt.Printf("Draw a dot at x: %d - y: %d\n", d.x, d.y)
}

//type circle struct {
//	x,y,r int
//}
//
//func (c *circle) draw() {
//	fmt.Printf("Draw a circle at x: %d - y: %d with radius: %d\n", c.x, c.y, c.r)
//}

type CompoundGraphic struct {
	children []Graphic
}

func (cg *CompoundGraphic) add(child Graphic) {
	cg.children = append(cg.children, child)
	fmt.Printf("Added a new element to drawing machine\n")
}

func (cg *CompoundGraphic) draw() {

	if len(cg.children) == 2 {
		fmt.Println("Drawing a line")
		for _, child := range cg.children {
			child.draw()
		}
		fmt.Println("A line has been drawn")
	}

	if len(cg.children) == 3 {
		fmt.Println("Drawing a triangle")
		for _, child := range cg.children {
			child.draw()
		}
		fmt.Println("A triangle has been drawn")
	}

	if len(cg.children) == 4 {
		fmt.Println("Drawing a rectangle")
		for _, child := range cg.children {
			child.draw()
		}
		fmt.Println("A rectangle has been drawn")
	}
}

func main() {
	figure1 := &dot{2, 3}
	figure2 := &dot{6, 7}
	figure3 := &dot{10, 5}
	figure4 := &dot{0, 1}

	drawing := &CompoundGraphic{}
	drawing.add(figure1)
	drawing.add(figure2)
	drawing.add(figure3)
	drawing.add(figure4)
	drawing.draw()
}
