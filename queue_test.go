package go_queue

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue("2")
	q.Enqueue(3)

	q.RemoveByValue(2)

	q.Traverse(func(n *Node) {
		fmt.Println(n.Value)
	})
}
