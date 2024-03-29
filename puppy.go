package puppy

import (
	"fmt"

	"github.com/adarsh-shahi/doggo"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return doggo.WhenGrownUp(Bark())
}

func BigBarks() string {
	return doggo.WhenGrownUp(Barks())
}

func From1() {
	fmt.Println("version 1")
}
func From2() {
	fmt.Println("version 2")
}
