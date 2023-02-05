package main

import (
	"myapp/packageone"
)

var myVar = "This is myVar and I am package level in main"

func main() {
	// variables
	// declare a package level variable for the main
	// package named myVar

	// declare a block variable for the main function
	// called blockVar
	var blockVar = "This is blockVar and I am block level in main function"

	// declare a package level vairable in the packageone
	// package name PackageVar

	// create an exported fuction in packageone called PrintMe

	// in the main function, print out the values of myVar,
	// blockVar, and PackageVar on one line using the PrintMe
	// function in packageone
	packageone.PrintMe(myVar)
	packageone.PrintMe(blockVar)
	packageone.PrintMe(packageone.PackageVar)

}
