package puppy

import "fmt"

const sound string = "Woof!"

func Bark() {
	question := "how many times does the dog bark"
	numTimes := "twice"

	fmt.Printf("%v? %v. %v", question, numTimes, sound)
}
