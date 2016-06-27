package main

import (
	"encoding/json"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/HorayNarea/go-mplayer"
)

func reloadSounds(w http.ResponseWriter, r *http.Request) {
	snippets = nil
	snippetlist = map[string]string{}

	filepath.Walk(conf.Sounds, addSound)
	for _, s := range snippets {
		snippetlist[s.name] = s.path
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snippetlist)
}

func stop(w http.ResponseWriter, r *http.Request) {
	mplayer.SendCommand("stop")
}

func playPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	checkErr(err)

	filename := r.PostFormValue("filename")
	content := map[string]string{}

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

func playGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	filename := strings.TrimPrefix(r.URL.Path, "/play/")
	content := map[string]string{}

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
