package numbers

import (
	"testing"
)



func TestFindQuestion(t *testing.T)  {
	numberOfQuestionsByChapters := [17]int{9,8,6,12,8,10,12,14,8,11,6,11,8,7,7,26,26}
	
	chapter, question := FindQuestion()
	
	if chapter < 1 || chapter > 17 {
		t.Errorf("Chapter number cannot be larger than 17, or smaller than 1/nChapter number is %s", chapter)
	}
	
	if question < 1 || question > numberOfQuestionsByChapters[chapter-1] {
		t.Errorf("Question number for chapter %d cannot be larger than %d, or smaller than 1\nQuestion number is %d", 
						chapter, numberOfQuestionsByChapters[chapter-1], question)
	}
}

func TestFindQuestionInChapter(t *testing.T)  {
	numberOfQuestionsByChapters := [17]int{9,8,6,12,8,10,12,14,8,11,6,11,8,7,7,26,26}
	
	chapter := 17
	
	question := FindQuestionInChapter(chapter)
	
	if question < 1 || question > numberOfQuestionsByChapters[chapter-1] {
		t.Errorf("Question number for chapter %d cannot be larger than %d, or smaller than 1\nQuestion number is %d", 
						chapter, numberOfQuestionsByChapters[chapter-1], question)
	}
}

