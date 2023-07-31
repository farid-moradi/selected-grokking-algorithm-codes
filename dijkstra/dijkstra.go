package main

type graph struct {
	v *[]node
	e *[]edge
}

func NewGraph(v *[]node, e *[]edge) *graph {
	return &graph{v, e}
}

func (g *graph) AddNode(v *node) {
	*g.v = append(*g.v, *v)
}

func (g *graph) AddEdge(e *edge) {
	*g.e = append(*g.e, *e)
}

func (g *graph) Dijkstra(start *node) *[]node {
	return nil
}

type node struct {
	name      string
	neighbors *[]node
}

func NewNode(name string) *node {
	return &node{name, nil}
}

type edge struct {
	w   uint
	src *node
	dst *node
}

func NewEdge(w uint, src *node, dst *node) *edge {
	*src.neighbors = append(*src.neighbors, *dst)
	return &edge{w, src, dst}
}

func main() {
	start := NewNode("START")
	a := NewNode("A")
	b := NewNode("B")
	fin := NewNode("FIN")

	e1 := NewEdge(6, start, a)
	e2 := NewEdge(1, a, fin)
	e3 := NewEdge(3, b, a)
	e4 := NewEdge(2, start, b)
	e5 := NewEdge(5, b, fin)

	g := NewGraph(nil, nil)
	g.AddNode(start)
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(fin)

	g.AddEdge(e1)
	g.AddEdge(e2)
	g.AddEdge(e3)
	g.AddEdge(e4)
	g.AddEdge(e5)

}
