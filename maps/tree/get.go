package tree 

import core "github.com/sunshower-io/anvil/collections"

func (t *TreeMap) Get(key core.Value) core.Value {
	node, match := t.greaterThanOrEqualTo(key)
	if match {
		return node.value
	}
	return nil
}

func (t *TreeMap) greaterThanOrEqualTo(key core.Value) (*node, bool) {
	n := t.root
	for {
		if n == nil {
			return nil, false
		}

		c := t.compare(key, n.key)

		if c == 0 {
			return n, true
		} else if c < 0 {
			if n.left != nil {
				n = n.left
			} else {
				return n, false
			}
		} else {
			if n.right != nil {
				n = n.right
			} else {
				s := n.next()
				if s == nil {
					return nil, false
				} else {
					c = t.compare(key, s.key)
					return s, c == 0
				}
			}
		}
	}
}
