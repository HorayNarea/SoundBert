package main

import (
	"flag"
	"net"
	"net/http"
	"path"
	"path/filepath"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/GeertJohan/go.rice"
	"github.com/HorayNarea/go-mplayer"
)

var (
	snippets    []Sound
	conf        Config
	staticBox   *rice.HTTPBox
	conffile    = flag.String("c", "config.toml", "Configuration file, must be valid TOML")
	snippetlist = map[string]interface{}{}
)

func main() {
	flag.Parse()

	_, err := toml.DecodeFile(*conffile, &conf)
	checkErr(err)

	staticBox = rice.MustFindBox("static").HTTPBox()
	conf.Sounds = path.Clean(conf.Sounds)

	filepath.Walk(conf.Sounds, addSound)
	for _, s := range snippets {
		snippetlist[s.name] = s.path
	}

	mplayer.StartSlave()

	http.Handle("/", http.FileServer(staticBox))
	http.HandleFunc("/list", list)
	http.HandleFunc("/reload_sounds", reload_sounds)
	http.HandleFunc("/stop", stop)
	http.HandleFunc("/play", play)

	http.ListenAndServe(net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port)), nil)
}
