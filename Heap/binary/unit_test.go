package heap

import (
	"testing"
)

func Test_Nothing(t *testing.T) {
	//Do Nothing
}

/*

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const size = 16
	rand.Seed(time.Now().Unix())

	var list [size]int
	for i := 0; i < size; i++ {
		var key = rand.Int() % 1000
		fmt.Printf("%d ", key)
		list[i] = key
	}

	var heap BinaryHeap
	heap.Build(list[:])
	for i := 0; i < size; i++ {
		var key = rand.Int() % 1000
		fmt.Printf("%d ", key)
		heap.Push(key)
	}
	fmt.Println()

	for !heap.IsEmpty() {
		fmt.Printf("%d ", heap.Pop())
	}
	fmt.Println()

}
*/
