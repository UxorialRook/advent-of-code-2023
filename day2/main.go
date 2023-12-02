package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
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
	inputRed := 12
	inputGreen := 13
	inputBlue := 14
	
	file, err := os.Open("data.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	gameNumber := 0
	total := 0
	
	for scanner.Scan() {
		gameNumber += 1
		var text = scanner.Text()
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		maxByColor := map[string]int {};
		data := strings.Split(text, ":");
		subsets := strings.Split(data[1], ";")
		
		// Split all the subsets in each game
		for _, subset := range subsets {
			cubesBySubset := strings.Split(subset, ",")
			// Split all the cubes in the subset
			for _, cube := range cubesBySubset {
				numberCube := strings.Split(cube, " ")
				currentColor := numberCube[2]
				currentNumber, _ := strconv.Atoi(numberCube[1])
				
				_, ok := maxByColor[currentColor]
				if !ok || currentNumber > maxByColor[currentColor] {
					maxByColor[currentColor] = currentNumber
				}
			}
			
		}
		isOk := true
		for color, value := range maxByColor {
			if color == "red" && value > inputRed  {
				isOk = false;
				break
			}
			if color == "blue" && value > inputBlue {
				isOk = false
				break
			}
			if color == "green" && value > inputGreen {
				isOk = false
				break
			}
		}

		if isOk {
			total += gameNumber
		}
		
	}
	fmt.Println(total)
	check(scanner.Err())
}