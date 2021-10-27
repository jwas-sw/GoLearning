package fileProcessor

import (
	"testing"
	cmp "github.com/google/go-cmp/cmp"
)

func TestFund(t *testing.T){
	expected := []StringInt{}
	expected = append(expected, StringInt{Word: "a", Count: 3})
	expected = append(expected, StringInt{Word: "b", Count: 2})
	expected = append(expected, StringInt{Word: "c", Count: 1})

    searching := StringInt{Word: "a", Count: 3}

    _, exist := Find(expected, searching)
	if !exist {
		t.Log("Eror ExpectedResult is diffrent!")
        t.Log("expected", true)
        t.Log("result", exist)
		t.Fail()
	}
}

func TestWordCount(t *testing.T) {
	result := WordCount("a b c a b a")

	expected := []StringInt{}
	expected = append(expected, StringInt{Word: "a", Count: 3})
	expected = append(expected, StringInt{Word: "b", Count: 2})
	expected = append(expected, StringInt{Word: "c", Count: 1})


	if !cmp.Equal(result, expected) {
		t.Log("Eror ExpectedResult is diffrent!")
        t.Log("expected", expected)
        t.Log("result", result)
		t.Fail()
	}
}
