package main

import (
	"fmt"
)

func main() {
	statesNeeded := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	stations := make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktow"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"az", "ca"}

	fmt.Println(findStationsGreedy(statesNeeded, stations))
}

// TODO: have to deal with making the subset of the stations set
// func findStationsExact(needed []string, stations map[string][]string) []string {
// }

func findStationsGreedy(needed []string, stations map[string][]string) []string {
	// TODO: it doesn't kill if we check if all the needed states are offered by the stations
	finalStations := []string{}
	// go until we find all the needed states
	for len(needed) != 0 {
		// loop through every station to find the one that covers the most states in needed
		var bestStation string
		stateCovered := []string{}
		for station, states := range stations {
			covered, n := slicesIntersection(states, needed)
			if n > len(stateCovered) {
				bestStation = station
				stateCovered = covered
			}
		}

		// remove added states from needed
		for _, i := range stateCovered {
			for index, j := range needed {
				if i == j {
					needed = append(needed[:index], needed[index+1:]...)
				}
			}
		}

		// add the best station that we found in this iteration
		finalStations = append(finalStations, bestStation)
	}
	return finalStations
}

func slicesIntersection(s1 []string, s2 []string) ([]string, int) {
	inboth := []string{}
	count := 0
	for _, i := range s1 {
		for _, j := range s2 {
			if i == j {
				inboth = append(inboth, i)
				count++
			}
		}
	}
	return inboth, count
}
