package main

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func isValidOption(param string, options []string) bool {
	for option_i := range options {
		if param == options[option_i] {
			return true
		}
	}
	return false
}

func param_parse(params []string, options []string) (map[string]string, error) {
	result := make(map[string]string)

	skip := false

	for param_i := range params {

		if skip {
			skip = false
			continue
		}

		param := params[param_i]

		if !isValidOption(param, options) {
			return nil, errors.New("Invalid option '" + param + "'")
		}

		var value string

		if param_i+1 != len(params) {
			if !isValidOption(params[param_i+1], options) {
				value = params[param_i+1]
				skip = true
			}
		}

		result[param] = value
	}

	return result, nil
}

func getJob(args []string, options []string, job *job) error {
	parsed_options, err := param_parse(args, options)

	if _, found := parsed_options["--help"]; found {
		job.help = true
		return nil
	}

	if _, found := parsed_options["-u"]; found {
		u, err := url.Parse(parsed_options["-u"])

		if err != nil {
			fmt.Println("Invald URL: " + err.Error())
			return err
		}

		folders := strings.Split(u.Path, "/")

		if len(folders) != 0 {
			job.problem_name = folders[len(folders)-1]
		}
	}

	if _, found := parsed_options["-n"]; found {
		if len(parsed_options["-n"]) != 0 {
			job.problem_name = parsed_options["-n"]
		}
	}

	if _, found := parsed_options["-b"]; found {
		base, err := url.Parse(parsed_options["-b"])

		if err != nil {
			return errors.New("Invalud URL: " + err.Error())
		}

		if base.String() == "" {
			return errors.New("Please specify a base")
		}
	}

	if _, found := parsed_options["--no-cache"]; found {
		job.cache = false
	}

	if _, found := parsed_options["-e"]; found {
		job.execute = parsed_options["-e"]
	}

	if _, found := parsed_options["-r"]; found {
		job.rescan = true
	}

	if _, found := parsed_options["-t"]; found {
		job.threads, err = strconv.Atoi(parsed_options["-t"])

		if err != nil {
			return errors.New("Invalid number of threads")
		}
	}

	return nil
}
