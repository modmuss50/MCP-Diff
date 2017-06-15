package mcp

import (
	"strings"
)

type Method struct {
	Searge string
	Name string
	Desc string
}

func LoadMethods(lines []string) map[string]Method {
	m := make(map[string]Method)
	for _,line := range lines {
		if strings.HasPrefix(line, "func_") {
			parts := strings.Split(line, ",")
			m[parts[0]] = Method{Searge:parts[0], Name:parts[1], Desc:parts[3]}
		}
	}
	return m
}