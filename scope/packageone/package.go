package packageone

import "fmt"

var privateVar = "I am private"
var PublicVar = "I am public or exported."
var PackageVar = "This is PackageVar and I am package level in packageone"

func notExported() {}

func Exported() {}

func PrintMe(input string) {
	fmt.Print(input)
}
