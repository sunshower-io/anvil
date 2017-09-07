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

type treemapIterator struct {
    collections.Iterator
    
    
    node *node
    tree *TreeMap
}

func (i *treemapIterator) HasNext() bool {
	return i.node != nil
}

func (i *treemapIterator) First() bool {
	return i.node == i.tree.minimum
}

func (i *treemapIterator) Last() bool {
	return i.node == i.tree.maximum
}

func (i *treemapIterator) NextKey() interface{} {
	return i.node.key

}

func (i *treemapIterator) NextValue() interface{} {
	return i.node.value
}

func (i *treemapIterator) NextEntry() Pair {
	return Pair{
		Key:   i.node.key,
		Value: i.node.value,
	}
}

func (i *treemapIterator) Next() (collections.Value, error) {
    r := i.node
    c := r.next()
    i.node = c
    return r, nil
}
