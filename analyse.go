package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getTests(name string, rescan bool) ([]test, error) {

	testDir := filepath.FromSlash(".cache/" + name + "/test_info.csv")

	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		return createAndGetMeta(name)
	} else {
		if rescan {
			os.Remove(testDir)
			return createAndGetMeta(name)
		}

		return getMeta(name)
	}
}

// Reads an existing meta file
func getMeta(name string) ([]test, error) {
	return nil, nil
}

// Creates the meta file
func createAndGetMeta(name string) ([]test, error) {

	var tests []test

	path := filepath.FromSlash(".cache/" + name + "/")

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	listing := make(map[string]bool)

	for _, f := range files {
		listing[f.Name()] = false
	}

	testDir := filepath.FromSlash(".cache/" + name + "/test_info.csv")
	f, err := os.Create(testDir)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	for filename, used := range listing {
		if used {
			continue
		}

		lastInd := strings.LastIndex(filename, ".")
		test_name := filename[:lastInd]
		ext := filename[lastInd+1:]

		output := ""

		if ext == "in" {
			if _, found := listing[test_name+".ans"]; found {
				output = test_name + ".ans"
				listing[output] = true
			}
		} else if ext == "ans" {
			continue
		}

		if output == "" {
			output = requestChoice(listing, filename)
		}

		if output == "" {
			continue
		}

		tests = append(tests, test{
			input:  filename,
			output: output,
		})

		_, err := f.WriteString(filename + "," + output + "\n")

		if err != nil {
			return nil, err
		}

	}

	return tests, nil
}

func requestChoice(listing map[string]bool, name string) string {
	fmt.Println()
	fmt.Println()
	fmt.Println("Unable to find an output file for " + name)
	fmt.Println("Please select from the list below: ")
	fmt.Println()

	counter := 1
	ids := make(map[int]string)

	fmt.Println("[0] This is an output file")

	for filename, used := range listing {
		if used {
			continue
		}

		ids[counter] = filename

		fmt.Print("[")
		fmt.Print(counter)
		fmt.Println("] " + filename)

		counter++
	}

	var input int

	for input == 0 {
		fmt.Print("~# ")
		fmt.Scanln(&input)

		if input == 0 {
			return ""
		}

		if input >= counter {
			input = 0
			fmt.Println("Too big of a number")
		}
	}

	return ids[input]
}
