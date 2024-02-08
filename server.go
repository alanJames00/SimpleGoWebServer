package main

import (
	"fmt"
	"net/http"
	"os"
)

func getFile(fPath string) string {
	dat, err := os.ReadFile("./public" + fPath)
	if err != nil {
		fmt.Println(err)
		return "fError"
	}

	datString := string(dat)
	return datString
}

func main() {

	http.HandleFunc("/", handler1)
	http.ListenAndServe("localhost:4001", nil)
}

func handler1(w http.ResponseWriter, r *http.Request) {

	// find the URL Path
	uPath := r.URL.Path
	fmt.Println(uPath)
	if uPath == "/" {
		// serve index.html
		resp := getFile("/index.html")
		if resp == "fError" {

			// show 404
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "404 NOT FOUND")

		} else {

			fmt.Fprintln(w, resp)

		}
	} else {

		// match the pathName and check for exact filename
		// remove the initial / optional
		resp := getFile(uPath)
		if resp == "fError" {

			// add .html to the filename and try again
			resp2 := getFile(uPath + ".html")

			if resp2 == "fError" {

				// show 404
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintln(w, "404 NOT FOUND")

			} else {

				// serve the file with .html
				fmt.Fprintln(w, resp2)
			}

		} else {

			// serve the file with exact filename
			fmt.Fprintln(w, resp)
		}

	}
}
