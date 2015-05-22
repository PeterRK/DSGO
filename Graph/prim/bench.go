package prim

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
	var table = transform(roads, size)

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

func readGraph() (roads []graph.PathX, size int, err error) {
	var total int
	_, err = fmt.Scan(&size, &total)
	if err != nil || size < 2 || size > total {
		return []graph.PathX{}, 0, err
	}
	roads = make([]graph.PathX, total)
	for i := 0; i < total; i++ {
		_, err = fmt.Scan(&roads[i].A, &roads[i].B, &roads[i].Dist)
		if err != nil {
			return []graph.PathX{}, 0, err
		}
	}
	return roads, size, nil
}

func transform(roads []graph.PathX, size int) [][]graph.Path {
	var table = make([][]graph.Path, size)
	for _, path := range roads {
		table[path.A] = append(table[path.A], graph.Path{Next: path.B, Dist: path.Dist})
	}
	return table
}
