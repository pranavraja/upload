package main

import (
	"github.com/codegangsta/martini"
	"io"
	"net/http"
	"os"
)

func main() {
	m := martini.Classic()
	m.Post("/upload/:file", func(params martini.Params, w http.ResponseWriter, r *http.Request) {
		f, err := os.Create(params["file"])
		if err != nil {
			println(err.Error())
			http.Error(w, "Couldn't open file", 500)
			return
		}
		_, err = io.Copy(f, r.Body)
		if err != nil {
			println(err.Error())
			http.Error(w, "Couldn't save file", 500)
			return
		}
		w.WriteHeader(200)
	})
	m.Run()
}
