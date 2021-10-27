package fileProcessor

import (
	"testing"
	cmp "github.com/google/go-cmp/cmp"
)

func TestFindAndUpdate(t *testing.T) {
    searchIn := []StringInt{}
    data := []StringInt{}
	data = append(data, StringInt{Word: "a", Count: 3})
	data = append(data, StringInt{Word: "b", Count: 2})
	data = append(data, StringInt{Word: "a", Count: 1})

    result := FindAndUpdate(data, searchIn)
    i, exist := Find(result, StringInt{Word: "a", Count: 3} )

    if !exist || result[i].Count != 4 {
		t.Log("Eror ExpectedResult is diffrent!")
        t.Log("expected exist true, got", exist)
        t.Log("expected result[i].Count 4, got", 4)
		t.Fail()
	}
}

func TestFind(t *testing.T){
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
