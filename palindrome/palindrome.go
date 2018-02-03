package palindrome

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func IsPalindrome(pal string) bool {

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}

	pal = reg.ReplaceAllString(pal, "")
	pal = strings.ToLower(pal)

	isPal := true
	lenPal := len(pal)

	half := lenPal / 2
	fmt.Println(half)
	//O(n) time
	for i := 0; i <= half; i++ {
		if pal[i] != pal[lenPal-1-i] {
			isPal = false
			break
		}
	}

	return isPal

}

// race car
// racecar  (7)
// r     r  (1 and 7)
//  a   a   (2 and 6)
//   c c    (3 and 5)
// doesn't matter

// anna			(4)
// a  a     (1 and 4)
//  nn      (2 and 3)
// done
