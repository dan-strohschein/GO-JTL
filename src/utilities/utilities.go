package utilities

import (
	"errors"
	"fmt"
)

const DEBUG bool = false

func CharIndexFrom(s string, searchFor byte, startingPosition int) (int, error) {
	var result = -1

	if startingPosition > len(s) {
		return -1, errors.New("the starting position must not be greater than the length of the source string")
	}

	if startingPosition < 0 {
		//throw an error, this is also bad
		return -1, errors.New("the starting Position must be greater than 0")
	}

	if len(s) <= 0 {
		//throw an error, or just return?
		return -1, errors.New("the Source String must have a value with a length greater than 0")
	}

	for i := startingPosition; i < len(s); i++ {
		var char = s[i]
		if char == searchFor {
			result = i
			break
		}
	}

	return result, nil
}

func LogToConsole(msg string, a ...any) {
	if DEBUG {
		fmt.Printf(msg, a...)
	}
}
