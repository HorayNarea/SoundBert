package main

type sound struct {
	name string
	path string
}

type config struct {
	Host           string
	Port           int
	Sounds         string
	AllowedFormats []string
}
