package main

import (

	"net/http"
	"log"
	"time"
	"encoding/json"
	"text/template"
	"io/ioutil"
	"fmt"
	"strings"
)

type ParkeringsOmraade struct {
	Breddegrad 		float64 		`json:"breddegrad"`
	Lengdegrad 		float64 		`json:"lengdegrad"`
	AktivVersjon 	AktivVersjon 	`json:"aktivVersjon"`
}

type AktivVersjon struct {
	Navn 							string 	`json:"navn"`
	Adresse 						string 	`json:"adresse"`
	Postnummer						string 	`json:"postnummer"`
	Poststed 						string 	`json:"poststed"`
	AntallLadeplasser 				int 	`json:"antallLadeplasser"`
	AntallAvgiftsbelagtePlasser 	int 	`json:"antallAvgiftsbelagtePlasser"`
	AntallAvgiftsfriePlasser		int 	`json:"antallAvgiftsfriePlasser"`
	AntallForflytnigshemmede 		int 	`json:"antallForflytningshemmede"`
	VurderingForflytningshemmede 	string 	`json:"vurderingForflytningshemmede"`
	TypeParkeringsomrade 			string 	`json:"typeParkeringsomrade"`
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
	parkeringer := make([]ParkeringsOmraade, 0)
	getJSON("https://www.vegvesen.no/ws/no/vegvesen/veg/parkeringsomraade/parkeringsregisteret/v1/parkeringsomraade?datafelter=alle", &parkeringer)
	var parkeringsSøk []ParkeringsOmraade
	antallTreff := 0
	search := r.URL.Query().Get("search")

	for _, parkering := range parkeringer {
		lowerParkeringPoststed := strings.ToLower(parkering.AktivVersjon.Poststed)
		lowerSøk := strings.ToLower(search)
		lowerSøkParkeringNavn := strings.ToLower(parkering.AktivVersjon.Navn)

		if strings.Contains(lowerParkeringPoststed, lowerSøk) || strings.Contains(lowerSøkParkeringNavn, lowerSøk) {
			parkeringsSøk = append(parkeringsSøk, parkering)
			antallTreff++

			if antallTreff > 50 {
				break
			}
		}
	}


	t, err := template.ParseFiles("Oblig4\\parkeringsomraade.html")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	t.Execute(w, parkeringsSøk)
}

func main() {
	http.HandleFunc("/", parkeringsomraade)
	log.Fatal(http.ListenAndServe(":8080", nil))
}