package main

import (
	"fmt"
	"github.com/modmuss50/MCP-Diff/mcpDiff"
)

func main(){
	fmt.Print(mcpDiff.GetMCPDiff("stable-29-1.10.2", "stable-32-1.11"))
}

