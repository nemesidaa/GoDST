package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GraphNode struct {
	Id           int
	Time         int
	Dependencies []*GraphNode
	Indegree     int
	FinishTime   int
}

type Graph struct {
	Nodes  []*GraphNode
	buffer map[int][]int // нужно чтобы хранить неинициализированные депы, id[]id, типо чтобы создать своего рода воткни и забей и не мучаться о порядке вставки, сохранить айдишник
}

// ВАЖНО!!! СУПЕРКОСТЫЛЬ, было бы больше времени, сэкономил бы по time complexity через реаллокацию
func removeElement(slice []int, value int) []int {
	result := []int{}
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

func NewGraph() *Graph {
	return &Graph{
		Nodes:  make([]*GraphNode, 0),
		buffer: make(map[int][]int),
	}
}

func (g *Graph) AddNode(id, time int, deps ...int) *GraphNode {
	g.buffer[id] = make([]int, 0)
	node := &GraphNode{
		Id:           id,
		Time:         time,
		Dependencies: make([]*GraphNode, 0),
	}
	for _, dep := range deps {
		g.buffer[id] = append(g.buffer[id], dep)
		for _, n := range g.Nodes {
			if dep == n.Id {
				node.Dependencies = append(node.Dependencies, n)
				g.buffer[id] = removeElement(g.buffer[id], dep)
			}
		}
	}
	// закрыли вопрос с вставляемой, теперь к старым рофланам через перебор
	for n, deps := range g.buffer {
		for _, dep := range deps {
			if node.Id == dep {
				g.Nodes[n-1].Dependencies = append(g.Nodes[n-1].Dependencies, node)
				g.buffer[n] = removeElement(g.buffer[n], dep)
			}
		}
	}
	// проверка на наличие устаревшей инфы в буфере
	g.Nodes = append(g.Nodes, node)
	return node
}

func (g *Graph) Solution() int {
	n := len(g.Nodes)
	if n == 0 {
		return 0
	}

	queue := []*GraphNode{}
	for _, node := range g.Nodes {
		if node.Indegree == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		finishTime := node.Time
		for _, dep := range node.Dependencies {
			if dep.FinishTime > finishTime {
				finishTime = dep.FinishTime
			}
		}
		node.FinishTime = finishTime

		for _, neighbor := range g.Nodes {
			for _, dep := range neighbor.Dependencies {
				if dep == node {
					neighbor.Indegree--
					if neighbor.Indegree == 0 {
						queue = append(queue, neighbor)
					}
				}
			}
		}
	}

	maxFinishTime := 0
	for _, node := range g.Nodes {
		if node.FinishTime > maxFinishTime {
			maxFinishTime += node.FinishTime
		}
	}

	return maxFinishTime
}

func Sixth() {
	reader := bufio.NewReader(os.Stdin)

	// Чтение количества процессов
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	graph := NewGraph()
	for i := 1; i <= n; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Split(line, " ")

		t, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			panic(err)
		}

		deps := make([]int, 0)

		for _, v := range parts[1:] {
			d, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				panic(err)
			}
			deps = append(deps, d)
		}
		graph.AddNode(i, t, deps...)

	}

	fmt.Println(graph.Solution())
}
