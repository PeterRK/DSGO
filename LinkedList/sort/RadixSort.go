package sort

import (
	"LinkedList/list"
	"unsafe"
)

//基数排序，不依赖比较操作，具有稳定性
func RadixSort(head *list.Node) *list.Node {
	const bytesOfUint = uint(unsafe.Sizeof(uint(0)))
	const base = -int((^uint(0))>>1) - 1
	for node := head; node != nil; node = node.Next {
		node.Val += base
	}

	var parts, tails [256]*list.Node
	for step := uint(0); step < bytesOfUint*8; step += 8 {
		for i := 0; i < 256; i++ {
			parts[i] = nil
			tails[i] = list.FakeHead(&parts[i])
		}
		for node := head; node != nil; node = node.Next {
			var radix = uint8((node.Val >> step) & 0xFF)
			tails[radix].Next, tails[radix] = node, node
		}
		var tail = list.FakeHead(&head)
		for i := 0; i < 256; i++ {
			if parts[i] != nil {
				tail.Next = parts[i]
				tail = tails[i]
			}
		}
		tail.Next = nil
	}

	for node := head; node != nil; node = node.Next {
		node.Val -= base
	}
	return head
}
