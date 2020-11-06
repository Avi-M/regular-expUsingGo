package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {

	words := [...]string{"Seven", "even", "Maven", "Amen", "eleven"}

	re, err := regexp.Compile(".even")

	if err != nil {
		log.Fatal(err)
	}

	for _, word := range words {

		found := re.MatchString(word)

		if found {

			fmt.Printf("%s matches\n", word)
		} else {

			fmt.Printf("%s does not match\n", word)
		}
	}
}
