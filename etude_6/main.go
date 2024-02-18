package main

import (
	"log"

	"github.com/redmorin01/LivreGo-source/etude_6/elections/model"
)

func main() {
	m := model.FromMongo{Server: "mongodb://localhost:8088", DbName: "elections", PoliticiansCollection: "politicians", VotesCollection: "votes"}

	p, err := m.Winner()
	if err != nil {
		log.Printf("error: %s \n", err)
		return
	}
	log.Printf("Le gagnant est %s \n", p.String())
}
