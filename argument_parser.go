package main

import (
	"errors"
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
