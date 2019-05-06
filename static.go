package main

import (
	"net/http"
	"os"
)

func staticHandler(notFound http.HandlerFunc) http.HandlerFunc {

	fileServer := http.FileServer(http.Dir("static"))
	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/static/" {
			notFound.ServeHTTP(w, r)
			return
		}

		if file, err := os.Stat("." + r.URL.Path); err != nil {
			notFound.ServeHTTP(w, r)
			return
		} else {

			if file.Mode().IsDir() {
				notFound.ServeHTTP(w, r)
				return
			}

			http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)

		}

	}

}
