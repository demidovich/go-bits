package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Value(t *testing.T) {
	l := New[string]()
	n := l.PushHead("a")

	assert.Equal(t, "a", n.Value())
}

func Test_PushHead(t *testing.T) {
	l := New[string]()
	push1 := l.PushHead("a")
	push2 := l.PushHead("b")

	node2, ok2 := l.Head()
	assert.Equal(t, push2, node2)
	assert.True(t, ok2)

	node1, ok1 := node2.Next()
	assert.Equal(t, push1, node1)
	assert.True(t, ok1)

	_, ok3 := node1.Next()
	assert.False(t, ok3)
}

func Test_PushTail(t *testing.T) {
	l := New[string]()
	push1 := l.PushTail("a")
	push2 := l.PushTail("b")
	push3 := l.PushTail("c")

	node3, ok3 := l.Tail()
	assert.Equal(t, push3, node3)
	assert.True(t, ok3)

	node2, ok2 := node3.Prev()
	assert.Equal(t, push2, node2)
	assert.True(t, ok2)

	node1, ok1 := node2.Prev()
	assert.Equal(t, push1, node1)
	assert.True(t, ok1)

	_, okMissing := node1.Prev()
	assert.False(t, okMissing)
}

func Test_MoveToHead_OneNode(t *testing.T) {
	l := New[string]()
	n1 := l.PushTail("a")
	l.MoveToHead(n1)

	f1, _ := l.Head()
	assert.Equal(t, n1, f1)
}

func Test_MoveToHead_ThreeNode(t *testing.T) {
	l := New[string]()
	n1 := l.PushTail("a")
	n2 := l.PushTail("b")
	n3 := l.PushTail("c")
	l.MoveToHead(n2)

	f1, ok1 := l.Head()
	assert.True(t, ok1)
	assert.Equal(t, n2, f1)

	f2, ok2 := f1.Next()
	assert.True(t, ok2)
	assert.Equal(t, n1, f2)

	f3, ok3 := f2.Next()
	assert.True(t, ok3)
	assert.Equal(t, n3, f3)
}

func Test_MoveTailToHead_ThreeNode(t *testing.T) {
	l := New[string]()
	n1 := l.PushTail("a")
	n2 := l.PushTail("b")
	n3 := l.PushTail("c")
	l.MoveToHead(n3) // В предыдущем тесте переносили средний узел n2

	f1, ok1 := l.Head()
	assert.True(t, ok1)
	assert.Equal(t, n3, f1)

	f2, ok2 := f1.Next()
	assert.True(t, ok2)
	assert.Equal(t, n1, f2)

	f3, ok3 := f2.Next()
	assert.True(t, ok3)
	assert.Equal(t, n2, f3)
}

func Test_MoveToTail_OneNode(t *testing.T) {
	l := New[string]()
	n1 := l.PushTail("a")
	l.MoveToTail(n1)

	f1, _ := l.Head()
	assert.Equal(t, n1, f1)
}

func Test_MoveToTail_ThreeNode(t *testing.T) {
	l := New[string]()
	n1 := l.PushTail("a")
	n2 := l.PushTail("b")
	n3 := l.PushTail("c")
	l.MoveToTail(n2)

	f1, ok1 := l.Head()

	assert.True(t, ok1)
	assert.Equal(t, n1, f1)

	f2, ok2 := f1.Next()
	assert.True(t, ok2)
	assert.Equal(t, n3, f2)

	f3, ok3 := f2.Next()
	assert.True(t, ok3)
	assert.Equal(t, n2, f3)
}

func Test_MoveHeadToTail_ThreeNode(t *testing.T) {
	l := New[string]()
	n1 := l.PushTail("a")
	n2 := l.PushTail("b")
	n3 := l.PushTail("c")
	l.MoveToTail(n1) // В предыдущем тесте переносили средний узел n2

	f1, ok1 := l.Head()

	assert.True(t, ok1)
	assert.Equal(t, n2, f1)

	f2, ok2 := f1.Next()
	assert.True(t, ok2)
	assert.Equal(t, n3, f2)

	f3, ok3 := f2.Next()
	assert.True(t, ok3)
	assert.Equal(t, n1, f3)
}
