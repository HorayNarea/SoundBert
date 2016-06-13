package main

import (
	"flag"
	"net"
	"net/http"
	"path"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/GeertJohan/go.rice"
	"github.com/HorayNarea/go-mplayer"
)

var (
	snippets  []Sound
	conf      Config
	staticBox *rice.Box
	conffile  = flag.String("c", "config.toml", "Configuration file, must be valid TOML")
)

func main() {
	flag.Parse()

	_, err := toml.DecodeFile(*conffile, &conf)
	checkErr(err)

	staticBox = rice.MustFindBox("static")
	conf.Sounds = path.Clean(conf.Sounds)

	mplayer.StartSlave()

	http.Handle("/", http.FileServer(staticBox.HTTPBox()))
	http.HandleFunc("/list", list)
	http.HandleFunc("/stop", stop)
	http.HandleFunc("/play", play)

	http.ListenAndServe(net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port)), nil)
}
