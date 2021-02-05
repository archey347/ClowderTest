package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {

	options := []string{"--help", "-n", "-u", "-t", "-e", "-b", "--no-cache", "-r"}

	base, _ := url.Parse("https://open.kattis.com")

	mainJob := job{
		problem_name: "",
		base:         *base,
		cache:        true,
		threads:      1,
		execute:      "",
		rescan:       false,
		help:         false,
	}

	err := getJob(os.Args[1:], options, &mainJob)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if !mainJob.help {
		fmt.Println("Problem name: " + mainJob.problem_name)
		fmt.Println("Execute     : " + mainJob.execute)
		fmt.Print("Cache       : ")
		fmt.Println(mainJob.cache)
		fmt.Print("Threads     : ")
		fmt.Println(mainJob.threads)
		fmt.Println("Base        : " + base.String())
	} else {
		help()
	}

	err = downloadTests(*base, mainJob.problem_name, mainJob.cache)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = getTests(mainJob.problem_name, mainJob.rescan)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
