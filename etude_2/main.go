package main

import (
	"fmt"
	"log"
	"os/user"

	"github.com/redmorin01/LivreGo-source/etude_2/elections/helper"
	"github.com/redmorin01/LivreGo-source/etude_2/elections/model"
)

func main() {
	votesFilesNames := []string{}
	for i := 0; i < 100; i++ {
		votesFilesNames = append(votesFilesNames, fmt.Sprintf("votes_%d.json", i+1))
	}

	usr, err := user.Current()
	if err != nil {
		log.Printf("error1: %s \n", err)
		return
	}

	m := model.FromFiles{DirPath: usr.HomeDir + "/LivreGo-data", PoliticiansFileName: "politicians.json", VotesFileName: votesFilesNames}

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

	log.Printf("Le gagnant est %s \n", w.String())
}
