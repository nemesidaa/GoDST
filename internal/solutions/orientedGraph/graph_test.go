package orientedgraph_test

import (
	orientedgraph "GoDST/internal/solutions/orientedGraph"
	"fmt"
	"testing"
)

// -benchmem -run=^$ -bench
// BenchmarkGraph          20602             59520 ns/op              48 B/op          1 allocs/op
func BenchmarkGraph(b *testing.B) {
	g := orientedgraph.NewGraph()
	g.AddNode(1, 10)
	g.AddNode(2, 8, 3, 4)
	g.AddNode(3, 15, 5)
	g.AddNode(4, 17)
	g.AddNode(5, 1)

	for i := 0; i < b.N; i++ {
		fmt.Println(g.MaxNodeSTime())
	}

}
