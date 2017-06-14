package mcpDiff

import (
	"github.com/modmuss50/MCP-Diff/utils"
	"github.com/modmuss50/MCP-Diff/mcp"
	"strings"
)


func GetMCPDiff(oldMCP string, newMCP string) string {
	oldMCP = "mcp_snapshot-" + oldMCP
	newMCP = "mcp_snapshot-" + newMCP

	dataDir := "data"
	if !utils.FileExists(dataDir){
		utils.MakeDir(dataDir)
	}

	oldMCPFile := dataDir + "/" + oldMCP + ".zip"
	newMCPFile := dataDir + "/" + newMCP + ".zip"

	if !utils.FileExists(oldMCPFile){
		urlSub := strings.Replace(oldMCP, "-", "/", 1) + "/"
		if ! utils.DownloadURL("http://export.mcpbot.bspk.rs/" + urlSub + oldMCP + ".zip", oldMCPFile) {
			return "Failed to download old MCP export"
		}

	}

	if !utils.FileExists(newMCPFile){
		urlSub := strings.Replace(newMCP, "-", "//", 1) + "/"
		if ! utils.DownloadURL("http://export.mcpbot.bspk.rs/" + urlSub + newMCP + ".zip", newMCPFile){
			return "Failed to download new MCP export"
		}

	}

	utils.ExtractZip(oldMCPFile, dataDir + "/" + oldMCP)
	utils.ExtractZip(newMCPFile,dataDir + "/" + newMCP)

	oldMCPDir := dataDir + "/" + oldMCP
	newMCPDir := dataDir + "/" + newMCP

	oldFields := mcp.LoadFields(utils.ReadLinesFromFile(oldMCPDir + "/" + "fields.csv"))
	newFields := mcp.LoadFields(utils.ReadLinesFromFile(newMCPDir + "/" + "fields.csv"))

	response := ""

	for _, field := range newFields {
		if value, ok := oldFields[field.Searge]; ok {
			if value.Name != field.Name{
				response += "Changed Field: " + value.Name + " from " + field.Name + "\n"
			}
		} else {
			response += "Added Field: " + field.Name + " srg: " + field.Searge + "\n"
		}
	}

	oldMethods := mcp.LoadMethods(utils.ReadLinesFromFile(oldMCPDir + "/" + "methods.csv"))
	newMethods := mcp.LoadMethods(utils.ReadLinesFromFile(newMCPDir + "/" + "methods.csv"))

	for _, method := range newMethods {
		if value, ok := oldMethods[method.Searge]; ok {
			if value.Name != method.Name{
				response += "Changed Method: " + value.Name + " from " + method.Name + "\n"
			}
		} else {
			response += "Added Method: " + method.Name + " srg: " + method.Searge + "\n"
		}
	}
	return response
}
