package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {

	options := []string{"--help", "-n", "-u", "-t", "-e", "-b", "--no-cache"}

	parsed_options, err := param_parse(os.Args[1:], options)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var problem_name string = ""
	base, _ := url.Parse("https://open.kattis.com")
	var cache bool = true
	var threads int = 1
	var execute string

	if _, found := parsed_options["--help"]; found {
		help()
		return
	}

	if _, found := parsed_options["-u"]; found {
		u, err := url.Parse(parsed_options["-u"])

		if err != nil {
			fmt.Println("Invald URL: " + err.Error())
		}

		folders := strings.Split(u.Path, "/")

		if len(folders) != 0 {
			problem_name = folders[len(folders)-1]
		}
	}

	if _, found := parsed_options["-n"]; found {
		if len(parsed_options["-n"]) != 0 {
			problem_name = parsed_options["-n"]
		}
	}

	if _, found := parsed_options["-b"]; found {
		base, err := url.Parse(parsed_options["-b"])

		if err != nil {
			fmt.Println("Invald URL: " + err.Error())
		}

		if base.String() == "" {
			fmt.Println("Please specify a base")
			return
		}
	}

	if _, found := parsed_options["--no-cache"]; found {
		cache = false
	}

	if _, found := parsed_options["-e"]; found {
		execute = parsed_options["-e"]
	}

	if _, found := parsed_options["-t"]; found {
		threads, err = strconv.Atoi(parsed_options["-t"])

		if err != nil {
			fmt.Println("Invalid number")
			return
		}
	}

	fmt.Println("Problem name: " + problem_name)
	fmt.Println("Execute     : " + execute)
	fmt.Print("Cache       : ")
	fmt.Println(cache)
	fmt.Print("Threads     : ")
	fmt.Println(threads)
	fmt.Println("Base        : " + base.String())

	_, err = getTests(*base, problem_name, cache)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Download, unzip and analyse tests

}
