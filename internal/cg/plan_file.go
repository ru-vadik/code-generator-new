package cg

import (
	"bufio"
	"os"
	"strconv"
	"time"
)

type PlanFile struct {
	config *Config

	FileName string
	file     *os.File
	writer   *bufio.Writer

	C chan string
}

func NewPlanFile(config *Config) *PlanFile {
	planFile := &PlanFile{
		config:   config,
		FileName: "codes_" + strconv.FormatInt(time.Now().Unix(), 10) + ".txt",
		C:        make(chan string, config.BufferSize),
	}
	planFile.open()

	go planFile.write()
	return planFile
}

func (p *PlanFile) open() {
	var err error

	p.file, err = os.OpenFile(p.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	p.writer = bufio.NewWriter(p.file)
}

func (p *PlanFile) close() {
	p.writer.Flush()
	p.file.Close()
}

func (p *PlanFile) write() {
	for code := range p.C {
		p.writer.WriteString(code)
		p.writer.WriteString(p.config.EOF)
		if p.config.InLine {
			p.writer.WriteString(p.config.EOF)
		}
	}
	p.close()
}
