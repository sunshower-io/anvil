package tree 

import core "github.com/sunshower-io/anvil/collections"

func (t *TreeMap) Put(key, value core.Value) bool {
	n := t.insert(key, value)
	if n == nil {
		return false
	}
	n.color = red

	for {
		if n.parent == nil {
			n.color = black
			break
		}

		if n.parent.color == black {
			break
		}

		gp := n.parent.parent
		var u *node

		if n.parent.isLeft() {
			u = gp.right
		} else {
			u = gp.left
		}

		if u != nil && u.color == red {
			n.parent.color = black
			u.color = black
			gp.color = red
			n = gp
			continue
		}

		if n.isRight() && n.parent.isLeft() {
			t.rotateLeft(n.parent)
			n = n.left
			continue
		}

		if n.isLeft() && n.parent.isRight() {
			t.rotateRight(n.parent)
			n = n.right
			continue
		}

		n.parent.color = black
		gp.color = red
		if n.isLeft() {
			t.rotateRight(gp)
		} else {
			t.rotateLeft(gp)
		}
		break
	}
	return true
}

func (t *TreeMap) insert(key, value core.Value) *node {

	if t.root == nil {
		n := &node{key: key, value: value}
		t.root = n
		t.minimum = n
		t.maximum = n
		t.size += 1
		return n
	}

	p := t.root

	for {
		c := t.compare(key, p.key)

		if c == 0 {
			return nil
		} else if c < 0 {
			if p.left == nil {
				n := &node{key: key, value: value, parent: p}
				p.left = n
				t.size += 1
				t.setMinimum(n)
				return n
			} else {
				p = p.left
			}
		} else {
			if p.right == nil {
				n := &node{key: key, value: value, parent: p}
				p.right = n
				t.size += 1
				t.setMaximum(n)
				return n
			} else {
				p = p.right
			}
		}
	}
	panic(IllegalState)
}
