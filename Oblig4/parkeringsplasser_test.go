package main

import (
	"testing"
	"math"
	"net/http"
	"net/http/httptest"
)

func TestHsin(t *testing.T) {
	expected := math.Pow(math.Sin(5/2),2)
	actual := math.Pow(math.Sin(5/2),2)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s' got: '%s'", expected, actual)
	}
}

func TestApiForTekstSøk(t *testing.T) {
	//
	req, err := http.NewRequest("GET", "/api/parkering/search/?search=tommelitenvei", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(apiForTekstSøk)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":7385,"breddegrad":58.160647,"lengdegrad":8.068607,"aktivVersjon":{"navn":"Søm - Tommelitenvei","adresse":"Tommelitenvei 23","postnummer":"4638","poststed":"KRISTIANSAND S","antallLadeplasser":0,"antallAvgiftsbelagtePlasser":0,"antallAvgiftsfriePlasser":6,"antallForflytningshemmede":0,"vurderingForflytningshemmede":"","typeParkeringsomrade":"AVGRENSET_OMRADE"},"Avstand":0}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
