package model

import "fmt"

// FromFiles holds the information and instruments the liaison of the model with flat JSON files
type FromFiles struct {
	DirPath             string
	PoliticiansFileName string
	VotesFileName       []string
}

// Politician contains all data about one given politician
type Politician struct {
	Name  string `json:"name" bson:"name"`
	ID    int    `json:"id,omitempty" bson:"id,omitempty"`
	Party string `json:"party,omitempty" bon:"party,omitempty"`
}

// Politician is a set of politicians
type Politicians []Politician

// Vote is the information registered when a voter votes
type Vote struct {
	Name         string `json:"name" bson:"name"`
	ID           string `json:"id" bson:"id"`
	PoliticianID int    `json:"politician_id" bson:"politician_id"`
}

// Votes is a set of votes
type Votes []Vote

func (p *Politician) String() string {
	return fmt.Sprintf("%s, de \"%s\"", (*p).Name, (*p).Party)
}

func (v *Vote) String() string {
	return v.Name
}
