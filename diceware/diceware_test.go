package diceware

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

func generate(testList, testRolls []string, rollCount int) error {
	dw := NewDiceware(testList, testRolls)
	for i := 1; i < 7; i++ {
		pass, err := dw.GeneratePass(i, rollCount)
		if err != nil {
			log.Fatal(err)
		}
		if len(pass) != i {
			return errors.New(fmt.Sprintf("password length expected %d, got %d (%s)", i, len(pass), pass))
		}
	}
	return nil
}

func TestGeneratePass(t *testing.T) {
	testList := []string{"a", "b", "c", "d", "e", "f"}
	testRolls := []string{"1", "2", "3", "4", "5", "6"}
	if err := generate(testList, testRolls, 1); err != nil {
		t.Error(err)
	}
	testList = []string{
		"a", "b", "c", "d", "e", "f",
		"a", "b", "c", "d", "e", "f",
		"a", "b", "c", "d", "e", "f",
		"a", "b", "c", "d", "e", "f",
		"a", "b", "c", "d", "e", "f",
		"a", "b", "c", "d", "e", "f",
	}
	testRolls = []string{
		"11", "12", "13", "14", "15", "16",
		"21", "22", "23", "24", "25", "26",
		"31", "32", "33", "34", "35", "36",
		"41", "42", "43", "44", "45", "46",
		"51", "52", "53", "54", "55", "56",
		"61", "62", "63", "64", "65", "66",
	}
	if err := generate(testList, testRolls, 2); err != nil {
		t.Error(err)
	}
}
