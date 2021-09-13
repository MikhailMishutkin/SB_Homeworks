// Подмножества

package main

import "fmt"

func main() {
	langs := []string{"python", "go", "c", "js"}
	compiledLangs := langs[1:3]

	fmt.Println(langs)
	fmt.Println(compiledLangs) // [go c]
}
