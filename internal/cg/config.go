package cg

import (
	"encoding/json"
	"os"
	"runtime"
)

const (
	configFileName = "config.json"
)

type Config struct {
	Qty        int
	InLine     bool
	EOF        string
	Threads    int
	BufferSize int

	Code Code
}

type Code struct {
	Length                      int
	Set                         []string
	Restrictions                bool
	RepetitionCharacterInString int
	RepetitionSameSymbol        int
}

func NewConfig() *Config {
	c := &Config{}

	c.Qty = 1000
	c.InLine = false
	c.EOF = "\r\n"
	c.Threads = runtime.NumCPU() - 1
	if c.Threads < 1 {
		c.Threads = 1
	}
	c.BufferSize = 100

	c.Code.Length = 20
	c.Code.Set = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	c.Code.Restrictions = true
	c.Code.RepetitionCharacterInString = 3
	c.Code.RepetitionSameSymbol = 2

	c.load()

	// Для маленьких значений Qty выходной файл получается пустым
	if c.Qty < 100 {
		c.Qty = 100
	}

	return c
}

func (p *Config) load() {
	_, err := os.Stat(configFileName)
	if !os.IsNotExist(err) {
		byteValue, err := os.ReadFile(configFileName)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(byteValue, p)
		if err != nil {
			p.removeCorrupted()
			panic(err)
		}
	}
}

func (p *Config) Save() {
	jsonFile, err := os.Create(configFileName)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.WriteString(p.String())
	if err != nil {
		panic(err)
	}
}

func (p *Config) removeCorrupted() {
	_, err := os.Stat(configFileName)
	if !os.IsNotExist(err) {
		err := os.Remove(configFileName)
		if err != nil {
			err = os.Remove(configFileName)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (p *Config) String() string {
	byteValue, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(byteValue)
}
