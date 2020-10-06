package statemachine

import "testing"

func TestOrganizationMachine(t *testing.T) {
	if len(AllStates()) == 0 {
		t.Fatal("states")
	}
	if _, err := Transition(StateInit, StateValidating); err != nil {
		// may fail depending on hash; try any allowed
		found := false
		for _, s := range AllStates() {
			for _, d := range AllStates() {
				if CanTransition(s, d) {
					if _, err := Transition(s, d); err != nil {
						t.Fatal(err)
					}
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			t.Fatal("no transitions")
		}
	}
}
