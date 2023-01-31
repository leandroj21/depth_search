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
	visited   bool
	previo    int // indice a nodo anterior en recorrido
	neighbors []Edge
}

type Graph struct {
	nodes        []*Node
	visitedNodes []int
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Something bad just happened")
		}
	}(file)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	_, _ = fmt.Sscanf(line, "%d %d", &nrnod, &nrlin)
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
	fmt.Printf("\nNodes visited in order:\n")
	breakLine := 0
	for _, node := range g.visitedNodes {
		fmt.Printf("%-3d  ", node)
		if breakLine == 9 {
			fmt.Println("")
			breakLine = 0
		} else {
			breakLine++
		}
	}
	fmt.Println("")
}

func (g *Graph) depthSearch(index int, rollback bool) {
	if index == 0 {
		return
	}

	node := g.nodes[index]

	if !rollback {
		g.nodes[index].visited = true

		// For printing
		g.visitedNodes = append(g.visitedNodes, index)
	}

	// Look for the next node
	minimum := MaxInt
	for _, neighbor := range node.neighbors {
		if neighbor.label < minimum && !g.nodes[neighbor.label].visited {
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

func testDepthSearch(graph *Graph) {
	for _, node := range graph.nodes {
		if node == nil {
			continue
		}

		if !node.visited {
			fmt.Printf("Error: %d missing.\n", node.label)
		}
	}
}

func printRequiredNodes(graph *Graph) {
	length := len(graph.visitedNodes)
	firstNodes := ""
	lastNodes := ""
	c, breakLine := 0, 0
	separator := " "
	for i, j := 0, length-15; i <= length-1 && j <= length-1; i, j = i+1, j+1 {
		// Printing options
		if c == 15 {
			break
		}
		if breakLine == 4 {
			separator = "\n"
			breakLine = -1
		}

		firstNodes += fmt.Sprintf("=> %-3d%s", graph.visitedNodes[i], separator)
		lastNodes += fmt.Sprintf("=> %-3d%s", graph.visitedNodes[j], separator)

		// Variables used for printing
		c++
		breakLine++
		separator = " "
	}
	fmt.Printf("First %d nodes:\n", c)
	fmt.Println(firstNodes)

	fmt.Printf("Last %d nodes:\n", c)
	fmt.Println(lastNodes)
}

func main() {
	// Reading file and inserting nodes
	start := time.Now()
	name := "./data/data_graf.txt"
	if len(os.Args) > 1 {
		name = "./data/" + os.Args[1]
	}
	var gr Graph
	(&gr).readFile(name)
	elapsed := time.Since(start)
	fmt.Println("Inserting nodes time elapsed:", elapsed)

	gr.search()

	elapsed = time.Since(start)
	fmt.Println("Finish time:", elapsed)

	// Test that all nodes are visited
	testDepthSearch(&gr)

	gr.display()

	printRequiredNodes(&gr)
}
