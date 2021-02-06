package main

import (
	"fmt"
)

func help() {
	//options := []string{"--help", "-n", "-u", "-t", "-e", "-b", "--no-cache"}

	fmt.Println("CowderTest Help")
	fmt.Println("=============================================================")
	fmt.Println("About: Tests solutions for Kattis Kattis Problems")
	fmt.Println("")
	fmt.Println("Options              Description")
	fmt.Println("-------------------------------------------------------------")
	fmt.Println(" --help              Displays this help")
	fmt.Println("  -u <url>           The URL of the problem to test")
	fmt.Println("  -n <name>          The name of the problem to test")
	fmt.Println("                     (Will override -u)")
	fmt.Println("  -e <command>       What to execute to run the solution")
	fmt.Println("  -t                 The amount of threads to use when ")
	fmt.Println("                     running the tests. Defaults to 1")
	fmt.Println(" --no-cache          Always fetch a new copy of the tests")
	fmt.Println(" -r                  Rescan the tests scored in cache.")
	fmt.Println("  -b                 The base URL for Kattis.")
	fmt.Println("                     Defaults to https://open.kattis.com/")
}
