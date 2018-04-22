package main

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
)

func pathOne(w http.ResponseWriter, r *http.Request){
	res, err := http.Get("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
	if err != nil {
		panic(err.Error())
	}

	type miljo struct {
	Name string `json:"navn"`
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	garbage1 := miljo{}
	jsonErr := json.Unmarshal(body, &garbage1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(garbage1.Name)
}

func main() {
http.HandleFunc("/1", pathOne)
if err := http.ListenAndServe(":8080", nil); err != nil {
//panic(err)
}
}