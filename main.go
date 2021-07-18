package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

const dir = "files"

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	savePath := filepath.Join(dir, handler.Filename)
	err = ioutil.WriteFile(savePath, fileBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	addr := "0.0.0.0:8080"
	fmt.Printf("listening to %s", addr)
	http.HandleFunc("/upload", uploadFile)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(addr, nil)
}

func main() {
	setupRoutes()
}
