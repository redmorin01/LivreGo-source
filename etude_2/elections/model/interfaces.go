package model

// Reader contains all the function to read the necessary data
type Reader interface {
	allPoliticians() (Politicians, error)
	PoliticianFromID(ID int) (Politician, error)
	AllVotes() (Votes, error)
}
