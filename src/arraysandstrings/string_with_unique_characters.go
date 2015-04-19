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


func ContainsOnlyUniqueCharsForLowerAscii(str string) bool {
	var tally uint = 0
	var offset uint;
	var mask uint = 0
	
	for _, c := range str {
		
		offset = uint(c - 'a')
		
		mask = 0
		mask = 1 << offset 		
	
		if (tally & mask > 0) {
			return false
			}
		tally |= mask
		}
	
	return true
		
	}

