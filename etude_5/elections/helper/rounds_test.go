package helper

import (
	"errors"
	"reflect"
	"testing"

	"github.com/redmorin01/LivreGo-source/etude_5/elections/model"
)

func TestAlwaysPasses(t *testing.T) {
	t.Log("This test always passes")
}

func TestAddition(t *testing.T) {
	computedInt := 2 + 2
	expectedInt := 4
	if computedInt != expectedInt {
		t.Errorf("Addition is not expected. Computed: %d. Expected: %d", computedInt, expectedInt)
	}
}

var computeRoundTests = []struct {
	in  model.Votes
	out Round
}{
	{model.Votes{
		model.Vote{Name: "John Doe", ID: "a", PoliticianID: 1},
		model.Vote{Name: "John Doe", ID: "b", PoliticianID: 0},
		model.Vote{Name: "John Doe", ID: "c", PoliticianID: 1},
		model.Vote{Name: "John Doe", ID: "d", PoliticianID: 2},
		model.Vote{Name: "John Doe", ID: "e", PoliticianID: 2},
		model.Vote{Name: "John Doe", ID: "f", PoliticianID: 2},
	}, Round{0: 1, 1: 2, 2: 3}},
	{model.Votes{
		model.Vote{Name: "John Doe", ID: "a", PoliticianID: 1},
		model.Vote{Name: "John Doe", ID: "b", PoliticianID: 0},
		model.Vote{Name: "John Doe", ID: "c", PoliticianID: 1},
		model.Vote{Name: "John Doe", ID: "d", PoliticianID: 2},
		model.Vote{Name: "John Doe", ID: "e", PoliticianID: 2},
		model.Vote{Name: "John Doe", ID: "f", PoliticianID: 3},
	}, Round{0: 1, 1: 2, 2: 2, 3: 1}},
	{model.Votes{}, Round{}},
}

func TestComputeRound(t *testing.T) {

	for _, testCase := range computeRoundTests {
		computedRound := ComputeRound(testCase.in)
		if !reflect.DeepEqual(computedRound, testCase.out) {
			t.Errorf("Round is not as expected. Computed: %v. Expected: %v", computedRound, testCase.out)
		}
	}
}

var WinnerTests = []struct {
	in              Round
	outPoliticianID int
	outErr          error
}{
	{
		in:              Round{1: 3, 2: 2, 3: 1},
		outPoliticianID: 1,
		outErr:          nil,
	}, {
		in:              Round{},
		outPoliticianID: 0,
		outErr:          errors.New("il ne semble pas y avoir de vote enregistré pour le moment"),
	},
	{
		in:              Round{1: 2, 2: 2, 3: 1},
		outPoliticianID: 0,
		outErr:          errors.New("deux candidat sont à égalité! John Doe, de \"GOP\" et John Doe, de \"GOP\" et ont tous deux 2 votes"),
	},
}

type mockedReader struct{}

func (m mockedReader) PoliticianFromID(ID int) (model.Politician, error) {
	return model.Politician{ID: ID, Name: "John Doe", Party: "GOP"}, nil
}

func (m mockedReader) AllPoliticians() (model.Politicians, error) {
	return nil, nil
}

func (m mockedReader) AllVotes() (model.Votes, error) {
	return nil, nil
}

func errorsUnequal(err1, err2 error) bool {
	return (err1 != nil && err2 != nil) && ((err1 == nil || err2 == nil) || (err1.Error() != err2.Error()))
}

func TestWinner(t *testing.T) {
	for _, testCase := range WinnerTests {
		computePolitician, err := testCase.in.Winner(mockedReader{})

		if errorsUnequal(err, testCase.outErr) {
			t.Errorf("Error status unexpected. Computed: %v. Expected: %v", err, testCase.outErr)
		}

		if computePolitician.ID != testCase.outPoliticianID {
			t.Errorf("Unexpected winning politician. Computed: %v. Expected: %v", computePolitician.ID, testCase.outPoliticianID)
		}
	}
}
