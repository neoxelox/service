// Package fortune is working!
// testing
// 	- 1
// 	- 2
// 	- 3
// 	- 4
package fortune

import "fmt"

func init() {
	fmt.Println("Fortune working!")
}

// ExportedFunc is exported!
func ExportedFunc() {
	fmt.Println("Exporetd func!")
}
