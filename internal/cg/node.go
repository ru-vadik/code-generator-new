package cg

type node struct {
	value string
	left  *node
	right *node
}

func (p *node) find(val string) *node {
	if val == p.value {
		return p
	}

	if val < p.value && p.left != nil {
		return p.left.find(val)
	}

	if val > p.value && p.right != nil {
		return p.right.find(val)
	}

	return nil
}

func (p *node) insert(val string) {
	if val < p.value {
		if p.left != nil {
			p.left.insert(val)
		}
		p.left = &node{value: val}
	}

	if val > p.value {
		if p.right != nil {
			p.right.insert(val)
		}
		p.right = &node{value: val}
	}
}
