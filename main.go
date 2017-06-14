package main

import (
	"fmt"
	"github.com/modmuss50/MCP-Diff/mcpDiff"
)

func main(){
	fmt.Print(mcpDiff.GetMCPDiff("20170613-1.12", "20170614-1.12"))
}

