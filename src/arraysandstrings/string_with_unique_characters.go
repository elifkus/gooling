package arraysandstrings 


func ContainsOnlyUniqueChars(str string) bool {
	m := make(map[rune]int)
	
	for _, c := range str {
		_, exists := m[c]

		if (exists) {
			return false
			}
		
		m[c] = 1 

	}
		
	return true
	}


func ContainsOnlyUniqueCharsWithoutDataStructure(str string) bool {
	
	var substr string
	
	for i, c := range str {
		
		substr = str[i+1:]
		
		for _, o := range substr {
			if (c == o) {
				return false
				}
			}
		
	}
		
	return true
	}

func Exists(m int, c rune) bool {
	return true
	}

func Add(m int, c rune) {
	
	}
