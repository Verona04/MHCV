package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Astros struct {
	Number int `json: number`
	People []People `json: people`
}

type People struct {
	Name string `json: name`
	Craft string `json: craft`
}



func getJSON(url string, target interface{}) error {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func peopleInSpace(w http.ResponseWriter, r *http.Request) {
	foo1 := new(Astros)
	getJSON("http://api.open-notify.org/astros.json", foo1)

	t, _ := template.ParseFiles("Oblig3\\pepoleInSpace.html")
	t.Execute(w, foo1)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello, client"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/1", peopleInSpace)
	http.HandleFunc("/3", chuckNorris)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
