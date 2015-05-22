package tree

import (
	"Graph/graph"
	"fmt"
	"time"
)

func DoBenchMark() {
	var roads, size, err = readGraph()
	if err != nil {
		fmt.Println("Illegal Input")
		return
	}
	fmt.Println("Read: OK")
	var table = transform(roads, size)
	fmt.Println("Prepare: OK")

	var start = time.Now()
	var ret, fail = Kruskal(roads, size)
	var tm = time.Since(start)
	if fail {
		fmt.Println("fail")
	} else {
		fmt.Printf("Kruskal %v: %v\n", tm, ret)
	}

	start = time.Now()
	ret, fail = Prim(table)
	tm = time.Since(start)
	if fail {
		fmt.Println("fail")
	} else {
		fmt.Printf("Prim %v: %v\n", tm, ret)
	}
}

func readGraph() (roads []graph.Edge, size int, err error) {
	var total int
	_, err = fmt.Scan(&size, &total)
	if err != nil || size < 2 || size > total {
		return []graph.Edge{}, 0, err
	}
	roads = make([]graph.Edge, total)
	for i := 0; i < total; i++ {
		_, err = fmt.Scan(&roads[i].A, &roads[i].B, &roads[i].Dist)
		if err != nil {
			return []graph.Edge{}, 0, err
		}
	}
	return roads, size, nil
}

func transform(roads []graph.Edge, size int) [][]graph.Path {
	var table = make([][]graph.Path, size)
	for _, path := range roads {
		table[path.A] = append(table[path.A], graph.Path{Next: path.B, Dist: path.Dist})
	}
	return table
}
