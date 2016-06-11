package main

import (
	"log"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sanitizeName(filename string) string {
	for _, ext := range conf.AllowedFormats {
		filename = strings.TrimSuffix(filename, "."+ext)
	}

	for _, s := range []string{"_", "-", "."} {
		filename = strings.Replace(filename, s, " ", -1)
	}

	filename = strings.Title(filename)

	return filename
}

func addSound(p string, info os.FileInfo, err error) error {
	checkErr(err)
	if info.IsDir() == false {
		for _, ext := range conf.AllowedFormats {
			if strings.HasSuffix(p, "."+ext) {
				name := strings.TrimPrefix(p, conf.Sounds+"/")
				snippets = append(snippets, Sound{sanitizeName(name), name})
			}
		}
	}
	return nil
}
