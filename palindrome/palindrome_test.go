package palindrome

import (
	"testing"
)

func Test(t *testing.T) {

	result := IsPalindrome("race car")

	if !result {
		t.Errorf("race car is a palindrome but false returned")
	}

	result = IsPalindrome("Race Car")

	if !result {
		t.Errorf("Race Car is a palindrome but false returned")
	}

	result = IsPalindrome("A Toyota! Race fast, safe car! A Toyota!")

	if !result {
		t.Errorf("'A Toyota! Race fast, safe car! A Toyota!' is a palindrome but false returned")
	}

	result = IsPalindrome("this Really is not a palindrome... for sure")

	if result {
		t.Errorf("'this Really is not a palindrome... for sure' is NOT a palindrome but true returned")
	}

	result = IsPalindrome("anna")

	if !result {
		t.Errorf("'anna' is a palindrome but false returned")
	}

	result = IsPalindrome("Aibohphobia")
	if !result {
		t.Errorf("'Aibohphobia' is a palindrome but false returned")
	}

}
