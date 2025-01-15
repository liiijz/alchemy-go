package main

import (
	"alchemy"
	"fmt"
)

func main() {
	client := alchemy.NewClient("121212")
	client.GetNews("bitcoin")

	fmt.Println("Hello World")
}
