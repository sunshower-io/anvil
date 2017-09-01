package tree 

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
    "github.com/sunshower-io/anvil/collections"
)

var IllegalState = errors.New("Illegal state")



type TreeMap struct {
    collections.SortedMap
	size int

	root *node

	minimum *node
    
	maximum *node

	comparator collections.Comparator
}



func NewTreeMap(f func(
        collections.Value, 
        collections.Value,
) int) *TreeMap {
	return &TreeMap{
		size:    0,
		root:    nil,
		minimum: nil,
		maximum: nil,
		comparator: &collections.FunctionComparator{
			Ord: f,
		},
	}
}

func (t *TreeMap) FirstValue() collections.Value {
	if t.size >= 0 {
		return t.minimum.value
	}
	return nil
}

func (t *TreeMap) FirstKey() collections.Key {
	if t.size >= 0 {
		return t.minimum.key
	}
	return nil
}

func (t *TreeMap) Iterator() Iterator {
	return Iterator{tree: t, node: t.minimum}
}

func (t *TreeMap) IsEmpty() bool {
	return t.size == 0
}

func (t *TreeMap) Size() int {
	return t.size
}

func (t *TreeMap) String() string {

	b := new(bytes.Buffer)

	write(t.root, 0, b)
	return b.String()
}

func write(n *node, depth int, b *bytes.Buffer) {
	if n != nil {
		b.WriteString(strings.Repeat(" ", depth))
		b.WriteString(fmt.Sprintf("Node{key: %s, value: %s", n.key, n.value))
		b.WriteString("\n")
		write(n.left, depth+1, b)
		write(n.right, depth+1, b)
	}
}

func (t *TreeMap) compare(lhs, rhs collections.Value) int {
	return t.comparator.Compare(lhs, rhs)
}

func (t *TreeMap) setMinimum(n *node) {
	if t.minimum == nil {
		t.minimum = n
		t.maximum = n
	} else if t.compare(n.key, t.minimum.key) <= 0 {
		t.minimum = n
	}
}

func (t *TreeMap) setMaximum(n *node) {

	if t.maximum == nil {
		t.maximum = n
		t.minimum = n
	} else if t.compare(n.key, t.maximum.key) > 0 {
		t.maximum = n
	}
}
