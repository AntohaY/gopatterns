package main

import "fmt"

type game interface {
	insertIntoRoundHole()
}

type roundFigure struct {
}

func (rf *roundFigure) insertIntoRoundHole() {
	fmt.Println("Insert round figure into round hole")
}

type squareFigure struct{}

func (sf *squareFigure) insertInSquareHole() {
	fmt.Println("Insert square figure into square hole")
}

type squareAdapter struct {
	squareM *squareFigure
}

func (sa *squareAdapter) insertIntoRoundHole() {
	sa.squareM.insertInSquareHole()
}

type client struct {
}

func (c *client) insertRoundFigureIntoGameBoard(gm game) {
	gm.insertIntoRoundHole()
}

func main() {
	client := &client{}
	roundFigure := &roundFigure{}
	client.insertRoundFigureIntoGameBoard(roundFigure)
	squareM := &squareFigure{}
	squareAdapterMachine := &squareAdapter{
		squareM: squareM,
	}
	client.insertRoundFigureIntoGameBoard(squareAdapterMachine)
}
