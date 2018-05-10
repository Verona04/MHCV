package main

import (

	"net/http"
	"log"
	"time"
	"encoding/json"

	"io/ioutil"
	"fmt"
	"strings"
	"math"
	"strconv"
)

type ParkeringsOmraade struct {
	Id 				int 		`json:"id"`
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

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
func distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180
	r = 6378100 // Earth radius in METERS
	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)
	return 2 * r * math.Asin(math.Sqrt(h))
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
	raw, err := ioutil.ReadFile("Oblig4\\parkeringsomraade.html")
	if err != nil {
		fmt.Print(w, "OOps")
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(raw)
}

func apiForRadiusSøk(w http.ResponseWriter, r *http.Request) {
	var longitude, latitude, radiusMeters float64
	longitude, _ = strconv.ParseFloat(r.URL.Query().Get("longitude"), 64)
	latitude, _ = strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
	radiusMeters, _ = strconv.ParseFloat(r.URL.Query().Get("radius"), 64)

	parkeringer := make([]ParkeringsOmraade, 0)
	getJSON("https://www.vegvesen.no/ws/no/vegvesen/veg/parkeringsomraade/parkeringsregisteret/v1/parkeringsomraade?datafelter=alle", &parkeringer)

	var parkeringsSøk []ParkeringsOmraade
	antallTreff := 0
	for _, parkering := range parkeringer {
		parkDistance := distance(parkering.Breddegrad, parkering.Lengdegrad, latitude, longitude)
		if parkDistance <= radiusMeters {
			parkeringsSøk = append(parkeringsSøk, parkering)
			antallTreff++
			if antallTreff > 150 {
				break
			}
		}
	}
	var result []byte
	result, _ = json.Marshal(parkeringsSøk)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func apiForTekstSøk(w http.ResponseWriter, r *http.Request) {
	parkeringer := make([]ParkeringsOmraade, 0)
	getJSON("https://www.vegvesen.no/ws/no/vegvesen/veg/parkeringsomraade/parkeringsregisteret/v1/parkeringsomraade?datafelter=alle", &parkeringer)
	var parkeringsSøk []ParkeringsOmraade
	antallTreff := 0
	search := r.URL.Query().Get("search")
	ladestasjoner := r.URL.Query().Get("ladestasjoner")
	hc := r.URL.Query().Get("hc")

	for _, parkering := range parkeringer {
		lowerParkeringPoststed := strings.ToLower(parkering.AktivVersjon.Poststed)
		lowerSøk := strings.ToLower(search)
		lowerSøkParkeringNavn := strings.ToLower(parkering.AktivVersjon.Navn)

		if strings.Contains(lowerParkeringPoststed, lowerSøk) ||
			strings.Contains(lowerSøkParkeringNavn, lowerSøk) {

			if hc == "on" && parkering.AktivVersjon.AntallForflytnigshemmede == 0 {
				continue
			}

			if ladestasjoner == "on" && parkering.AktivVersjon.AntallLadeplasser == 0{
				continue
			}
			parkeringsSøk = append(parkeringsSøk, parkering)
			antallTreff++

			if antallTreff > 150 {
				break
			}
		}
	}
	var result []byte
	result, _ = json.Marshal(parkeringsSøk)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func main() {
	http.HandleFunc("/", parkeringsomraade)
	// JSON-APIer for å gjøre diverse søk
	http.HandleFunc("/api/parkering/radius", apiForRadiusSøk)
	http.HandleFunc("/api/parkering/search", apiForTekstSøk)
	// Gjør filer i public/ tilgjengelig fra websiden.
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("Oblig4/public/"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}