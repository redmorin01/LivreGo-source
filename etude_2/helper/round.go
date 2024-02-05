package helper

import (
	"errors"
	"fmt"

	"github.com/redmorin01/LivreGo-source/etude_2/elections/model"
)

// Round is a count, for each politician, of the number of votes they got
type Round map[int]int

// ComputeRound computes the summary of the round
func ComputeRound(vs model.Votes) Round {
	r := make(Round)
	for _, v := range vs {
		val, exists := r[v.PoliticianID]
		if exists {
			r[v.PoliticianID] = val + 1
			continue
		}
		r[v.PoliticianID] = 1
	}
	return r
}

// Winner finds the winner from a round
func (r *Round) Winner(m model.Reader) (model.Politician, error) {
	currentMaxScore := 0
	secondMaxScore := 0
	var curreentWinner int
	var secondToWinner int

	for p, s := range *r {
		if s >= currentMaxScore {
			secondMaxScore = currentMaxScore
			currentMaxScore = s
			secondToWinner = curreentWinner
			curreentWinner = p
		}
	}

	if currentMaxScore == 0 {
		return model.Politician{}, errors.New("il ne semble pas y avoire de vote enregistré pour le moment")
	}

	if currentMaxScore == secondMaxScore {
		currentWinnerPolitician, err := m.PoliticianFromID(curreentWinner)
		if err != nil {
			return model.Politician{}, err
		}
		secondToWinnerPolitician, err := m.PoliticianFromID(secondToWinner)
		if err != nil {
			return model.Politician{}, err
		}

		errString := fmt.Sprintf("Deux candidat sont à égalité! %s et %s et ont tous deux %d votes", currentWinnerPolitician.String(), secondToWinnerPolitician.String(), currentMaxScore)
		return model.Politician{}, errors.New(errString)
	}

	currentWinnterPolitician, err := m.PoliticianFromID(curreentWinner)
	if err != nil {
		return model.Politician{}, err
	}

	return currentWinnterPolitician, nil
}
