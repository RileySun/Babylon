package babylon

import(
	"testing"
	"strings"
)

func TestBabble(t *testing.T) {
	//Default 1 Count
	babbler := NewBabylon(1)
	
	//1 Count
	oneWord := babbler.Babble()
	if len(strings.Split(oneWord, " ")) != 1 {
		t.Error("Failed to produce one word", oneWord)
	}

	//10 Count
	babbler.Number = 100
	tenWords := babbler.Babble()
	if len(strings.Split(tenWords, " ")) != 100 {
		t.Error("Failed to produce ten words", tenWords)
	}
	
	//100 Count
	babbler.Number = 100
	hundredWords := babbler.Babble()
	if len(strings.Split(hundredWords, " ")) != 100 {
		t.Error("Failed to produce hundred words", hundredWords)
	}
	
	//1000 Count
	babbler.Number = 1000
	thousandWords := babbler.Babble()
	if len(strings.Split(thousandWords, " ")) != 1000 {
		t.Error("Failed to produce thousand words", thousandWords)
	}
}