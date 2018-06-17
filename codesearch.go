package main

import (
	"codesearch/indexer"
	"codesearch/util"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	templateDir = "templates/"
	uploadHTML  = templateDir + "upload.html"
	searchHTML  = templateDir + "search.html"
	resultHTML  = templateDir + "result.html"
)

// UploadHandler Handles upload requests
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(uploadHTML)
	util.CheckError(err)
	t.Execute(w, nil)
}

// SearchHandler Handles search requests
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(searchHTML)
	util.CheckError(err)
	t.Execute(w, nil)
}

// ResultHandler Handles result display requests
func ResultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pat := strings.Join(r.Form["pattern"], "")
	repoName := strings.Join(r.Form["repo"], "")
	result := indexer.QueryIndex(pat, repoName)

	t, err := template.ParseFiles(resultHTML)
	util.CheckError(err)
	t.Execute(w, result)
}

// RepoIndexHandler Handles repo upload requests
func RepoIndexHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	repoURL := strings.Join(r.Form["repoURL"], "")
	indexer.IndexRepo(repoURL)

	files, err := ioutil.ReadDir(indexer.IndexDir)
	util.CheckError(err)

	t, err := template.ParseFiles(searchHTML)
	util.CheckError(err)
	t.Execute(w, files)
}

func main() {
	http.HandleFunc("/", UploadHandler)
	http.HandleFunc("/search/", SearchHandler)
	http.HandleFunc("/result/", ResultHandler)
	http.HandleFunc("/repoindex/", RepoIndexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
