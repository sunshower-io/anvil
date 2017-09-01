package tree 

import (
    "errors"
    "github.com/sunshower-io/anvil/collections"
)

var NoSuchElementException = errors.New("No element")

type Pair struct {
	Key   interface{}
	Value interface{}
}

type Iterator struct {
    collections.Iterator
	node *node
	tree *TreeMap
}

func (i Iterator) HasNext() bool {
	return i.node != nil
}

func (i Iterator) First() bool {
	return i.node == i.tree.minimum
}

func (i Iterator) Last() bool {
	return i.node == i.tree.maximum
}

func (i Iterator) NextKey() interface{} {
	return i.node.key

}

func (i Iterator) NextValue() interface{} {
	return i.node.value
}

func (i Iterator) NextEntry() Pair {
	return Pair{
		Key:   i.node.key,
		Value: i.node.value,
	}
}

func (i Iterator) Next() Iterator {
	if i.HasNext() {
		return Iterator{
			tree: i.tree,
			node: i.node.next(),
		}
	}
	panic(NoSuchElementException)
}
