package main

import "fmt"

type graph struct {
	gr map[*node][]edge
	v  int
	e  int
}

func NewGraph() *graph {
	g := make(map[*node][]edge)
	return &graph{g, 0, 0}
}

func (g *graph) AddNode(v *node) {
	g.gr[v] = make([]edge, 0)
}

func (g *graph) AddEdge(v *node, e *edge) {
	g.gr[v] = append(g.gr[v], *e)
}

func (g *graph) Print() {
	for v, e := range g.gr {
		for _, edges := range e {
			fmt.Printf("(src: %s, dst: %s, weigth: %d)\n", v.name, edges.dst.name, edges.w)
		}
	}
}

func (g *graph) PrintNeighbor(n *node) {
	for _, edges := range g.gr[n] {
		fmt.Println(edges.dst.name)
	}
}

// func (g *graph) Dijkstra(start *node) *[]node {
// 	fmt.Printf("%v\n", g.v)
// 	return nil
// }

// this prints every single path from node n in a DAG
func (g *graph) PrintAllThePathsNode(n *node, path []*node) {
	if len(g.gr[n]) == 0 {
		if len(path) != 0 {
			for _, i := range path {
				fmt.Printf("%s ", i.name)
			}
			fmt.Printf("%s\n", n.name)
		}
		return
	}
	path = append(path, n)
	for _, edges := range g.gr[n] {
		g.PrintAllThePathsNode(edges.dst, path)
	}
}

type node struct {
	name string
}

func NewNode(name string) *node {
	return &node{name}
}

type edge struct {
	w   uint
	dst *node
}

func NewEdge(w uint, dst *node) *edge {
	return &edge{w, dst}
}

func main() {
	start := NewNode("START")
	a := NewNode("A")
	b := NewNode("B")
	fin := NewNode("FIN")

	e1 := NewEdge(6, a)
	e2 := NewEdge(1, fin)
	e3 := NewEdge(3, a)
	e4 := NewEdge(2, b)
	e5 := NewEdge(5, fin)

	g := NewGraph()
	g.AddNode(start)
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(fin)

	g.AddEdge(start, e1)
	g.AddEdge(a, e2)
	g.AddEdge(b, e3)
	g.AddEdge(start, e4)
	g.AddEdge(b, e5)

	g.PrintAllThePathsNode(b, nil)

}
