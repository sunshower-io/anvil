package tree 

import core "github.com/sunshower-io/anvil/collections"

type color int

const (
	red color = iota
	black
)

type node struct {
	key    core.Key
	value  core.Value
	parent *node
	left   *node
	right  *node
	color  color
}

func (n *node) Key() core.Key {
    return n.key
}

func (n *node) Value() core.Value {
    return n.value
}

var leftMost *node

func (n *node) next() *node {
	if n.right != nil {
		r := n.right
		for r.left != nil {
			r = r.left
		}
		return r
	}

	for n != nil {
		r := n.parent
		if r == nil {
			return nil
		}

		if n.isLeft() {
			return r
		}
		n = r
	}
	return nil
}

func (n *node) isLeft() bool {
	return n == n.parent.left
}

func (n *node) isRight() bool {
	return n == n.parent.right
}

func (n *node) sibling() *node {
	if n.isLeft() {
		return n.parent.right
	}
	return n.parent.left
}
