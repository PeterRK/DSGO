package sort

import (
	ll "DSGO/linkedlist"
	"golang.org/x/exp/constraints"
	"unsafe"
)

// 基数排序，不依赖比较操作，具有稳定性
func RadixSort[T constraints.Integer](head *ll.Node[T]) *ll.Node[T] {
	byteWidth := uint(unsafe.Sizeof(T(0)))
	bitWidth := byteWidth * 8
	base := T(T(1) << (bitWidth - 1))
	signed := (base >> bitWidth) != 0

	if signed {
		for node := head; node != nil; node = node.Next {
			node.Val += base
		}
	}
	const bk = 1 << 8
	parts, tails := new([bk]*ll.Node[T]), new([bk]*ll.Node[T])

	for step := uint(0); step < bitWidth; step += 8 {
		for i := 0; i < bk; i++ {
			parts[i] = nil
			tails[i] = ll.FakeHead(&parts[i])
		}
		for node := head; node != nil; node = node.Next {
			radix := uint8(node.Val >> step)
			tails[radix].Next, tails[radix] = node, node
		}
		tail := ll.FakeHead(&head)
		for i := 0; i < bk; i++ {
			if parts[i] != nil {
				tail.Next = parts[i]
				tail = tails[i]
			}
		}
		tail.Next = nil
	}

	if signed {
		for node := head; node != nil; node = node.Next {
			node.Val -= base
		}
	}
	return head
}
