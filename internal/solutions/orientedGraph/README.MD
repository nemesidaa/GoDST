***Ориентированный граф зависимостей***

Методы:

```go
type Graph struct { // size=40 (0x28)
    Nodes  []*GraphNode
    buffer map[int][]int
    state  State
}
func (g *Graph) AddNode(id int, time int, deps ...int) *GraphNode
func (g *Graph) _State() bool // incapsulated
func (g *Graph) _UpdateState() //incapsulated
func (g *Graph) MaxNodeSTime() (int, error)
func (g *Graph) PopNode(id int) *GraphNode
```

Usage:

```go
reader := bufio.NewReader(os.Stdin)

	// Чтение количества процессов
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))
	// Инициализация графа
	graph := orientedgraph.NewGraph()
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
		// добавление узла
		graph.AddNode(i, t, deps...)

	}
	// вывод алги поиска
	fmt.Println(graph.MaxNodeSTime())
```


Структура узла:

```go
type GraphNode struct { // size=40 (0x28)
    Id           int
    Time         int
    Dependencies []*GraphNode
}
func (g *GraphNode) TimeToReach() int
```


Для получения величины пути n-ного узла графа делайте что-то типо этого:

```go
nodeTime := <*Graph>.Nodes[<n>].TimeToReach()
```
