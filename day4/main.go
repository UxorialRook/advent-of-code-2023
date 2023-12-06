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
	exe2()
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

func exe2() {
	res := map[int]game {}
	file, err := os.Open("data.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	numGame := 0
	for scanner.Scan() {
		numGame += 1
		var text = scanner.Text()
		text = strings.Split(text, ":")[1]
		subset := strings.Split(text, "|");
		var onlyNumberRegex = "[0-9]+"
		regex := regexp.MustCompile(onlyNumberRegex)

		winNumbers := regex.FindAllString(subset[0],-1)
		fmt.Println(winNumbers)

		allNumbers  := regex.FindAllString(subset[1],-1)
		fmt.Println(allNumbers)

		nbCopies := 0
		for _, number := range allNumbers {
			if contains(winNumbers, number) {
				nbCopies += 1
			}
		}

		res[numGame] = game{numGame, nbCopies, 1}

	}

	total := 0
	// For each game, we need to add nbTimesToExecute for each nbCopies found
	// Ex: if we found 4 copies, we add nbTimesToExecute to the next 4 games.
	// nbTimesToExecute being the number of time we are going to execute the current game
	for i := 1; i <= len(res); i++ {
		game := res[i]
		for game.nbCopies > 0 {
			otherGame := res[game.numGame + game.nbCopies]
			otherGame.nbTimesToExecute += game.nbTimesToExecute
			res[game.numGame + game.nbCopies] = otherGame
			game.nbCopies -= 1
		}
		total += game.nbTimesToExecute
	}

	fmt.Println(total)
}

// Basic contain function
func contains(haystack []string, needel string) bool {
	for _, element := range haystack {
		if element == needel {
			return true
		}
	}
	return false
}

type game struct {
	numGame int // ID of the game (counter)
	nbCopies int // Number of copies (score of the current game)
	nbTimesToExecute int // Number of times we need to execute the current game
}