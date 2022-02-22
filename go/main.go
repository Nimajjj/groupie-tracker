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
	Result []struct {
		Nom      string `json:"Nom"`
		Prenom   string `json:"Prenom"`
		Email    string `json:"Email"`
		Github   string `json:"Github,omitempty"`
		Linkedin string `json:"Linkedin,omitempty"`
	} `json:"result"`
}

func loadAPI() ViewData {
  vd := ViewData{}

	url := "https://raw.githubusercontent.com/Nimajjj/groupie-tracker/main/API/etudiant.json"

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

	jsonErr := json.Unmarshal(body, &vd)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

  return vd
}


func main() {
  viewData := loadAPI()

  fmt.Println("\nStarting server -> localhost:80")

  indexTemplate := template.Must(template.ParseFiles("../page/index.html"))

  cssFolder := http.FileServer(http.Dir("../css"))
  http.Handle("/css/", http.StripPrefix("/css/", cssFolder))

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    indexTemplate.Execute(w, viewData)
  })

  http.ListenAndServe(":80", nil)
}
