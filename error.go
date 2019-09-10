package baas

import (
	"fmt"
)

type BaasError struct {
	code    int
	message string
}

var errorMap map[int]BaasError

func addError(code int, message string) {
	err := BaasError{code, message}
	errorMap[code] = err
}

func init() {
	errorMap = make(map[int]BaasError)
	addError(1, "Played the same card")
	addError(2, "No more card to play")
	addError(3, "Wrong playerID")
}

func NewError(code int) BaasError {
	return errorMap[code]
}

func (e BaasError) Error() string {
	return fmt.Sprintf("ERROR %d, %s\n", e.code, e.message)
}
