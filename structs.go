package main

type Sound struct {
	name string
	path string
}

type Config struct {
	Host           string
	Port           int
	Sounds         string
	AllowedFormats []string
}
