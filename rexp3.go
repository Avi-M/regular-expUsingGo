package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	var content = `Foxes are omnivorous mammals belonging to several genera 
    of the family Canidae. Foxes have a flattened skull, upright triangular ears, 
    a pointed, slightly upturned snout, and a long bushy tail. Foxes live on every 
    continent except Antarctica. By far the most common and widespread species of 
    fox is the red fox.`

	re := regexp.MustCompile("(?i)fox(es)?")

	found := re.FindAllString(content, -1)

	fmt.Printf("%q\n", found)

	if found == nil {
		fmt.Printf("no match found\n")
		os.Exit(1)
	}

	for _, word := range found {
		fmt.Printf("%s\n", word)
	}

}
