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



