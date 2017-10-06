package mcpDiff

import (
	"github.com/modmuss50/MCP-Diff/utils"
	"github.com/modmuss50/MCP-Diff/mcp"
	"strings"
	"errors"
	"github.com/patrickmn/go-cache"
	"time"
	"fmt"
	"strconv"
)

var lookupCache = cache.New(10*time.Minute, 10*time.Minute)

func GetMCPDiff(oldMCP string, newMCP string) (string, string,  error) {
	if strings.Contains(oldMCP, "stable-"){
		oldMCP = "mcp_" + oldMCP
	} else {
		oldMCP = "mcp_snapshot-" + oldMCP
	}

	if strings.Contains(newMCP, "stable-"){
		newMCP = "mcp_" + newMCP
	} else {
		newMCP = "mcp_snapshot-" + newMCP
	}

	dataDir := "data"

	if(utils.FileExists(dataDir)){
		utils.DeleteDir(dataDir)
	}

	utils.MakeDir(dataDir)

	oldMCPFile := dataDir + "/" + oldMCP + ".zip"
	newMCPFile := dataDir + "/" + newMCP + ".zip"

	if !utils.FileExists(oldMCPFile){
		urlSub := strings.Replace(oldMCP, "-", "/", 1) + "/"
		if ! utils.DownloadURL("http://export.mcpbot.bspk.rs/" + urlSub + oldMCP + ".zip", oldMCPFile) {
			return "","", errors.New("Failed to download old MCP export")
		}

	}

	if !utils.FileExists(newMCPFile){
		urlSub := strings.Replace(newMCP, "-", "//", 1) + "/"
		if ! utils.DownloadURL("http://export.mcpbot.bspk.rs/" + urlSub + newMCP + ".zip", newMCPFile){
			return "","",  errors.New("Failed to download new MCP export")
		}

	}
	oldMCPDir := dataDir + "/" + oldMCP
	newMCPDir := dataDir + "/" + newMCP

	utils.ExtractZip(oldMCPFile, oldMCPDir)
	utils.ExtractZip(newMCPFile, newMCPDir)

	if !utils.FileExists(oldMCPDir){
		return  "","",  errors.New("Failed to extract old MCP export, is that a correct mcp name?")
	}

	if !utils.FileExists(newMCPDir){
		return "","",  errors.New("Failed to extract old MCP export, is that a correct mcp name?")
	}


	oldFields := mcp.LoadFields(utils.ReadLinesFromFile(oldMCPDir + "/" + "fields.csv"))
	newFields := mcp.LoadFields(utils.ReadLinesFromFile(newMCPDir + "/" + "fields.csv"))

	response := ""

	changed := 0
	added := 0
	lost := 0

	for _, field := range newFields {
		if value, ok := oldFields[field.Searge]; ok {
			if value.Name != field.Name{
				response += "Changed Field: " + field.Name + " -> " + value.Name  + "\n"
				changed++
			}
		} else {
			response += "Added Field: " + field.Name + " srg: " + field.Searge + "\n"
			added++
		}
	}
	for _, field := range oldFields {
		if _, ok := newFields[field.Searge]; ok {
			//nope
		} else {
			response += "Lost Field: " + field.Name + " srg: " + field.Searge + "\n"
			lost++
		}
	}

	oldMethods := mcp.LoadMethods(utils.ReadLinesFromFile(oldMCPDir + "/" + "methods.csv"))
	newMethods := mcp.LoadMethods(utils.ReadLinesFromFile(newMCPDir + "/" + "methods.csv"))

	for _, method := range newMethods {
		if value, ok := oldMethods[method.Searge]; ok {
			if value.Name != method.Name{
				response += "Changed Method: " + method.Name + " -> " + value.Name + "\n"
				changed++
			}
		} else {
			response += "Added Method: " + method.Name + " srg: " + method.Searge + "\n"
			added++
		}
	}
	for _, method := range oldMethods {
		if _, ok := newMethods[method.Searge]; ok {
			//nope
		} else {
			response += "Lost Method: " + method.Name + " srg: " + method.Searge + "\n"
			lost++
		}
	}
	return response, "Added: " + strconv.Itoa(added) + " Changed: " + strconv.Itoa(changed) + " Lost: " + strconv.Itoa(lost),  nil
}

func LookupMethod(input string) string {
	methodURL := "http://export.mcpbot.bspk.rs/methods.csv"
	methodCSV := "data/methods.csv"
	utils.DownloadURL(methodURL, methodCSV)

	var methods map[string]mcp.Method

	mcache, found := lookupCache.Get("methods")
	if found {
		methods = mcache.(map[string]mcp.Method)
		fmt.Println("Using cache")
	} else {
		methods = mcp.LoadMethods(utils.ReadLinesFromFile(methodCSV))
		lookupCache.Set("methods", methods, cache.DefaultExpiration)
		fmt.Println("Adding to cache")
	}

	for _,method := range methods {
		if(input == method.Name || input == method.Searge){
			return "Method: SRG=`" + method.Searge + "`	Name=`" + method.Name +"` 	Description=`" + method.Desc + "`"
		}
	}
	return "Failed to find method"
}

func LookupField(input string) string {
	fieldURL := "http://export.mcpbot.bspk.rs/fields.csv"
	fieldCSV := "data/fields.csv"
	utils.DownloadURL(fieldURL, fieldCSV)

	var fields map[string]mcp.Field

	mcache, found := lookupCache.Get("fields")
	if found {
		fields = mcache.(map[string]mcp.Field)
		fmt.Println("Using cache")
	} else {
		fields = mcp.LoadFields(utils.ReadLinesFromFile(fieldCSV))
		lookupCache.Set("fields", fields, cache.DefaultExpiration)
		fmt.Println("Adding to cache")
	}

	for _,field := range fields {
		if(input == field.Name || input == field.Searge){
			return "Field: SRG=`" + field.Searge + "`	Name=`" + field.Name +"` 	Description=`" + field.Desc + "`"
		}
	}
	return "Failed to find Field"
}
