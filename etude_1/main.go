package main

import (
	"errors"
	"fmt"
	"log"
)

type politician struct {
	Name  string
	ID    int
	Party string
}

type voter struct {
	Name string
	ID   string
}

type votes map[voter]*politician
type round map[politician]int

func (p *politician) String() string {
	return fmt.Sprintf("%s, de \"%s\"", (*p).Name, (*p).Party)
}

func (v *voter) String() string {
	return v.Name
}

func (v *votes) computeRound() round {
	r := make(round)
	for _, p := range *v {
		val, exists := r[*p]
		if exists {
			r[*p] = val + 1
			continue
		}
		r[*p] = 1
	}
	return r
}

func (r *round) winner() (politician, error) {
	currentMaxScore := 0
	secondMaxScore := 0
	var currentWinner politician
	var secondToWinner politician

	for p, s := range *r {
		if s >= currentMaxScore {
			secondMaxScore = currentMaxScore
			currentMaxScore = s
			secondToWinner = currentWinner
			currentWinner = p

		}
	}

	if currentMaxScore == 0 {
		return politician{}, errors.New("Il ne semble y avoir aucun vote enregistré pour le moment")
	}

	if currentMaxScore == secondMaxScore {
		return politician{}, fmt.Errorf("deux candidats sont à égalité ! %s et %s ont tous deux %d votes.", currentWinner.String(), secondToWinner.String(), currentMaxScore)
	}
	return currentWinner, nil
}

func main() {
	blanc := politician{Name: "Vote blanc"}
	rouleau := politician{Name: "Chantal Rouleau", ID: 1, Party: "Ensemble demain"}
	lagace := politician{Name: "Sylvain Lagacé", ID: 2, Party: "Ecologie"}
	dorion := politician{Name: "Guillaume Dorion", ID: 3, Party: "Par et pour le peuple"}

	v := make(votes)
	v[voter{Name: "Alice", ID: "1"}] = &rouleau
	v[voter{Name: "Bob", ID: "2"}] = &rouleau
	v[voter{Name: "Charlie", ID: "3"}] = &dorion
	v[voter{Name: "David", ID: "4"}] = &dorion
	v[voter{Name: "Eve", ID: "5"}] = &dorion
	v[voter{Name: "Frank", ID: "6"}] = &dorion
	v[voter{Name: "Grace", ID: "7"}] = &dorion
	v[voter{Name: "Hector", ID: "8"}] = &lagace
	v[voter{Name: "Isabelle", ID: "9"}] = &rouleau

	r := v.computeRound()
	delete(r, blanc)
	w, err := r.winner()
	if err != nil {
		log.Printf("erreur: %s\n", err)
		return
	}
	fmt.Printf("Le gagnant est %s! avec %d votes.\n", w.String(), r[w])
}
