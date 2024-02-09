package main

import (
	"fmt"
	"log"

	"github.com/redmorin01/LivreGo-source/etude_4/elections/helper"
	"github.com/redmorin01/LivreGo-source/etude_4/elections/model"
)

func main() {
	votesFilesNames := []string{}
	for i := 0; i < 100; i++ {
		votesFilesNames = append(votesFilesNames, fmt.Sprintf("votes_%d.json", i+1))
	}

	m := model.FromMongo{Server: "mongodb://localhost:8088", DbName: "elections", PoliticiansCollection: "politicians", VotesCollection: "votes"}

	allVotes, err := m.AllVotes()
	if err != nil {
		log.Printf("error2: %s \n", err)
		return
	}

	r := helper.ComputeRound(allVotes)
	delete(r, 0)

	w, err := r.Winner(&m)
	if err != nil {
		log.Printf("error3: %s \n", err)
		return
	}

	log.Println(len(allVotes))
	log.Printf("Le gagnant est %s \n", w.String())
}
