package mcp

import (
	"strings"
)

type Field struct {
	Searge string
	Name string
}

func LoadFields(lines []string) map[string]Field {
	m := make(map[string]Field)
	for _,line := range lines {
		if strings.HasPrefix(line, "field_") {
			parts := strings.Split(line, ",")
			m[parts[0]] = Field{Searge:parts[0], Name:parts[1]}
		}
	}
	return m
}