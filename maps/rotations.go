package maps 

func (t *TreeMap) rotateRight(n *node) {
	l := n.left
	n.left = l.right

	if l.right != nil {
		l.right.parent = n
	}

	l.parent = n.parent

	if n.parent == nil {
		t.root = l
	} else {
		if n.isLeft() {
			n.parent.left = l
		} else {
			n.parent.right = l
		}
	}
	l.right = n
	n.parent = l
}

func (t *TreeMap) rotateLeft(n *node) {
	r := n.right
	n.right = r.left
	if r.left != nil {
		r.left.parent = n
	}
	r.parent = n.parent

	if n.parent == nil {
		t.root = r
	} else {
		if n.isLeft() {
			n.parent.left = r
		} else {
			n.parent.right = r
		}
	}
	r.left = n
	n.parent = r
}
