package main

import "fmt"

func main() {
	greetEmpty := hello("")
	fmt.Println(greetEmpty)

	greetMsg := hello("John")
	fmt.Println(greetMsg)

}
