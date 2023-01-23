package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const MaxInt = int((^uint(0)) >> 1)

var (
	nodIni = 1
)

type Edge struct {
	label int
}

type Node struct {
	label     int // nombre del nodo
	visited   int // 0 not visited, -1 gray, 1 visited
	previo    int // indice a nodo anterior en recorrido
	dist      int // distancia al nodo inicial
	neighbors []Edge
}

type Graph struct {
	nodes []*Node
}

type Data struct {
	source int
	dest   int
}

func check(st string, e error) {
	if e != nil {
		fmt.Println(st)
		panic(e)
	}
}

func enlista(pos int, ls []int) (res bool) {
	for _, dat := range ls {
		if dat*1_000 == pos {
			res = true
			break
		}
	}
	return
}

func presentar(pg *Graph, nd int, start time.Time) {
	fmt.Printf("%5.5s %9.9s %8.8s \n", "Nodo", "Distancia", "Tiempo")
	elapsed := time.Since(start)
	dist := pg.nodes[nd].dist
	fmt.Printf("%05d    %06d    %s\n", nd, dist, elapsed)
	for nd > 1 {
		fmt.Print(nd, pg.nodes[nd].dist, " -> ")
		nd = pg.nodes[nd].previo
	}
	fmt.Println(nd)
}

func iniGraph(tsize int) (g Graph) {
	g.nodes = make([]*Node, tsize+1)
	return
}

func (g *Graph) insertNode(data Data) {
	var edge Edge
	var node Node

	// Definir como -1 para evitar errores
	node.previo = -1

	edge.label = data.dest
	i := data.source
	node.label = i
	if g.nodes[i] == nil {
		g.nodes[i] = &node
	}
	if g.nodes[edge.label] == nil {
		g.nodes[edge.label] = new(Node)
		g.nodes[edge.label].label = edge.label
	}
	g.nodes[i].neighbors = append(g.nodes[i].neighbors, edge)
	return
}

func (g *Graph) readFile(name string) (nrnod, nrlin int) {
	file, err := os.Open(name)
	check("No se abrio archivos ", err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nrnod, &nrlin)
	g.nodes = make([]*Node, nrnod+1)

	for scanner.Scan() {
		lineStr := scanner.Text()
		arr := strings.Fields(lineStr)
		inic, _ := strconv.Atoi(arr[0])

		node := Node{label: inic}
		g.nodes[inic] = &node
		for _, sj := range arr[1:] {
			j, _ := strconv.Atoi(sj)
			edge := Edge{label: j}
			node.neighbors = append(node.neighbors, edge)
		}
	}
	return
}

func (g *Graph) display() {
	s := ""
	for i := 1; i < len(g.nodes); i++ {
		near := g.nodes[i]
		if near == nil {
			continue
		}
		s += fmt.Sprintf(" %5d --> ", near.label)
		for _, edge := range near.neighbors {
			//             s += fmt.Sprintf(" %5d %8d | ",edge.label , edge.distance)
			s += fmt.Sprintf(" %5d | ", edge.label)
		}
		fmt.Println(s)
		s = ""
	}
}

func (g *Graph) depthSearch(index int, rollback bool) {
	if index == 0 {
		return
	}

	node := g.nodes[index]

	if !rollback {
		g.nodes[index].visited = 1
		fmt.Println(node.label)
	}

	// Look for the next node
	minimum := MaxInt
	for _, neighbor := range node.neighbors {
		if neighbor.label < minimum && g.nodes[neighbor.label].visited != 1 {
			minimum = neighbor.label
		}
	}

	if minimum == MaxInt {
		// Do not change nodesVisited
		g.depthSearch(node.previo, true)
		return
	}

	// Continue to the next min node
	g.nodes[minimum].previo = index
	g.depthSearch(minimum, false)

	return
}

// Depth search with min
func (g *Graph) search() {
	g.depthSearch(nodIni, false)
}

func main() {
	start := time.Now()
	name := "./data/data_graf.txt"
	if len(os.Args) > 1 {
		name = "./data/" + os.Args[1]
	}
	fmt.Println("main Inic")
	var gr Graph
	(&gr).readFile(name)
	elapsed := time.Since(start)
	fmt.Printf("ReadFile time %s \n", elapsed)
	//(&gr).display()

	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)

	gr.search()
}
