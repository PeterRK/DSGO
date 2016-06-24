package span

import (
	"DSGO/Graph/graph"
	"fmt"
	"time"
)

func BenchMark() {
	var start = time.Now()
	var edges, size, err = readGraph() //IO就是慢！！！
	if err != nil {
		fmt.Println("Illegal Input")
		return
	}
	var roads = transform(edges, size)
	fmt.Printf("Prepare Graph [%d vertexes & %d edges] in %v\n", size, len(edges), time.Since(start))

	start = time.Now()
	ret1, err := Kruskal(edges, size)
	var tm1 = time.Since(start)
	if err != nil {
		fmt.Println(err)
	}

	start = time.Now()
	ret2, err := Prim(roads)
	var tm2 = time.Since(start)
	if err != nil {
		fmt.Println(err)
	}

	if ret1 != ret2 {
		fmt.Printf("Kruskal[%d] != Prim[%d]\n", ret1, ret2)
	} else {
		fmt.Printf("result = %d\n", ret1)
		fmt.Println("Kruskal:", tm1)
		fmt.Println("Prim:   ", tm2)
	}
}

func readGraph() (edges []graph.Edge, size int, err error) {
	var total int
	_, err = fmt.Scan(&size, &total)
	if err != nil || size < 2 || size > total {
		return nil, 0, err
	}
	edges = make([]graph.Edge, total)
	for i := 0; i < total; i++ {
		_, err = fmt.Scan(&edges[i].A, &edges[i].B, &edges[i].Weight)
		if err != nil {
			return nil, 0, err
		}
	}
	return edges, size, nil
}

func transform(edges []graph.Edge, size int) [][]graph.Path {
	var roads = make([][]graph.Path, size)
	for _, path := range edges {
		roads[path.A] = append(
			roads[path.A], graph.Path{Next: path.B, Weight: path.Weight})
	}
	return roads
}
