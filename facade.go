package main

import (
	"fmt"
	"math/rand"
)

type navigationFacade struct {
	startAddress string
	destination  string
}

func newNavigationFacade(start string, dest string) *navigationFacade {
	fmt.Println("GPS IS STARTING..... BEEP BOOP")
	navigationFacade := &navigationFacade{
		startAddress: start,
		destination:  dest,
	}
	fmt.Println("GPS HAS BEEN TURNED ON")
	return navigationFacade
}

func (n *navigationFacade) loadData(start string, dest string) string {
	fmt.Println()
	fmt.Printf("Loading data for %s and %s\n", start, dest)
	res := "Navigation data for " + start + " and " + dest
	return res
}

func (n *navigationFacade) calculateRoute(start string, dest string) {
	navigationData := n.loadData(start, dest)
	fmt.Println()
	fmt.Printf("Calculating route for %s and %s using %s\n", start, dest, navigationData)
	fmt.Println()
	fmt.Print(locateTraffic())
	fmt.Println()
	fmt.Printf("The route between %s and %s is following: -____----___----___--\n", start, dest)
	fmt.Println("GPS HAS BEEN TURNED OFF, BYE BYE!")
}

func locateTraffic() string {
	res := rand.Intn(2)
	if res == 2 {
		return "Traffic is present :("
	} else {
		return "No traffic :)"
	}
}

func main() {
	fmt.Println()
	navFacade := newNavigationFacade("Kherson", "Kiev")
	fmt.Println()
	navFacade.calculateRoute("Kherson", "Kiev")
}
