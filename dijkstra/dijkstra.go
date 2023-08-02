package main

import (
	"fmt"
	"math"
)

type graph struct {
	nodes map[*node]map[*node]*edge
}

func NewGraph() *graph {
	grph := make(map[*node]map[*node]*edge)
	return &graph{grph}
}

func (g *graph) AddNode(v *node) {
	g.nodes[v] = make(map[*node]*edge)
}

func (g *graph) AddEdge(p *node, v *node, e *edge) {
	g.nodes[p][v] = e
}

func (g *graph) Print() {
	for i := range g.nodes {
		for j := range g.nodes[i] {
			fmt.Printf("(src: %s, dst: %s, weigth: %d)\n", i.name, j.name, g.nodes[i][j].w)
		}
	}

}

// func (g *graph) PrintNeighbor(n *node) {
// 	for _, edges := range g.gr[n] {
// 		fmt.Println(edges.dst.name)
// 	}
// }

func (g *graph) Dijkstra(start *node) *[]node {
	costs := make(map[*node]float64)
	parents := make(map[*node]*node)
	infinity := math.Inf(1)

	for v := range g.nodes {
		if _, ok := g.nodes[start][v]; ok {
			parents[v] = start
			costs[v] = float64(g.nodes[start][v].w)
		} else {
			parents[v] = nil
			costs[v] = infinity
		}
	}

	var processed []*node
	nodeCheck := findLowestCostNode(costs, processed)
	for nodeCheck != nil {
		cost := costs[nodeCheck]
		neighbors := g.nodes[nodeCheck]
		for n := range neighbors {
			newCost := cost + float64(g.nodes[nodeCheck][n].w)
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = nodeCheck
			}
		}
		processed = append(processed, nodeCheck)
		nodeCheck = findLowestCostNode(costs, processed)
	}
	fmt.Println("costs:")
	for i, v := range costs {
		fmt.Println(i.name, v)
	}
	fmt.Println("parents:")
	for i, v := range parents {
		if v != nil {
			fmt.Println(i.name, v.name)
		}
	}
	// TODO
	return nil
}

func findLowestCostNode(costs map[*node]float64, processed []*node) *node {
	lowestCost := math.Inf(1)
	var lowestCostNode *node
	for node := range costs {
		cost := costs[node]
		found := false
		for _, n := range processed {
			if n == node {
				found = true
			}
		}
		if cost < lowestCost && !found {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

// func (g *graph) Weigth(p *node, n *node) (uint, error) {
// 	for _, e := range g.gr[p] {
// 		if e.dst == n {
// 			return e.w, nil
// 		}
// 	}
// 	err := errors.New("not an edge")
// 	return 0, err
// }

// this prints every single path from node n in a DAG
func (g *graph) PrintAllThePathsNode(n *node, path []*node) {
	if len(g.nodes[n]) == 0 {
		if len(path) != 0 {
			for _, i := range path {
				fmt.Printf("%s ", i.name)
			}
			fmt.Printf("%s\n", n.name)
		}
		return
	}
	path = append(path, n)
	for nei := range g.nodes[n] {
		g.PrintAllThePathsNode(nei, path)
	}
}

// func (g *graph) IsParent(n *node, p *node) bool {
// 	for _, t := range g.parents[n] {
// 		if t == p {
// 			return true
// 		}
// 	}
// 	return false
// }

type node struct {
	name string
}

func NewNode(name string) *node {
	return &node{name}
}

type edge struct {
	w uint
}

func NewEdge(w uint) *edge {
	return &edge{w}
}

func main() {
	start := NewNode("START")
	a := NewNode("A")
	b := NewNode("B")
	fin := NewNode("FIN")

	e1 := NewEdge(6)
	e2 := NewEdge(1)
	e3 := NewEdge(3)
	e4 := NewEdge(2)
	e5 := NewEdge(5)

	g := NewGraph()
	g.AddNode(start)
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(fin)

	g.AddEdge(start, a, e1)
	g.AddEdge(a, fin, e2)
	g.AddEdge(b, a, e3)
	g.AddEdge(start, b, e4)
	g.AddEdge(b, fin, e5)

	// g.Print()
	// g.PrintAllThePathsNode(start, nil)
	g.Dijkstra(start)

}
