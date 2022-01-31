package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func main() {
  fmt.Println("Starting server -> localhost:80")

  indexTemplate := template.Must(template.ParseFiles("../page/index.html"))

  cssFolder := http.FileServer(http.Dir("../css"))
  http.Handle("/css/", http.StripPrefix("/css/", cssFolder))

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    indexTemplate.Execute(w, nil)
  })

  http.ListenAndServe(":80", nil)
}
