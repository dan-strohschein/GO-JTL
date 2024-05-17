package lexicon

import (
	"errors"

	jsonmarshaller "jtl/json-marshaller"
	"strings"
)

type Command int
type OperandType int

type QueryToken struct {
	CND     Command
	Operand string
}

// ------------------------------------ Command Enums --------------------------------------

const (
	FIND     = iota //When given a tree structure property name, returns the object of that property. A dot returns the whole root object
	SEEK            // searches the entire object graph for the specified property name
	REMOVE          //Deletes the entire property from the object tree
	REPLACE         // Replaces any property value with another value
	CHANGE          // Changes an property value - only for primitives
	WHERE           //Set up for filtering criteria
	COUNT           //For arrays, search for arrays that have a particular count, or return the count of the array elements
	GET             //Returns a specific part of the GRAPH
	VALUE           //Indicates the executor to use the VALUE part of the KVP
	PROPERTY        // Indicates the executor to use the PROPERTY part of the KVP
	CHILDREN        // Indicates to return all of the children of the current key
	LIKE            // Indicates that matches should be wildcard based

)

var (
	CommandsMap = map[string]Command{
		"FIND":     FIND,
		"SEEK":     SEEK,
		"REMOVE":   REMOVE,
		"REPLACE":  REPLACE,
		"CHANGE":   CHANGE,
		"WHERE":    WHERE,
		"COUNT":    COUNT,
		"GET":      GET,
		"VALUE":    VALUE,
		"PROPERTY": PROPERTY,
		"CHILDREN": CHILDREN,
		"LIKE":     LIKE,
	}
)

func (cmd Command) String() string {
	return [...]string{"FIND", "SEEK", "REMOVE", "REPLACE", "CHANGE", "WHERE", "COUNT", "Get", "Value", "Property", "Children", "Like"}[cmd-1]
}

func (cmd Command) EnumIndex() int {
	return int(cmd)
}

func ParseCommandString(str string) (Command, bool) {
	c, ok := CommandsMap[strings.ToUpper(str)]
	return c, ok
}

// ----------------------------------- OperandType -------------------------------------------------

const (
	String  = iota //
	Number         //
	Object         //
	Array          //
	Boolean        //
	Path
)

var (
	OperandsTypeMap = map[string]OperandType{
		"String":  String,
		"Number":  Number,
		"Object":  Object,
		"Array":   Array,
		"Boolean": Boolean,
		"Path":    Path,
	}
)

func (ot OperandType) String() string {
	return [...]string{"String", "Number", "Object", "Array", "Boolean", "Path"}[ot]
}

func (ot OperandType) EnumIndex() int {
	return int(ot)
}

// func ParseCommandString(str string) (Command, bool) {
// 	c, ok := CommandsMap[strings.ToUpper(str)]
// 	return c, ok
// }

// ----------------------------------- Query Parser ------------------------------------------------
func CheckQuery(query string, tree jsonmarshaller.JSONTreeNode) (string, error) {
	var queryParts = strings.Split(query, " ")

	//Validate that the string is a legit json object

	for i := 0; i < len(queryParts); i++ {
		var token, _ = ParseCommandString(queryParts[i])
		switch token {
		case FIND:
			// There must be an operand after the FIND command.
			if i+1 > len(queryParts) {
				return "", errors.New("there must be an operand or command after the FIND keyword ")
			}

			var operandWord = queryParts[i+1]

			switch operandWord[0] {
			case '[':
				// array. Recurse because arrays can hold the same types recursively
			case '{':
				// object
			case '"':
				//string
			case 't':
				if operandWord == "true" {
					//bool
				}
			case 'f':
				if operandWord == "false" {
					//bool
				}
			default: // numbers
			}
		case SEEK:

		default:
			// freebsd, openbsd,
			// plan9, windows...
			//fmt.Printf("%s.\n", os)
		}

	}

	return "", nil
}

// func DetermineOperandType( operandWord string) OperandType {
// 	switch operandWord[0] {
// 	case '[':
// 		// array. Dig for deeper type leveling
// 	case '{':
// 		// object
// 	case '"':
// 		//string
// 	case 't':
// 		if operandWord == "true" {
// 			//bool
// 		}
// 	case 'f':
// 		if operandWord == "false" {
// 			//bool
// 		}
// 		default: // numbers
// 	}
// 	return nil
// }
