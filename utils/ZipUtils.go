package utils

import (
	"github.com/mholt/archiver"
	"fmt"
)

func ExtractZip(zip string, dest string){
	fmt.Println("Extracting " + zip)
	archiver.DefaultZip.Unarchive(zip, dest)
}