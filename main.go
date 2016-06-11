package main

import (
	"flag"
	"path"
	"path/filepath"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/GeertJohan/go.rice"
	"github.com/HorayNarea/go-mplayer"
	"github.com/gin-gonic/gin"
)

var snippets []Sound
var conf Config
var staticBox *rice.Box
var conffile string
var debug bool

func init() {
	flag.StringVar(&conffile, "c", "config.toml", "Configuration file, must be valid TOML (shorthand)")
	flag.BoolVar(&debug, "d", false, "Start the webserver in debug-mode (shorthand)")
	flag.Parse()

	_, err := toml.DecodeFile(conffile, &conf)
	checkErr(err)
}

func main() {
	staticBox = rice.MustFindBox("static")
	conf.Sounds = path.Clean(conf.Sounds)
	filepath.Walk(conf.Sounds, addSound)

	mplayer.StartSlave()

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/", index)
	router.GET("/help.html", help)
	router.GET("/favicon.ico", favicon)
	router.GET("/logo.png", logo)
	router.StaticFS("/assets", staticBox.HTTPBox())

	router.GET("/list", list)
	router.GET("/stop", stop)
	router.POST("/play", play)

	router.Run(conf.IP + ":" + strconv.Itoa(conf.Port))
}
