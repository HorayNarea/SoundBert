package main

import (
	"flag"
	"net"
	"path"
	"path/filepath"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/GeertJohan/go.rice"
	"github.com/HorayNarea/go-mplayer"
	"github.com/gin-gonic/gin"
)

var (
	snippets  []Sound
	conf      Config
	staticBox *rice.Box
	conffile  = flag.String("c", "config.toml", "Configuration file, must be valid TOML")
	debug     = flag.Bool("d", false, "Start the webserver in debug-mode")
)

func main() {
	flag.Parse()

	_, err := toml.DecodeFile(*conffile, &conf)
	checkErr(err)

	staticBox = rice.MustFindBox("static")
	conf.Sounds = path.Clean(conf.Sounds)
	filepath.Walk(conf.Sounds, addSound)

	mplayer.StartSlave()

	if !*debug {
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

	router.Run(net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port)))
}
