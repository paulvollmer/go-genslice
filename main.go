package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/paulvollmer/go-genslice/generator"
)

func usage() {
	fmt.Print("Usage: go-genslice [flags]\n\n")
	fmt.Print("Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flagDatasource := flag.String("d", "", "path to a line separated input source")
	flagPackageName := flag.String("p", "main", "the package name")
	flagVarName := flag.String("v", "Data", "the variable name")
	flagVarType := flag.String("t", "string", "the variable type can be 'string' or 'byte'")
	flagSort := flag.Bool("s", false, "sort the input data")
	flagPrintDefaultTemplate := flag.Bool("print-default-template", false, "print the default template to stdout and exit")
	flag.Usage = usage
	flag.Parse()

	if *flagPrintDefaultTemplate {
		fmt.Println(generator.DefaultTemplate)
		os.Exit(0)
	}

	if *flagDatasource == "" {
		fmt.Println("missing -d flag")
		os.Exit(2)
	}

	if *flagVarType != "string" && *flagVarType != "byte" {
		fmt.Println("unknown type. use 'string' or 'byte'")
		os.Exit(2)
	}

	dataBytes, err := ioutil.ReadFile(*flagDatasource)
	if err != nil {
		fmt.Printf("could not read file %q", *flagDatasource)
		os.Exit(2)
	}

	tpl := generator.DefaultTemplate
	config := generator.Config{
		PackageName: *flagPackageName,
		VarName:     *flagVarName,
		VarType:     *flagVarType,
		Sort:        *flagSort,
		Data:        strings.Split(string(dataBytes), "\n"),
	}
	code, err := generator.GenerateCode(tpl, config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(code))
}
