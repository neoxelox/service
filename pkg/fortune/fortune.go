package fortune

import "fmt"

func init() {
	fmt.Println("Fortune working!")
}

// ExportedFunc is exported!
func ExportedFunc() {
	fmt.Println("Exporetd func!")
}
