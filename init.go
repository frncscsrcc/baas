package baas

var symbolToScore map[rune]int
var symbolToRelativeScore map[rune]int

func init() {
	symbolToScore = make(map[rune]int)
	symbolToScore['2'] = 0
	symbolToScore['3'] = 10
	symbolToScore['4'] = 0
	symbolToScore['5'] = 0
	symbolToScore['6'] = 0
	symbolToScore['7'] = 0
	symbolToScore['J'] = 2
	symbolToScore['Q'] = 3
	symbolToScore['K'] = 4
	symbolToScore['A'] = 11

	symbolToRelativeScore = make(map[rune]int)
	symbolToRelativeScore['2'] = 2
	symbolToRelativeScore['4'] = 4
	symbolToRelativeScore['5'] = 5
	symbolToRelativeScore['6'] = 6
	symbolToRelativeScore['7'] = 7
}
