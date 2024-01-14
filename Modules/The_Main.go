package main

import (
	"fmt"
	addingnumber "module/adding_number"
	anothermodule "module/another_module" // je_name_hoy_te/folder_name
)

// Please note that Which data access to another module that must have first latter is capital
// command to create go.mod file :: go mod init module_name


func main() {
	fmt.Println(anothermodule.Testing())
	fmt.Println(addingnumber.Add(1, 2))
}
