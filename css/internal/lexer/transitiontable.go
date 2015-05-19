
package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 35 : // ['#','#']
				return 2
			case r == 42 : // ['*','*']
				return 3
			case r == 43 : // ['+','+']
				return 4
			case r == 45 : // ['-','-']
				return 5
			case r == 46 : // ['.','.']
				return 6
			case r == 47 : // ['/','/']
				return 7
			case 48 <= r && r <= 57 : // ['0','9']
				return 8
			case r == 58 : // [':',':']
				return 9
			case r == 59 : // [';',';']
				return 10
			case r == 62 : // ['>','>']
				return 11
			case 65 <= r && r <= 90 : // ['A','Z']
				return 5
			case r == 95 : // ['_','_']
				return 5
			case r == 97 : // ['a','a']
				return 5
			case r == 98 : // ['b','b']
				return 12
			case 99 <= r && r <= 122 : // ['c','z']
				return 5
			case r == 123 : // ['{','{']
				return 13
			case r == 124 : // ['|','|']
				return 14
			case r == 125 : // ['}','}']
				return 15
			
			
			
			}
			return NoState
			
		},
	
		// S1
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			
			
			
			}
			return NoState
			
		},
	
		// S2
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S3
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S4
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S5
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 5
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 5
			case r == 95 : // ['_','_']
				return 5
			case 97 <= r && r <= 122 : // ['a','z']
				return 5
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 17
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 18
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 5
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 5
			case r == 95 : // ['_','_']
				return 5
			case 97 <= r && r <= 110 : // ['a','n']
				return 5
			case r == 111 : // ['o','o']
				return 21
			case 112 <= r && r <= 122 : // ['p','z']
				return 5
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 5
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 5
			case r == 95 : // ['_','_']
				return 5
			case 97 <= r && r <= 122 : // ['a','z']
				return 5
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 17
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 22
			
			
			default:
				return 18
			}
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 17
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 20
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 5
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 5
			case r == 95 : // ['_','_']
				return 5
			case 97 <= r && r <= 99 : // ['a','c']
				return 5
			case r == 100 : // ['d','d']
				return 23
			case 101 <= r && r <= 122 : // ['e','z']
				return 5
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 22
			case r == 47 : // ['/','/']
				return 24
			
			
			default:
				return 18
			}
			
		},
	
		// S23
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 5
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 5
			case r == 95 : // ['_','_']
				return 5
			case 97 <= r && r <= 120 : // ['a','x']
				return 5
			case r == 121 : // ['y','y']
				return 25
			case r == 122 : // ['z','z']
				return 5
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S25
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 5
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 5
			case r == 95 : // ['_','_']
				return 5
			case 97 <= r && r <= 122 : // ['a','z']
				return 5
			
			
			
			}
			return NoState
			
		},
	
}
