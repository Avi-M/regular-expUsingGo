package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {

	var data = `22, 1, 3, 4, 5, 17, 4, 3, 21, 4, 5, 1, 48, 9, 42`

	sum := 0

	re := regexp.MustCompile(",\\s*")

	vals := re.Split(data, -1)

	for _, val := range vals {

		n, err := strconv.Atoi(val)

		sum += n

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(sum)
}
