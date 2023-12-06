package main

import (
	
	"fmt"
	"os"
	"strings"
	"regexp"
	"strconv"
	"math"
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

func exe1() {
	var text, err = os.ReadFile("data.txt")
	check(err)
	
	split := strings.Split(string(text), "\n")
	
	var onlyNumberRegex = "[0-9]+"
	regexNumber := regexp.MustCompile(onlyNumberRegex)
	
	games := []game{}
	
	times := regexNumber.FindAllString(split[0],-1)
	distances := regexNumber.FindAllString(split[1],-1)
	for i := 0; i < len(times); i++ {
		time,_ := strconv.Atoi(times[i])
		distance,_ := strconv.Atoi(distances[i])
		games = append(games, game{time,distance})
	}
	
	total := 1
	for _, game := range games {
		combinations := calcCombinations(game)
		fmt.Println(total, " * ", combinations, total * combinations)
		total *= combinations
		fmt.Println(total)
	}
	
	fmt.Println(total)
}

func exe2() {
	var text, err = os.ReadFile("data.txt")
	check(err)

	split := strings.Split(string(text), "\n")

	var onlyNumberRegex = "[0-9]+"
	regexNumber := regexp.MustCompile(onlyNumberRegex)

	games := []game{}

	times := regexNumber.FindAllString(split[0],-1)
	distances := regexNumber.FindAllString(split[1],-1)
	var totalTime string
	var totalDistance string
	for i := 0; i < len(times); i++ {
		totalTime += times[i]
		totalDistance += distances[i]
	}
	
	distance,_ := strconv.Atoi(totalDistance)
	time,_ := strconv.Atoi(totalTime)
	fmt.Println(totalTime, totalDistance)
	games = append(games, game{time,distance})

	total := 1
	for _, game := range games {
		combinations := calcCombinations(game)
		total *= combinations
	}
	
	fmt.Println(total)
}

// Optimized version of calcCombinations using a quadratic function
// If we take the problem as a mathematical equation we find the following equality:
// d = (acceleration * accelerationTime) * movingTime
//
// But we know that movingTime = totalTime - accelerationTime:
//
// d = (acceleration * accelerationTime) * (totalTime - accelerationTime)
// d = acceleration * accelerationTime * totalTime - acceleration * accelerationTime^2
// d = acceleration * totalTime * accelerationTime - acceleration * accelerationTime^2
//
// We found the quadratic function:
//
// acceleration * totalTime * accelerationTime - acceleration * accelerationTime^2 + d = 0
//
func calcCombinations(game game) int {
	b := float64(game.time)
	sqrt := math.Sqrt(b*b - 4 * float64(game.distance))
	
	res1 := (-b - sqrt) / 2
	res2 := (-b + sqrt) / 2
	
	return int (res2 - res1)
}

type game struct {
	time int
	distance int
}