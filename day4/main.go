package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"regexp"
)

// Check if there is an error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	exe1()
}

func exe1() {
	file, err := os.Open("data.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	bigTotal := 0
	for scanner.Scan() {
		var text = scanner.Text()
		text = strings.Split(text, ":")[1]
		subset := strings.Split(text, "|");
		var onlyNumberRegex = "[0-9]+"
		regex := regexp.MustCompile(onlyNumberRegex)
		
			winNumbers := regex.FindAllString(subset[0],-1)
			fmt.Println(winNumbers)
			
			allNumbers  := regex.FindAllString(subset[1],-1)
			fmt.Println(allNumbers)
			
			total := 0
			for _, number := range allNumbers {
				if contains(winNumbers, number) {
					if total == 0 {
						total = 1
					} else {
						total *= 2
					}
				}
			}
			
			bigTotal += total
	}
	fmt.Println(bigTotal)
}


func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}