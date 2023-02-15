package main

import (
	_ "embed"
	"fmt"
)

var fileString string


var fileByte []byte



//var folder embed.FS

func main() {

	fmt.Print(fileString)
	fmt.Print(string(fileByte))

	// content1, _ := folder.ReadFile("D:\vivasoftOffice\\GoLand\\EmbedDirective\test.txt")
	// fmt.Print(string(content1))

	// content2, _ := folder.ReadFile("D:\\vivasoftOffice\\GoLand\\EmbedDirective\\test.txt")
	// fmt.Println(string(content2))
}
