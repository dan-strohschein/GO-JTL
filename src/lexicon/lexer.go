package lexicon

import (
	"strings"
)

type Command int

type QueryToken struct {
	CND     Command
	Operand string
}

// ------------------------------------ Command Enums --------------------------------------

const (
	FIND    = iota //When given a tree structure property name, returns the object of that property. A dot returns the whole root object
	SEEK           // searches the entire object graph for the specified property name
	REMOVE         //Deletes the entire property from the object tree
	REPLACE        // Replaces any property value with another value
	CHANGE         // Changes an property value - only for primitives
	WHERE          //Set up for filtering criteria
	COUNT          //For arrays, search for arrays that have a particular count, or return the count of the array elements
)

var (
	CommandsMap = map[string]Command{
		"FIND":    FIND,
		"SEEK":    SEEK,
		"REMOVE":  REMOVE,
		"REPLACE": REPLACE,
		"CHANGE":  CHANGE,
		"WHERE":   WHERE,
		"COUNT":   COUNT,
	}
)

func (cmd Command) String() string {
	return [...]string{"FIND", "SEEK", "REMOVE", "REPLACE", "CHANGE", "WHERE", "COUNT"}[cmd-1]
}

func (cmd Command) EnumIndex() int {
	return int(cmd)
}

func ParseCommandString(str string) (Command, bool) {
	c, ok := CommandsMap[strings.ToLower(str)]
	return c, ok
}

// ----------------------------------- Query Parser ------------------------------------------------
func ParseQueryString(query string) (string, string) {
	var words = strings.Split(query, " ")

	//Validate that the string is a legit json object

	for i := 0; i < len(words); i++ {

		// switch token := w; token {
		// case  :
		// 	fmt.Println("OS X.")
		// case "linux":
		// 	fmt.Println("Linux.")
		// default:
		// 	// freebsd, openbsd,
		// 	// plan9, windows...
		// 	fmt.Printf("%s.\n", os)
		// }

	}

	return "", ""
}
