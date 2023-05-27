package cg

import (
	"math/rand"
	"strings"
	"time"
)

type RandomCode struct {
	config *Config

	rand        *rand.Rand
	codeBuilder *strings.Builder
	sameBuilder *strings.Builder
}

func NewRandomCode(config *Config) *RandomCode {
	return &RandomCode{
		config: config,

		rand:        rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
		codeBuilder: &strings.Builder{},
		sameBuilder: &strings.Builder{},
	}
}

func (p *RandomCode) GetCode() string {
	p.codeBuilder.Reset()
	char := ""

	for p.codeBuilder.Len() < p.config.Code.Length {
		char = p.config.Code.Set[p.rand.Intn(len(p.config.Code.Set))]

		if p.config.Code.Restrictions {
			if strings.Count(p.codeBuilder.String(), char) >= p.config.Code.RepetitionCharacterInString {
				continue
			}
			if strings.HasSuffix(p.codeBuilder.String(), p.getSameSymbols(char)) {
				continue
			}
		}

		p.codeBuilder.WriteString(char)
	}

	return p.codeBuilder.String()
}

func (p *RandomCode) getSameSymbols(char string) string {
	p.sameBuilder.Reset()
	for p.sameBuilder.Len() < p.config.Code.RepetitionSameSymbol {
		p.sameBuilder.WriteString(char)
	}
	return p.sameBuilder.String()
}
