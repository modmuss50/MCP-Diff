package main

import (
	"fmt"
	"github.com/modmuss50/MCP-Diff/mcpDiff"
)

func main(){
	fmt.Print(mcpDiff.GetMCPDiff("stable-29-1.10.2", "stable-32-1.11"))

	fmt.Print(mcpDiff.LookupMethod("func_74780_a"))

	val, _ := mcpDiff.GetLatestSnapshot("1.12")
	fmt.Print(val)

	fmt.Println(mcpDiff.ParseMCPString("1.12"))
	fmt.Println(mcpDiff.ParseMCPString("20140925"))
}

