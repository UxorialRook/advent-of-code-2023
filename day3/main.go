package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"regexp"
	"unicode"
	"math/rand"
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
	totalSum := 0
	elements := []element{}
	file, err := os.Open("data.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	gearRatios := map[int]gear {};

	nLine := 0
	for scanner.Scan() {
		var text = scanner.Text()
		nLine += 1
		subset := strings.Split(text, ".");
		repStr := ".............................."
		var onlyNumberRegex = "[0-9]+"
		regexNumber := regexp.MustCompile(onlyNumberRegex)
		for _, value := range subset {
			startPos := strings.Index(text, value)
			endPos := startPos + len(value) - 1
			// If the value is different than empty, it means there is an object (could be a number or a sign) that we want to parse
			if value != "" {
				if isNumeric(value) {
					elements = append(elements, element{rand.Int(),value, nLine, startPos, endPos, false, false})
				} else {
					// If the len is bigger than 1, it means the number is followed or precessed by a sign.
					if len(value) > 1 {
						for _, character := range value {
							if !unicode.IsDigit(character) {
								startPosCharacter := strings.Index(text, fmt.Sprintf("%c",character))
								endPosCharacter := startPosCharacter
								elements = append(elements, element{rand.Int(),fmt.Sprintf("%c",character), nLine, startPosCharacter,
									endPosCharacter, true, false})
								text = strings.Replace(text, fmt.Sprintf("%c",character), ".",1)
							}
						}
						allNumbers := regexNumber.FindAllString(value,-1)

						for _, number := range allNumbers {
							startPosNumber := strings.Index(text, number)
							endPosNumber := startPosNumber + len(number) - 1
							elements = append(elements, element{rand.Int(),number, nLine, startPosNumber, endPosNumber, false, false})
							text = strings.Replace(text, number, repStr[:len(number)],1)
						}
					} else {
						elements = append(elements, element{rand.Int(),value, nLine, startPos, endPos, true, false})
					}
				}
				text = strings.Replace(text, value, repStr[:len(value)],1)
			}
		}
	}
	for key, element := range elements {
		if !element.isOk {
			for _, otherElement := range elements {
				// If we are in the current line, it must be a sign that is right before or right after the current element
				if otherElement.isSign && otherElement.line == element.line && (otherElement.startPos == element.endPos + 1 || otherElement.endPos == element.startPos - 1) {
					elements[key].isOk = true
					if otherElement.value == "*" {
						currVal, ok := gearRatios[otherElement.id]
						valToInt, _ := strconv.Atoi(element.value)
						if !ok {
							currVal = gear{1,1}
						}
						gearRatios[otherElement.id] = shouldMultiply(otherElement.value,currVal,valToInt)
					}
				}

				// If we are in the line before, it must be a sign that is in the range [start -1 ; end + 1] with start and end being the startPos and endPod of the current element
				if otherElement.isSign && otherElement.line == element.line - 1 && otherElement.startPos >= element.startPos - 1 && otherElement.endPos <= element.endPos + 1 {
					elements[key].isOk = true
					if otherElement.value == "*" {
						currVal, ok := gearRatios[otherElement.id]
						valToInt, _ := strconv.Atoi(element.value)
						if !ok {
							currVal = gear{1,1}
						}
						gearRatios[otherElement.id] = shouldMultiply(otherElement.value,currVal,valToInt)
					}
				}

				// If we are in the line after, it must be a sign that is in the range [start -1 ; end + 1] with start and end being the startPos and endPod of the current element
				if otherElement.isSign && otherElement.line == element.line + 1 && otherElement.startPos >= element.startPos - 1 && otherElement.endPos <= element.endPos + 1 {
					elements[key].isOk = true
					if otherElement.value == "*" {
						currVal, ok := gearRatios[otherElement.id]
						valToInt, _ := strconv.Atoi(element.value)
						if !ok {
							currVal = gear{1,1}
						}
						gearRatios[otherElement.id] = shouldMultiply(otherElement.value,currVal,valToInt)
					}
				}
			}
		}
	}

	for _, gearRatio :=range gearRatios {
		fmt.Println(gearRatio.value)
		if gearRatio.number > 2 {
			totalSum += gearRatio.value
		}
	}

	fmt.Println(totalSum)
}

func exe1() {
	totalSum := 0
	elements := []element{}
	file, err := os.Open("data.txt")
	check(err)
	scanner := bufio.NewScanner(file)

	nLine := 0
	for scanner.Scan() {
		var text = scanner.Text()
		nLine += 1
		subset := strings.Split(text, ".");
		repStr := ".............................."
		var onlyNumberRegex = "[0-9]+"
		regexNumber := regexp.MustCompile(onlyNumberRegex)
		for _, value := range subset {
			startPos := strings.Index(text, value)
			endPos := startPos + len(value) - 1
			// If the value is different than empty, it means there is an object (could be a number or a sign) that we want to parse
			if value != "" {
				if isNumeric(value) {
					elements = append(elements, element{rand.Int(), value, nLine, startPos, endPos, false, false})
				} else {
					// If the len is bigger than 1, it means the number is followed or precessed by a sign.
					if len(value) > 1 {
						for _, character := range value {
							if !unicode.IsDigit(character) {
								startPosCharacter := strings.Index(text, fmt.Sprintf("%c",character))
								endPosCharacter := startPosCharacter
								elements = append(elements, element{rand.Int(), fmt.Sprintf("%c",character), nLine, startPosCharacter,
									endPosCharacter, true, false})
								text = strings.Replace(text, fmt.Sprintf("%c",character), ".",1)
							}
						}
						allNumbers := regexNumber.FindAllString(value,-1)

						for _, number := range allNumbers {
							startPosNumber := strings.Index(text, number)
							endPosNumber := startPosNumber + len(number) - 1
							elements = append(elements, element{rand.Int(), number, nLine, startPosNumber, endPosNumber, false, false})
							text = strings.Replace(text, number, repStr[:len(number)],1)
						}
					} else {
						elements = append(elements, element{rand.Int(),value, nLine, startPos, endPos, true, false})
					}
				}
				text = strings.Replace(text, value, repStr[:len(value)],1)
			}
		}
	}
	for key, element := range elements {
		if !element.isOk {
				for _, otherElement := range elements {
				// If we are in the current line, it must be a sign that is right before or right after the current element
				if otherElement.isSign && otherElement.line == element.line && (otherElement.startPos == element.endPos + 1 || otherElement.endPos == element.startPos - 1) {
					elements[key].isOk = true
				}

				// If we are in the line before, it must be a sign that is in the range [start -1 ; end + 1] with start and end being the startPos and endPod of the current element
				if otherElement.isSign && otherElement.line == element.line - 1 && otherElement.startPos >= element.startPos - 1 && otherElement.endPos <= element.endPos + 1 {
					elements[key].isOk = true
				}

				// If we are in the line after, it must be a sign that is in the range [start -1 ; end + 1] with start and end being the startPos and endPod of the current element
				if otherElement.isSign && otherElement.line == element.line + 1 && otherElement.startPos >= element.startPos - 1 && otherElement.endPos <= element.endPos + 1 {
					elements[key].isOk = true
				}
			}
		}
	}

	for _, element :=range elements {
		if element.isOk {
			value, _ := strconv.Atoi(element.value)
			fmt.Println(element.value)

			totalSum += value
		}
	}

	fmt.Println(totalSum)
}

func isNumeric(s string) bool {
	number, err := strconv.Atoi(s)
	return err == nil && number > 0 && !strings.Contains(s,"+")
}

// Test if we should multiply and return the result of the multiplication
func shouldMultiply(symbol string, currGear gear, newValue int) gear {
	if symbol == "*" && currGear.number < 3 {
		return gear{currGear.value * newValue, currGear.number + 1}
	}
	return gear{currGear.value, currGear.number}
}

type element struct {
	id int
	value string
	line int
	startPos int
	endPos int
	isSign bool
	isOk bool
}

type gear struct {
	value int
	number int
}