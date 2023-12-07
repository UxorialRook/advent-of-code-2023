package main

import (
	
	"fmt"
	"bufio"
	"os"
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
	exe1()
}

func exe1() {
	var file, err = os.Open("data.txt")
	check(err)
	
	scanner := bufio.NewScanner(file)
	games := []camelGame{}
	
	for scanner.Scan() {
		var text = scanner.Text()
		elements := strings.Split(string(text), " ")
		hand := elements[0]
		bid, _ := strconv.Atoi(elements[1])
		rate := rateType(hand)
		games = append(games, camelGame{hand, bid, rate})
	}
	
	sort.Slice(games, func(i, j int) bool {
		if games[i].rate == games[j].rate {
			for k := 0; k < len(games[i].hand); k++ {
				if charToNumber(games[i].hand[k]) < charToNumber(games[j].hand[k]) {
					return true
				} else if charToNumber(games[i].hand[k]) > charToNumber(games[j].hand[k]) {
					return false
				}
			} 
		}
		return games[i].rate < games[j].rate
	})
	
	total := 0
	for i := 0; i < len(games); i++ {
		total += (i+1) * games[i].bid
	}
	
	fmt.Println(games, total)
	
}

func charToNumber(character uint8) int {
	c := fmt.Sprintf("%c",character)
	res := map[string]int{"A":13,
		"K":12,
		"Q":11,
		"J":10,
		"T":9,
		"9":8,
		"8":7,
		"7":6,
		"6":5,
		"5":4,
		"4":3,
		"3":2,
		"2":1,
	}
	
	return res[c]
}
func rateType(hand string) int {
	letters := map[string]int{}
	runes := []rune(hand)
	for _, character := range runes {
		c := fmt.Sprintf("%c",character)
		_, ok := letters[c]
		if !ok {
			letters[c] = strings.Count(hand, c)
		}
	}
	
	return calcPoints(letters)
}

func calcPoints(letters map[string]int) int {
	keys := sortByValues(letters)
	if len(letters) == 1 {
		return 6
	} else if len(letters) == 2 {
		if letters[keys[0]] == 4 {
			return 5
		} else {
			return 4		
		}
	} else if len(letters) == 3 {
		fmt.Println(keys)
		if letters[keys[0]] == 3 {
			return 3
		} else {
			return 2
		}
	} else if len(letters) == 4 {
		return 1
	} else {
		return 0	
	}	
}

func sortByValues(letters map[string]int) []string {
	keys := make([]string, 0, len(letters))

	for key := range letters {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool{
		return letters[keys[i]] > letters[keys[j]]
	})
	
	return keys
}

func exe2() {
}

type camelGame struct {
	hand string
	bid int
	rate int
}