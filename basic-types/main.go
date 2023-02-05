package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}

var keyPressChan chan rune

func main() {

	/*
		x := 10

		myFirstPointer := &x

		fmt.Println("x is", x)
		fmt.Println("myFirstPointer is", myFirstPointer)
		*myFirstPointer = 15
		fmt.Println("x is now", x)
		changeValue(&x)
		fmt.Println("Value of x after changeValue.", x)

		fmt.Println("Slices code:")
		slices()
		fmt.Println("\nMap code:")
		usingMaps()

		fmt.Printf("\ntype func:\n")
		var dog Animal
		dog.Name = "dog"
		dog.Sound = "woof"
		dog.Says()

		cat := Animal{
			Name:         "cat",
			Sound:        "meow",
			NumberOfLegs: 4,
		}
		cat.Says()
	*/

	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}

}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}

}

func changeValue(num *int) {
	fmt.Println("What is the value of num?", *num)
	*num = 20
}

func slices() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "fish")
	animals = append(animals, "horse")

	fmt.Println(animals[0:1])

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Delete fish which is index 2")
	animals = deleteFromSlice(animals, 2)
	fmt.Println("animals:", animals)
}

func deleteFromSlice(a []string, index int) []string {
	a[index] = a[len(a)-1]
	//a[len(a)-1] = "" // Why did we have to do this?
	a = a[:len(a)-1]
	return a
}

func usingMaps() {
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	delete(intMap, "four")
	fmt.Println("After delete")

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	fmt.Println("\nFind in map:")
	el, ok := intMap["two"]
	if ok {
		fmt.Println("Element is in map", el, ok)
	} else {
		fmt.Println("Element is not in the map", el, ok)
	}
}
