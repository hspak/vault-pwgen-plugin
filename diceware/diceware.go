// Inspired by https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases
package diceware

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Diceware struct {
	wordlist map[string]string
}

func NewDiceware(wordList, rollList []string) *Diceware {
	list := make(map[string]string)
	for i, roll := range rollList {
		list[roll] = wordList[i]
	}
	return &Diceware{
		wordlist: list,
	}
}

// GeneratePass takes an integer for the number of words to use for the password.
// EFF recommends 6 for most cases, though this doesn't enforce it.
func (d *Diceware) GeneratePass(words, rolls int) (string, error) {
	if words < 1 {
		return "", errors.New("cannot generate password with less than 1 word")
	}
	if rolls < 1 {
		return "", errors.New("cannot generate password with less than 1 roll")
	}
	var password strings.Builder
	for i := 0; i < words; i++ {
		var rollIndex strings.Builder
		for j := 0; j < rolls; j++ {
			roll := rand.Intn(5) + 1
			if _, err := rollIndex.WriteString(strconv.Itoa(roll)); err != nil {
				return "", err
			}
		}
		word := d.wordlist[rollIndex.String()]
		if word == "" {
			msg := fmt.Sprintf("missing word for roll: %s", rollIndex.String())
			return "", errors.New(msg)
		}
		if _, err := password.WriteString(word); err != nil {
			return "", err
		}
	}
	return password.String(), nil
}
