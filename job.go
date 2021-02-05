package main

import (
	"net/url"
)

type job struct {
	problem_name string
	base         url.URL
	cache        bool
	rescan       bool
	threads      int
	execute      string
	help         bool
}
