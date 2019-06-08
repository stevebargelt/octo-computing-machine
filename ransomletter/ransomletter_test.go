package ransomletter

import (
	"testing"
)

func Test(t *testing.T) {
	newspaper := `
In other news, a ferret was stuck in a tree. 
A fireman saved him so he didn't fall. 
We are taking up a collection to pay the $700 fine. We already have $10 in pledges.
`
	result := IsRansomLetterReproducible("I want $1000 or the kitty dies.", newspaper)

	if !result {
		t.Errorf("Should get true")
	}

	result = IsRansomLetterReproducible("I want $100000 or the kitty dies.", newspaper)

	if result {
		t.Errorf("Should get false")
	}

}
