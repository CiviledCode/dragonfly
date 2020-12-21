package pathfinding

import (
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"golang.org/x/image/math/f32"
)

//TODO: Finish implementing
func Pathfind(start, end f32.Vec3, world world.World) {
	//dist := int(math.Sqrt(math.Pow(float64(start[0] - end[0]), 2) + math.Pow(float64(start[1] - end[1]), 2)))
	//mapp := NewEOMap(world, start, dist)
}

type Node struct {
	Value f32.Vec3

	Next *Node
}

type LinkedList struct {
	Start *Node

	len uint32
}

func (l *LinkedList) AppendToEnd(node *Node) {
	n := l.NodeAtEnd()
	n.Next = node
	l.len++
}

func (l *LinkedList) Length() uint32 {
	return l.len + 1
}

func (l *LinkedList) NodeAtEnd() (node *Node) {
	node = l.Start
	for {
		if node.Next != nil {
			node = node.Next
		}
		return
	}
}

func (l *LinkedList) NodeAtIndex(index uint32) (n *Node, e error) {
	if index >= l.Length() {
		return nil, fmt.Errorf("got index %d had length %d", index, l.Length())
	}

	n = l.Start
	for i := uint32(0); i < l.Length(); i++ {
		if i == index {
			return n, nil
		}
	}

	// This is hypothetically impossible
	return
}

func (l *LinkedList) AppendToBeginning(node *Node) {
	node.Next = l.Start
	l.Start = node.Next
	l.len++
}

func (l *LinkedList) AppendAtIndex(node *Node, index uint32) {
	n, e := l.NodeAtIndex(index - 1)

	if e != nil {
		return
	}

	node.Next = n.Next
	n.Next = node
	l.len++
}

func (l *LinkedList) DeleteAtIndex(index uint32) error {
	if index == 0 {
		l.Start = l.Start.Next
	}

	n, e := l.NodeAtIndex(index - 1)

	if e != nil {
		return e
	}

	o, e := l.NodeAtIndex(index)

	if e != nil {
		return e
	}

	n.Next = o.Next

	return nil
}
