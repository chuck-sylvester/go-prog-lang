// echo2.go prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Println("\nHere's another way to do it...")
	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println("\nAnd a quick and dirty method...")
	fmt.Println(os.Args[1:])

	fmt.Println("\nExercise 1.1 - include os.Args[0]")
	fmt.Println(os.Args[0:])

	fmt.Println("\nExercise 1.2 - print the index and value of each argument, one per line")
	for i := 0; i < len(os.Args); i++ {
		fmt.Print(i)
		fmt.Println(sep + os.Args[i])
	}

}
