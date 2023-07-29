package main

import (
	"breadthfirst/queue"
	"fmt"
	"log"
)

func breadthSearch(graph map[string][]string, name string) bool {
	q := queue.Init()
	searched := []string{}
	for _, p := range graph[name] {
		q.EnQueue(p)
		searched = append(searched, p)
	}

	var person string
	var err error
	for q.Length() != 0 {
		person, err = q.DeQueue()
		if err != nil {
			log.Fatal(err)
		}

		if personIsSeller(person) {
			fmt.Println("seller is", person)
			return true
		} else {
			for _, p := range graph[person] {
				if isSearched(searched, p) {
					continue
				}
				searched = append(searched, person)
				q.EnQueue(p)
			}
		}
	}
	return false
}

func isSearched(searched []string, person string) bool {
	for _, i := range searched {
		if i == person {
			return true
		}
	}
	return false
}

// dummy function to pick someone that their name ends with 6
// TODO: when the human type is implemented, this neends to be changed too
func personIsSeller(person string) bool {
	return person[len(person)-1] == '6'
}

func main() {
	graph := make(map[string][]string)
	graph["person1"] = []string{"person2", "person3", "person4", "person5"}
	graph["person2"] = []string{"person5", "person12"}
	graph["person3"] = []string{"person6"}
	graph["person5"] = []string{"person4"}

	fmt.Println(breadthSearch(graph, "person1"))

}

// TODO: Make the queue general for any type
// TODO: define a type named human with name and occupation fileds and pass that object to the graph queue
