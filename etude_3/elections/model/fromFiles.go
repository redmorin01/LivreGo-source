package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sync"
)

// FromFiles holds the information and instruments the liaison of the model with flat JSON files
type FromFiles struct {
	DirPath             string
	PoliticiansFileName string
	VotesFileName       []string
}

// Politician contains all data about one given politician
type Politician struct {
	Name  string `json:"name"`
	ID    int    `json:"id,omitempty"`
	Party string `json:"party,omitempty"`
}

// Politician is a set of politicians
type Politicians []Politician

// Vote is the information registered when a voter votes
type Vote struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	PoliticianID int    `json:"politician_id"`
}

// Votes is a set of votes
type Votes []Vote

func (p *Politician) String() string {
	return fmt.Sprintf("%s, de \"%s\"", (*p).Name, (*p).Party)
}

func (v *Vote) String() string {
	return v.Name
}

var allPoliticians Politicians

// AllPoliticians fetches all politicians from JSON file if cache is empty, returns the cache otherwise
func (m *FromFiles) AllPoliticians() (Politicians, error) {
	if allPoliticians == nil {
		file, err := os.ReadFile(path.Join(m.DirPath, m.PoliticiansFileName))
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(file, &allPoliticians); err != nil {
			return nil, err
		}
	}

	return allPoliticians, nil
}

// PoliticianFromID finds the Politician value when giver its ID
func (m *FromFiles) PoliticianFromID(ID int) (Politician, error) {
	ps, err := m.AllPoliticians()
	if err != nil {
		return Politician{}, err
	}

	for _, p := range ps {
		if p.ID == ID {
			return p, nil
		}
	}

	return Politician{}, fmt.Errorf("Politician of ID %d not found", ID)
}

var allVotes Votes

// AllVotes fetches all votes from JSON file if cache is empty, returns the cache otherwise
func (m *FromFiles) AllVotes() (Votes, error) {
	if allVotes == nil {
		wg := &sync.WaitGroup{}
		mutex := &sync.Mutex{}
		allVotes = Votes{}

		errs := make(chan error, len(m.VotesFileName))
		wg.Add(len(m.VotesFileName))
		for _, filename := range m.VotesFileName {
			fileNameRef := &filename

			go func() {
				defer wg.Done()
				var votesInFile Votes
				file, err := os.ReadFile(path.Join(m.DirPath, *fileNameRef))
				if err != nil {
					errs <- err
					return
				}
				if err = json.Unmarshal(file, &votesInFile); err != nil {
					errs <- err
					return
				}
				mutex.Lock()
				defer mutex.Unlock()
				allVotes = append(allVotes, votesInFile...)
				errs <- nil
			}()
		}
		wg.Wait()
		close(errs)
		for err := range errs {
			if err != nil {
				return nil, err
			}
		}
	}
	return allVotes, nil
}
