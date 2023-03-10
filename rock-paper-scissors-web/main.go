package main

import (
	"encoding/json"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")

}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "     ")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(result.Winner, result.ComputerChoice, result.RoundResult)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)

	log.Println("Stating web sever on prot 8080.")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println((err))
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println((err))
		return
	}
}
