package main

import (
	"net/http"
	"path"
	"strings"

	"github.com/HorayNarea/go-mplayer"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	content, err := staticBox.Bytes("index.html")
	checkErr(err)
	ctx.Data(http.StatusOK, "text/html", content)
}

func help(ctx *gin.Context) {
	content, err := staticBox.Bytes("help.html")
	checkErr(err)
	ctx.Data(http.StatusOK, "text/html", content)
}

func favicon(ctx *gin.Context) {
	content, err := staticBox.Bytes("favicon.ico")
	checkErr(err)
	ctx.Data(http.StatusOK, "image/x-icon", content)
}

func logo(ctx *gin.Context) {
	content, err := staticBox.Bytes("logo.png")
	checkErr(err)
	ctx.Data(http.StatusOK, "image/png", content)
}

func list(ctx *gin.Context) {
	json := gin.H{}
	for _, s := range snippets {
		json[s.name] = s.path
	}
	ctx.JSON(http.StatusOK, json)
}

func stop(ctx *gin.Context) {
	mplayer.SendCommand("stop")
}

func play(ctx *gin.Context) {
	filename := ctx.PostForm("filename")

	for _, pre := range []string{".", "..", "/"} {
		if strings.HasPrefix(filename, pre) {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}
	}

	mplayer.SendCommand("loadfile " + `"` + path.Join(conf.Sounds, filename) + `"`)
	ctx.JSON(http.StatusOK, gin.H{"playing": filename})
}
