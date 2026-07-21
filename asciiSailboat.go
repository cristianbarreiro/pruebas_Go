// ASCII Sailboat Drawing in Go
// This program prints a recognizable sailboat with a mast, a sail, and a hull
// using standard keyboard characters and fmt.Println statements.

package main

import "fmt"

func main() {
	// Print the sail and mast
	fmt.Println("          |")
	fmt.Println("         /|")
	fmt.Println("        / |")
	fmt.Println("       /  |")
	fmt.Println("      /   |")
	fmt.Println("     /    |")
	fmt.Println("    /_____|")
	fmt.Println("          |")
	
	// Print the hull (Note: the double \\ is required to print a single backslash)
	fmt.Println("\\---------|---------/")
	fmt.Println(" \\                 /")
	fmt.Println("  \\_______________/")
}