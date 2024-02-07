package main

import (
	"fmt"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFile(fPath string) string {
	dat, err := os.ReadFile("./public" + fPath)

	check(err)
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
		fmt.Fprintln(w, resp)
	} else if uPath == "/favicon.ico" {
		// do nothing
	} else {
		// match the pathName and check for exact filename
		// remove the initial / optional
		resp := getFile(uPath)
		fmt.Fprintln(w, resp)
	}
}

/*
	request lifecycle -> only serves

	1. check if file with exact file name exist -> then serve the file
	2. if file does not exist -> fileName.html -> if found then serve the file
	3. if not found .html -> return 404 NOT FOUND
*/
