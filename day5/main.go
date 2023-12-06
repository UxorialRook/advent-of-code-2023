package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
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
	
	// Get the seeds
	split := strings.Split(string(text), "\n\n")
	strings.Split(split[0], "")
	seeds := strings.Fields(strings.Split(split[0], ":")[1])
	
	// All the other lines crrespond to the map
	mapPipe := parseMaps(split[1:])
	
	minLoc := 0
	for _, seed := range seeds {
		location, _ := strconv.Atoi(seed)

		newLoc := scanMaps(location, mapPipe)
		if minLoc == 0 || newLoc < minLoc {
			minLoc = newLoc
		}
	}

	fmt.Println(minLoc)
}

// Best time: 1m3s to execute on my laptop
func exe2() {
	var text, err = os.ReadFile("data.txt")
	check(err)
	
	split := strings.Split(string(text), "\n\n")
	
	strings.Split(split[0], "")
	seedsRanges := strings.Fields(strings.Split(split[0], ":")[1])
	mapPipe := parseMaps(split[1:])

	var waitGroup sync.WaitGroup
	var mutex sync.Mutex
	var minLocation int

	start := time.Now()
	
	for i := 0; i < len(seedsRanges); i += 2 {
		waitGroup.Add(1)
		startRange, _ := strconv.Atoi(seedsRanges[i])
		lengthRange, _ := strconv.Atoi(seedsRanges[i+1])

		go func() {
			defer waitGroup.Done()
			var rangeMinLocation int

			for i := 0; i < lengthRange; i++ {
				seed := startRange + i
				location := seed

				newLoc := scanMaps(location, mapPipe)

				if rangeMinLocation == 0 || newLoc < rangeMinLocation {
					rangeMinLocation = newLoc
				}
			}

			mutex.Lock()
			defer mutex.Unlock()

			if minLocation == 0 || rangeMinLocation < minLocation {
				minLocation = rangeMinLocation
			}
		}()
	}

	waitGroup.Wait()
	elapsed := time.Since(start)
	fmt.Println("Exe 2 took %s", elapsed)
	fmt.Println(minLocation)
}

// Scan the maps to return the location
func scanMaps(location int, mapPipe map[int][]SeedMap) int {
	for i := 0; i <= len(mapPipe); i++ {
		for _, mapVal := range mapPipe[i] {
			if location >= mapVal.sourceRangeStart && location < mapVal.sourceRangeStart+mapVal.rangeLenght {
				diff := location - mapVal.sourceRangeStart
				location = mapVal.destRangeStart + diff
				break
			}
		}
	}

	return location
}

// Parse the map to create a matrice of SeedMap
func parseMaps(maps []string) map[int][]SeedMap {
	var mapPipe = map[int][]SeedMap{}

	for i, mapValues := range maps {
		values := strings.Split(mapValues, "\n")[1:]
		for _, mapVal := range values {
			if len(mapVal) == 0 {
				continue
			}
			input := strings.Fields(mapVal)
			destRangeStart, _ := strconv.Atoi(input[0])
			sourceRangeStart, _ := strconv.Atoi(input[1])
			rangeLength, _ := strconv.Atoi(input[2])
			mapPipe[i] = append(mapPipe[i], SeedMap{destRangeStart, sourceRangeStart, rangeLength})
		}
	}

	return mapPipe
}

type SeedMap struct {
	destRangeStart   int
	sourceRangeStart int
	rangeLenght      int
}