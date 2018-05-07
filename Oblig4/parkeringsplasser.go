package main

import (

	"net/http"
	"log"
	"time"
	"encoding/json"
	"text/template"

	"io/ioutil"
	"fmt"
)

func main() {
	http.HandleFunc("/", parkeringsomraade)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//func parking(w http.ResponseWriter, r *http.Request) {
//	message := "Parkeringsplasser"
//	w.Write([]byte(message))
//}

func getJSON(url string, target interface{}) error {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	json.Unmarshal([]byte(body), &target)
	return nil
//	return json.NewDecoder(r.Body).Decode(&target)
}

func parkeringsomraade (w http.ResponseWriter, r *http.Request) {
	foo1 := make([]ParkeringsOmraade, 0)
	getJSON("https://www.vegvesen.no/ws/no/vegvesen/veg/parkeringsomraade/parkeringsregisteret/v1/parkeringsomraade?datafelter=kart", &foo1)
	t, err := template.ParseFiles("Oblig4\\parkeringsomraade.html")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	t.Execute(w, foo1)
}

type ParkeringsOmraade struct {
	Navn string `json:"navn"`
	Adresse string `json:"adresse"`
	Postnummer string `json:"postnummer"`
	Poststed string `json:"poststed"`
	Breddegrad float64 `json:"breddegrad"`
	Lengdegrad float64 `json:"lengdegrad"`
}

/*
{
	"id":8192,
	"parkeringstilbyderNavn":"TROMSØ OFFENTLIGE PARKERING AS",
	"breddegrad":69.667223,
	"lengdegrad":19.019822,
	"deaktivert":null,
	"versjonsnummer":1,
	"navn":"Tomasjordnes",
	"adresse":"Tomasjordnes 37",
	"postnummer":"9024",
	"poststed":"TOMASJORD",g i henhold til parkeringsforskriften § 62",
	"aktiveringstidspunkt":"2018-01-08T11:01:34Z"
}
 */

