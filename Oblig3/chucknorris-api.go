package main

import ("net/http"
"text/template")

func chuckNorris(w http.ResponseWriter, r *http.Request) {
	foo1 := new(Astros)
	getJSON("https://api.chucknorris.io/jokes/random", foo1)

	t, _ := template.ParseFiles("Oblig3\\pepoleInSpace.html")
	t.Execute(w, foo1)
}
