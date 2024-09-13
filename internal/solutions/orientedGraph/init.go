package orientedgraph

type GraphNode struct {
	Id           int
	Time         int
	Dependencies []*GraphNode
}

// А после этого не захочет жить даже процессоры IBM...
func (g *GraphNode) TimeToReach() int {
	var t int
	if len(g.Dependencies) == 0 {
		t = g.Time
	} else {
		for _, dep := range g.Dependencies {
			t = max(t, dep.TimeToReach()+g.Time)
		}
	}
	return t
}

type Graph struct {
	Nodes  []*GraphNode
	buffer map[int][]int // нужно чтобы хранить неинициализированные депы, id[]id, типо чтобы создать своего рода воткни и забей и не мучаться о порядке вставки, сохранить айдишник
	state  State
}

func NewGraph() *Graph {
	return &Graph{
		Nodes:  make([]*GraphNode, 0),
		buffer: make(map[int][]int),
		state:  Ready,
	}
}

func (g *Graph) _UpdateState() {
	if len(g.buffer) > 0 || len(g.Nodes) == 0 {
		g.state = NotReady
	}
	g.state = Ready
}

func (g *Graph) _State() bool {
	return g.state == Ready
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
	g._UpdateState()
	return node
}

func (g *Graph) PopNode(id int) *GraphNode {
	var node *GraphNode
	if g._State() {
		for idx, n := range g.Nodes {
			if n.Id == id {
				node = n
				g.Nodes = append(g.Nodes[:idx], g.Nodes[idx+1:]...)
				break
			}
		}
		for n, deps := range g.buffer {
			for _, dep := range deps {
				if node.Id == dep {
					g.Nodes[n-1].Dependencies = removeElement(g.Nodes[n-1].Dependencies, node)
					g.buffer[n] = removeElement(g.buffer[n], dep)
				}
			}
		}
		g._UpdateState()
	}
	return node
}

func (g *Graph) MaxNodeSTime() (int, error) {
	if g._State() {
		ttr := make([]int, len(g.Nodes))
		for idx, n := range g.Nodes {
			ttr[idx] = n.TimeToReach()
		}
		return max(ttr...), nil
	} else {
		return 0, errStateNotReady
	}
}
