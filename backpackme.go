package main

import (
	"fmt"
	"github.com/ivanBereznev/backpackme/generator"
)

func main() {
	fmt.Println(generator.GenerateMenu(generator.InputParams{Days: 3, People: 4}))
}
