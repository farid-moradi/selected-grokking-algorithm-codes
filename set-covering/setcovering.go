package main

import "fmt"

const UNITED_STATE_STATES = 50

func main() {
	statesNeeded := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	stations := make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"az", "ca"}

	fmt.Println(findStationsGreedy(statesNeeded, stations))
	fmt.Println(findStationsExact(statesNeeded, stations))
}

func findStationsExact(needed []string, stations map[string][]string) []string {
	stationNames := []string{}
	for sname := range stations {
		stationNames = append(stationNames, sname)
	}
	subsets := subsetList(stationNames)

	bestSubset := []string{}
	bestSubsetCount := UNITED_STATE_STATES
	for _, subset := range subsets {
		subsetStations := []string{}
		for _, station := range subset {
			subsetStations = append(subsetStations, stations[station]...)
		}
		if doesItCovers(subsetStations, &needed) {
			if len(subset) < bestSubsetCount {
				bestSubsetCount = len(subset)
				bestSubset = subset
			}
		}
	}
	return bestSubset
}

func doesItCovers(s []string, needed *[]string) bool {
	length := 0
	for _, i := range *needed {
		for _, j := range s {
			if i == j {
				length++
				break
			}
		}
	}
	return length >= len(*needed)
}

func subsetList(s []string) [][]string {
	subsets := [][]string{}
	for i := range s {
		subsets = append(subsets, []string{s[i]})
	}
	subsetListRecursive(s, &subsets)
	return subsets
}

func subsetListRecursive(s []string, subsets *[][]string) {
	if len(s) == 1 {
		return
	}

	fullElements := []string{}
	// for _, element := range s {
	// 	fullElements = append(fullElements, element)
	// }
	fullElements = append(fullElements, s...)
	// fmt.Println(stringExist(subsets, fullElements))
	if !stringExist(subsets, fullElements) {
		*subsets = append(*subsets, fullElements)
	}

	tempslice := make([]string, len(s))
	for e := range s {
		copy(tempslice[:e], s[:e])
		t := append(tempslice[:e], s[e+1:]...)
		subsetListRecursive(t, subsets)
	}
}

func stringExist(s *[][]string, e []string) bool {

	flag := true
	for _, slice := range *s {
		if len(slice) == len(e) {
			flag = true
			for i := range e {
				if slice[i] != e[i] {
					flag = false
					break
				}
			}
			if flag {
				return true
			}
		}
	}
	return false
}

func findStationsGreedy(needed []string, stations map[string][]string) []string {
	tempneeded := make([]string, len(needed))
	copy(tempneeded, needed)
	// TODO: it doesn't kill if we check if all the needed states are offered by the stations
	finalStations := []string{}
	// go until we find all the needed states
	for len(tempneeded) != 0 {
		// loop through every station to find the one that covers the most states in needed
		var bestStation string
		stateCovered := []string{}
		for station, states := range stations {
			covered, n := slicesIntersection(states, tempneeded)
			if n > len(stateCovered) {
				bestStation = station
				stateCovered = covered
			}
		}

		// remove added states from needed
		for _, i := range stateCovered {
			for index, j := range tempneeded {
				if i == j {
					tempneeded = append(tempneeded[:index], tempneeded[index+1:]...)
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
