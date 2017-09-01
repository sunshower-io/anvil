package maps

import "github.com/sunshower-io/anvil/core"

func (t *TreeMap) Remove(key core.Value) bool {
	i := t.collectGreaterThanOrEqualTo(key)
	if i.node != nil {
		t.RemoveAll(i)
		return true
	}
	return false
}

func (t *TreeMap) RemoveAll(i Iterator) {
	t.remove(i.node)
}

func (t *TreeMap) remove(n *node) {
	if !(n.left == nil || n.right == nil) {
		p := greatestPredecessor(n)
		t.swap(n, p)
	}

	c := n.right

	if c == nil {
		c = n.left
	}

	if n.color == black {
		n.color = colorOf(c)
		t.deleteFirstCase(n)
	}

	t.replace(n, c)

	if n.parent == nil && c != nil {
		c.color = black
	}

	t.size -= 1

	if t.size == 0 {
		t.minimum = nil
		t.maximum = nil
	} else {
		if t.minimum == n {
			t.recomputeSmallest()
		}
		if t.maximum == n {
			t.recomputeLargest()
		}
	}
}

func (t *TreeMap) deleteFirstCase(n *node) {
	for {
		if n.parent != nil {
			s := n.sibling()
			if colorOf(s) == red {
				n.parent.color = red
				s.color = black

				if n == n.parent.left {
					t.rotateLeft(n.parent)
				} else {
					t.rotateRight(n.parent)
				}
			}
			s = n.sibling()

			if colorOf(n.parent) == black &&
				colorOf(s) == black &&
				colorOf(s.left) == black &&
				colorOf(s.right) == black {
				n.sibling().color = red
				n = n.parent
				continue
			} else {
				if colorOf(n.parent) == red &&
					colorOf(s) == black &&
					colorOf(s.left) == black &&
					colorOf(s.right) == black {
					s.color = red
					n.parent.color = black
				} else {
					t.deleteFifthCase(n, n.sibling())
				}
			}
		}
		break
	}
}

func (t *TreeMap) deleteFifthCase(n, s *node) {

	if n == n.parent.left &&
		colorOf(s) == black &&
		colorOf(s.left) == red &&
		colorOf(s.right) == black {
		s.color = red
		s.left.color = black
		t.rotateRight(s)
	} else if n == n.parent.right &&
		colorOf(s) == black &&
		colorOf(s.right) == red &&
		colorOf(s.left) == black {
		s.color = red
		s.right.color = black
		t.rotateLeft(s)
	}
	t.deleteSixthCase(n, n.sibling())
}

func (t *TreeMap) deleteSixthCase(n, s *node) {
	s.color = colorOf(n.parent)
	n.parent.color = black

	if n == n.parent.left {
		s.right.color = black
		t.rotateLeft(n.parent)
	} else {
		s.left.color = black
		t.rotateRight(n.parent)
	}
}

func (t *TreeMap) recomputeLargest() {
	t.maximum = t.root
	if t.maximum != nil {
		for t.maximum.right != nil {
			t.maximum = t.maximum.right
		}
	}
}

func (t *TreeMap) recomputeSmallest() {
	t.minimum = t.root

	if t.minimum != nil {
		for t.minimum.left != nil {
			t.minimum = t.minimum.left
		}
	}
}

func (t *TreeMap) collectGreaterThanOrEqualTo(k core.Value) Iterator {
	n, _ := t.greaterThanOrEqualTo(k)
	return Iterator{tree: t, node: n}
}

func (t *TreeMap) replace(old, new *node) {
	if old.parent == nil {
		t.root = new
	} else {
		if old == old.parent.left {
			old.parent.left = new
		} else {
			old.parent.right = new
		}
	}
	if new != nil {
		new.parent = old.parent
	}
}

func (t *TreeMap) swap(n *node, p *node) {

	left := p.isLeft()
	tmp := *p

	t.replace(n, p)
	p.color = n.color

	if tmp.parent == n {
		swapLeft(left, p, n, tmp)
	} else {
		swapRight(p, n, left, tmp)
	}
	n.color = tmp.color

}
func swapLeft(left bool, p *node, n *node, tmp node) {
	if left {
		p.left = n
		p.right = n.right

		if p.right != nil {
			p.right.parent = p
		}
	} else {
		p.left = n.left

		if p.left != nil {
			p.left.parent = p
		}
		p.right = n
	}
	n.key = tmp.key
	n.value = tmp.value
	n.parent = p
	n.left = tmp.left
	if n.left != nil {
		n.left.parent = n
	}
	n.right = tmp.right
	if n.right != nil {
		n.right.parent = n
	}
}
func swapRight(p *node, n *node, left bool, tmp node) {
	p.left = n.left
	if p.left != nil {
		p.left.parent = p
	}
	p.right = n.right
	if p.right != nil {
		p.right.parent = p
	}
	if left {
		tmp.parent.left = n
	} else {
		tmp.parent.right = n
	}
	n.key = tmp.key
	n.value = tmp.value
	n.left = tmp.left
	if n.left != nil {
		n.left.parent = n
	}
	n.right = tmp.right
	if n.right != nil {
		n.right.parent = n
	}
}

func greatestPredecessor(n *node) *node {
	m := n.left
	for l := m.right; l != nil; {
		m = l.right
	}
	return m
}

func colorOf(n *node) color {
	if n == nil {
		return black
	}
	return n.color
}
