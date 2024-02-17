package helper

import (
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

func TestComputeRound(t *testing.T) {
	computedRound := ComputeRound(model.Votes{
		model.Vote{Name: "John Doe", ID: "a", PoliticianID: 1},
		model.Vote{Name: "John Doe", ID: "b", PoliticianID: 0},
		model.Vote{Name: "John Doe", ID: "c", PoliticianID: 1},
		model.Vote{Name: "John Doe", ID: "d", PoliticianID: 2},
		model.Vote{Name: "John Doe", ID: "e", PoliticianID: 2},
		model.Vote{Name: "John Doe", ID: "f", PoliticianID: 2},
	})

	expectedRound := Round{0: 1, 1: 2, 2: 3}
	if !reflect.DeepEqual(computedRound, expectedRound) {
		t.Errorf("Round is not as expected. Computed: %v. Expected: %v", computedRound, expectedRound)
	}
}
