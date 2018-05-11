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
	handler := http.HandlerFunc(ApiForTekstSøk)

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

func TestApiForOmrådeSøk(t *testing.T) {
	//
	req, err := http.NewRequest("GET", "/api/parkering/radius?radius=10&longitude=8.000012999999978&latitude=58.076023999999975&hc=&ladestasjoner=", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ApiForRadiusSøk)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":1860,"breddegrad":58.076024,"lengdegrad":8.000013,"aktivVersjon":{"navn":"Flekkerøy- flere delområder","adresse":"Østerøya 3","postnummer":"4625","poststed":"FLEKKERØY","antallLadeplasser":0,"antallAvgiftsbelagtePlasser":0,"antallAvgiftsfriePlasser":125,"antallForflytningshemmede":0,"vurderingForflytningshemmede":"Forflytningshemmede:\n\nFlere parkeringsområder hovedsakelig i turterreng. Ingen registrerte behov fra de forflytningshemmedes organisasjoner. \n\nLadestasjoner:\n\nHovedsakelig lokalisert i turterreng uten strømforsyning. Ikke registrert behov for ladestasjoner. ","typeParkeringsomrade":"AVGRENSET_OMRADE"},"Avstand":1.9295712076477326e-9}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}