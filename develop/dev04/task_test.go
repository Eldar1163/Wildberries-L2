package main

import "testing"

func Test_findAnagrams(t *testing.T) {
	input := []string{
		"пятак", "листок", "пятка", "слиток", "тяпка", "столик",
	}
	expectedOutput := map[string][]string{
		"листок": {"слиток", "столик"},
		"пятак":  {"пятка", "тяпка"},
	}

	curOutput := findAnagrams(input)

	if !mapEquals(curOutput, expectedOutput) {
		t.Error("Wrong answer")
	}
}

func mapEquals(m1 map[string][]string, m2 map[string][]string) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if !sliceEquals(m2[k], v) {
			return false
		}
	}

	return true
}

func sliceEquals(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for ind, val := range s1 {
		if s2[ind] != val {
			return false
		}
	}

	return true
}
