package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Fetch the filename and command-line arguments
	filename := os.Args[0]
	args := os.Args[1:]
	fmt.Println("Filename:", filename)
	fmt.Println("Arguments:", args)

	// Define a flag for the name of the project
	nameOfProject := flag.String("project", "firstapplication", "Name of the project")

	// Parse the command-line flags
	flag.Parse()

	// Access the values of the parsed flags
	projectName := *nameOfProject

	// Now you can use the values in your program
	fmt.Println("Project Name:", projectName)

	// Rest of your program logic goes here...
	// For example, you can check the value of projectName and perform actions accordingly.

	// Example:
	if projectName == "firstapplication" {
		fmt.Println("Default project selected.")
	} else {
		fmt.Println("Custom project selected:", projectName)
	}

	// Uncomment the next line if you want to use the useapi flag
	// fmt.Println("Use API:", *useapi)
}
