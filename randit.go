package main

import (
    "net/http"
    "io/ioutil"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
    //w.Header().Set("Content-Type","text/html; charset=utf-8")
    //io.WriteString(w, readFile("/root/rand/index.html"))
    http.ServeFile(w, r, r.URL.Path[1:])
}

func readFile(filename string) string {
    content, _ := ioutil.ReadFile(filename) //handle err
    return string(content)
}

func main() {
    http.HandleFunc("/", serveHome);
    http.ListenAndServe(":8080", nil)    
}
