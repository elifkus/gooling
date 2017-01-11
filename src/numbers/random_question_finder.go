package numbers
	
import (
	"math/rand"
	"fmt"
	"time"
)

func FindQuestion() (int, int) {
	numberOfQuestionsByChapters := [17]int{9,8,6,12,8,10,12,14,8,11,6,11,8,7,7,26,26}
	
	sum := 0
	
	for _, numberOfQuestions := range numberOfQuestionsByChapters {
		sum += numberOfQuestions
	}
		
	rand.Seed(time.Now().UnixNano())
	
	number := rand.Intn(sum) + 1

	var n int 
	currentRunningSum := 0
	var previousRunningSum int
	
	for i, numberOfQuestions := range numberOfQuestionsByChapters {
		previousRunningSum = currentRunningSum
		currentRunningSum += numberOfQuestions
		n=i
		if currentRunningSum + 1 > number {
			break
		}
		
	}

	question := number - previousRunningSum
	chapter := n+1
	
	fmt.Printf("%d.%d\n", chapter,  question)
	
	return chapter, question
}

func FindQuestionInChapter(chapter int) int {
	numberOfQuestionsByChapters := [17]int{9,8,6,12,8,10,12,14,8,11,6,11,8,7,7,26,26}
	
	if (chapter < 1 || chapter > 17) {
		return -1;
	}
	
	rand.Seed(time.Now().UnixNano())
	
	question := rand.Intn(numberOfQuestionsByChapters[chapter-1]) + 1
	
	fmt.Printf("%d.%d\n", chapter,  question)

	return question
} 

