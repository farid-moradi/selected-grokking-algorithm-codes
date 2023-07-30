package main

import (
	"breadthfirst/queue"
	"fmt"
	"log"
)

type human struct {
	name       string
	age        int
	occupation string
}

func breadthSearch(graph map[human][]human, name human) bool {
	q := queue.Init()
	searched := []human{}
	for _, p := range graph[name] {
		q.EnQueue(p)
		searched = append(searched, p)
	}

	for q.Length() != 0 {
		p, err := q.DeQueue()
		if err != nil {
			log.Fatal(err)
		}
		person, _ := p.(human)

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

func isSearched(searched []human, person human) bool {
	for _, i := range searched {
		if i == person {
			return true
		}
	}
	return false
}

func personIsSeller(person human) bool {
	return person.occupation == "seller"
}

func main() {
	graph := make(map[human][]human)
	person1 := &human{name: "person 1", age: 30, occupation: "athlete"}
	person2 := &human{name: "person 2", age: 20, occupation: "student"}
	person3 := &human{name: "person 3", age: 25, occupation: "seller"}
	graph[*person1] = []human{*person2}
	graph[*person2] = []human{*person3}

	fmt.Println(breadthSearch(graph, *person1))

}
