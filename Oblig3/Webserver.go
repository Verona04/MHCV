package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Astros struct {
	Number int `json:"number"`
	People []People `json:"people"`
}

type People struct {
	Name string `json:"name"`
	Craft string `json:"craft"`
}

type Miljo struct {
	Antall int `json:"posts"`
	Stasjoner []Stasjoner `json:"entries"`
}

type Stasjoner struct {
	Navn string `json:"navn"`
	Plast string `json:"plast"`
}

type randomJoke struct {
	Value string `json:"value"`
}

type KommuneNummer struct {
	ContainedItems []ContainedItems `json:"containeditems"`
	ContentSummary string           `json:"contentsummary"`
}

type ContainedItems struct {
	Description string `json:"description"`
	Label string `json:"label"`
	Status string `json:"status"`
	LastUpdated string `json:"lastUpdated"`
}

type TrumpQuote struct {
	Quote string `json:"message"`
}

//type Quotes struct {
//	Quote string `json:"message"`}

func miljoStasjoner(w http.ResponseWriter, r *http.Request){
	foo1:= new(Miljo)
	getJSON("https://hotell.difi.no/api/json/stavanger/miljostasjoner", foo1)

	t, _ := template.ParseFiles("Oblig3\\miljoStasjoner.html")
	t.Execute(w, foo1)
}

func peopleInSpace(w http.ResponseWriter, r *http.Request) {
	foo1 := new(Astros)
	getJSON("http://api.open-notify.org/astros.json", foo1)

	t, _ := template.ParseFiles("Oblig3\\pepoleInSpace.html")
	t.Execute(w, foo1)
}

func chuckNorris(w http.ResponseWriter, r *http.Request) {
	foo1 := new(randomJoke)
	getJSON("https://api.chucknorris.io/jokes/random", foo1)

	t, _ := template.ParseFiles("Oblig3\\randomchucknorrisjoke.html")
	t.Execute(w, foo1)
}

func listeOverKommuneNummer(w http.ResponseWriter, r *http.Request) {
	foo1 := new(KommuneNummer)
	getJSON("https://register.geonorge.no/api/subregister/sosi-kodelister/kartverket/kommunenummer-alle.json?", foo1)
	t, _ := template.ParseFiles("Oblig3\\kommuner.html")
	t.Execute(w, foo1)
}

func trumpThinks(w http.ResponseWriter, r *http.Request) {
	foo1 := new(TrumpQuote)
	getJSON("https://api.whatdoestrumpthink.com/api/v1/quotes/random", foo1)

	t, _ := template.ParseFiles("Oblig3\\randomTrumpQuote.html")
	t.Execute(w, foo1)
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

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello, client"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/1", peopleInSpace)
	http.HandleFunc("/2", miljoStasjoner)
	http.HandleFunc("/3", chuckNorris)
	http.HandleFunc("/4", listeOverKommuneNummer)
	http.HandleFunc("/5", trumpThinks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}