package cg

import (
	"sync"
)

type BinaryTree struct {
	config   *Config
	planFile *PlanFile

	wg       *sync.WaitGroup
	mu       *sync.Mutex
	rootNode *node
	counter  int
}

func NewBinaryTree(config *Config, planFile *PlanFile) *BinaryTree {
	return &BinaryTree{
		config:   config,
		planFile: planFile,

		wg:       &sync.WaitGroup{},
		mu:       &sync.Mutex{},
		rootNode: nil,
		counter:  0,
	}
}

func (p *BinaryTree) Start() {
	defer func() {
		close(p.planFile.C)
	}()

	p.wg.Add(p.config.Threads)
	for i := 0; i < p.config.Threads; i++ {
		go p.work()
	}
	p.wg.Wait()
}

func (p *BinaryTree) work() {
	defer p.wg.Done()

	randomCode := NewRandomCode(p.config)
	p.mu.Lock()
	for p.counter < p.config.Qty {
		code := randomCode.GetCode()

		if p.find(code) == nil {
			p.insert(code)
		}
	}
	p.mu.Unlock()
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
