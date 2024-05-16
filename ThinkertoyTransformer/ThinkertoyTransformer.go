package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cont, err := os.ReadFile("thinkertoy.txt")
	str := string(cont)
	if err != nil {
		fmt.Println("LOL")
	}
	vale := strings.NewReplacer(
		"o", "a",
		"0", "c",
		"O", "A",
		"|", "I",
		"\\", "+",
	)
	os.WriteFile("result.txt", []byte((vale.Replace(str))), 0o644)
}
