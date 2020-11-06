package main

import (
	"fmt"
	"regexp"
)

func main() {

	websites := [...]string{"webcode.me", "zetcode.com", "freebsd.org", "netbsd.org"}

	re := regexp.MustCompile("(\\w+)\\.(\\w+)")

	for _, website := range websites {

		parts := re.FindStringSubmatch(website)

		for i, _ := range parts {
			fmt.Println(parts[i])
		}

		fmt.Println("---------------------")
	}
}
