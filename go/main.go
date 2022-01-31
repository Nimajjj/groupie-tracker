package main

import (
  "fmt"
  "net/http"
  "html/template"
  "encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type ViewData struct {
  Characters  []Character
}

type Character struct {
  Name      string
  Fullname  string
  Gender    string
}

func loadAPI() []Character {
  var characters []Character

	url := "https://adventuretimeapi.herokuapp.com/people"

	httpClient := http.Client{
		Timeout: time.Second * 2, // define timeout
	}

	//create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "API AT test <3")

	//make api call
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	//parse response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &characters)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

  return characters
}


func main() {
  viewData := ViewData{Characters: loadAPI()}

  fmt.Println("\nStarting server -> localhost:80")

  indexTemplate := template.Must(template.ParseFiles("../page/index.html"))

  cssFolder := http.FileServer(http.Dir("../css"))
  http.Handle("/css/", http.StripPrefix("/css/", cssFolder))

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    indexTemplate.Execute(w, viewData)
  })

  http.ListenAndServe(":80", nil)
}
