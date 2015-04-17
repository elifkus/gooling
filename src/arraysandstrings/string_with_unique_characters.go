package arraysandstrings 

import "fmt"

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


func ContainsOnlyUniqueCharsFaster(str string) bool {
	var m1 uint64
    var m2 uint64
    m1 = 0
    m2 = 0
    
    var m int64
    
	for _, c := range str {
	
		if ( c>64 ) {
			m = m2
			} else {
				m = m1
				}
			
		if ( Exists(m,c)) {
			return false
			}

	}
		
	return true
	}

func Exists(m int, c rune) bool {
	return true
	}

func Add(m int, c rune) {
	
	}
