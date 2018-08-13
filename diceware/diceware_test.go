package diceware

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"testing"
)

func generate(testList, testAnswers, testRolls []string, rollCount int) error {
	rand.Seed(42)
	dw := NewDiceware(testList, testRolls)
	for i := 0; i < 6; i++ {
		pass, err := dw.GeneratePass(i+1, rollCount)
		if err != nil {
			log.Fatal(err)
		}
		if pass != testAnswers[i] {
			return errors.New(fmt.Sprintf("password\n\texpected %s\n\tgot      %s", testAnswers[i], pass))
		}
	}
	return nil
}

func TestGeneratePass(t *testing.T) {
	testList := []string{"a", "b", "c", "d", "e", "f"}
	testRolls := []string{"1", "2", "3", "4", "5", "6"}
	testPasswords := []string{"a", "cd", "ada", "cbdd", "ecadc", "dcbaec"}
	if err := generate(testList, testPasswords, testRolls, 1); err != nil {
		t.Error(err)
	}
	testList = []string{
		"a", "b", "c", "d", "e", "f",
		"aa", "ba", "ca", "da", "ea", "fa",
		"ab", "bb", "cb", "db", "eb", "fb",
		"ac", "bc", "cc", "dc", "ec", "fc",
		"ad", "bd", "cd", "dd", "ed", "fd",
		"ae", "be", "ce", "de", "ee", "fe",
	}
	testPasswords = []string{
		"c",
		"acac",
		"bbdccd",
		"ddbbbe",
		"ebadbdebeb",
		"cbaddcdcbc",
	}
	testRolls = []string{
		"11", "12", "13", "14", "15", "16",
		"21", "22", "23", "24", "25", "26",
		"31", "32", "33", "34", "35", "36",
		"41", "42", "43", "44", "45", "46",
		"51", "52", "53", "54", "55", "56",
		"61", "62", "63", "64", "65", "66",
	}
	if err := generate(testList, testPasswords, testRolls, 2); err != nil {
		t.Error(err)
	}
}
