package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
	"strings"
	"sort"
)

// Check if there is an error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	exe2()	
}

func exe2() {
	file, err := os.Open("data.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	total := 0 
	for scanner.Scan() {
		var text = scanner.Text()
		textConverted := convertLettersToNumber(text);
		calibration, err := strconv.Atoi(textConverted);
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		
		total += calibration
		fmt.Println(text, calibration, total)
	}
	check(scanner.Err())

	fmt.Println(total)
}

func exe1() {
	file, err := os.Open("data.txt")
	check(err)

	var onlyNumberRegex = "[0-9]"
	regex := regexp.MustCompile(onlyNumberRegex)
	scanner := bufio.NewScanner(file)
	total := 0 
	for scanner.Scan() {
		var text = scanner.Text()
		res := regex.FindAllString(text, -1);
		calibration, err := strconv.Atoi(res[0] + res[len(res)-1]);
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		total += calibration
		fmt.Println(res, total, calibration)
	}
	check(scanner.Err())

	fmt.Println(total)
}

func convertLettersToNumber(line string) string {
	
	lineNumbers := map[int]string {};
	
	numbers := map[string]string { "one":"1",
		"two":"2",
		"three":"3",
		"four":"4",
		"five":"5",
		"six":"6",
		"seven":"7",
		"eight":"8",
		"nine":"9"}
	
	randomString := "lfksjadflkasdjflkasjdf";
	
	for index, number := range numbers {
		// Test the letter representation
		position := strings.Index(line, index)
		temp := strings.Replace(line,index,randomString[:len(index)],1);
		for position != -1 {
			lineNumbers[position] = number;
			position = strings.Index(temp, index)
			temp = strings.Replace(temp,index,randomString[:len(index)],1);
		}
		
		// Test the numerical representation
		position = strings.Index(line, number)
		temp = strings.Replace(line,number,randomString[:len(number)],1);
		for position != -1 {
			lineNumbers[position] = number
			position = strings.Index(temp, number)
			temp = strings.Replace(temp,number,randomString[:len(number)],1);
		}
		
	}
	
	// Extract keys from map
	keys := make([]int, 0, len(lineNumbers))
	for k := range lineNumbers {
		keys = append(keys, k)
	}

	// Sort keys
	sort.Ints(keys)
	fmt.Println(keys,lineNumbers)
	return lineNumbers[keys[0]] + lineNumbers[keys[len(keys) - 1]]
}