package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strings"

	"github.com/HorayNarea/go-mplayer"
)

func list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snippetlist)
}

func stop(w http.ResponseWriter, r *http.Request) {
	mplayer.SendCommand("stop")
}

func play(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	checkErr(err)

	filename := r.PostFormValue("filename")
	content := map[string]interface{}{}

	for _, pre := range []string{".", "..", "/"} {
		if strings.HasPrefix(filename, pre) {
			http.Error(w, "Path not allowed", http.StatusBadRequest)
			return
		}
	}

	mplayer.SendCommand("loadfile " + `"` + path.Join(conf.Sounds, filename) + `"`)

	content["playing"] = filename
	json.NewEncoder(w).Encode(content)
}