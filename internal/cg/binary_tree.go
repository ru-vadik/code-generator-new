package cg

type BinaryTree struct {
	config   *Config
	planFile *PlanFile

	rootNode *node
	counter  int
}

func NewBinaryTree(config *Config, planFile *PlanFile) *BinaryTree {
	return &BinaryTree{
		config:   config,
		planFile: planFile,

		rootNode: nil,
		counter:  0,
	}
}

func (p *BinaryTree) Start() {
	defer func() {
		close(p.planFile.C)
	}()

	randomCode := NewRandomCode(p.config)
	for p.counter < p.config.Qty {
		code := randomCode.GetCode()

		if p.find(code) == nil {
			p.insert(code)
		}
	}
}

func (p *BinaryTree) find(code string) *node {
	if p.rootNode != nil {
		return p.rootNode.find(code)
	}

	return nil
}

func (p *BinaryTree) insert(code string) {
	p.planFile.C <- code
	p.counter++

	if p.rootNode == nil {
		p.rootNode = &node{value: code}
		return
	}

	p.rootNode.insert(code)
}
