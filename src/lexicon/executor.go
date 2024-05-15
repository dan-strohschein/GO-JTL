package lexicon

import (
	"fmt"
	"strings"
)

func ExecuteCommand(cmd Command, operand string, json string) (string, string) {

	switch cmd {
	case FIND:
		println("Executing find")
		if operand == "." {
			return json, ""
		}
	}

	return "", "Failed!"
}

func find(operand string, json string) (string, string) {
	var errMessage string = ""
	//If the operand has a path, walk the path. Paths have to have a dot in them, and should always start with a dot
	if isPath(operand) {

	} else {
		errMessage = "The first character of a path must be the root, which is noted as a dot. Please put a dot in front of the path."
	}
	return "", fmt.Sprintf("Failed!! %s", errMessage)
}

func isPath(str string) bool {
	if strings.Contains(str, ".") {
		if str[0] == '.' {
			return true
		}

	}
	return false
}

func walkPath(path string, json string) any {
	return nil
}
