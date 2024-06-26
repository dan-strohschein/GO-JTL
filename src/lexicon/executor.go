package lexicon

import (
	"errors"
	"fmt"
	jsonmarshaller "jtl/json-marshaller"
	"jtl/utilities"
	"strconv"
	"strings"
)

func ParseQuery(query string, tree jsonmarshaller.JSONTreeNode) (string, error) {
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

			//var operandWord = queryParts[i+1]

		case SEEK:

		default:
			// freebsd, openbsd,
			// plan9, windows...
			//fmt.Printf("%s.\n", os)
		}

	}

	return "", nil
}

func ExecuteCommand(cmd Command, operand string, json string) (string, error) {

	//We turn this string into an object
	var objectTree = jsonmarshaller.MarshallJSON(json)

	var errHandler error = nil
	var foundProp jsonmarshaller.KeyValuePair
	var jsonString string = ""

	switch cmd {
	case FIND:
		// There must be an operand after the FIND command.
		if len(operand) <= 0 {
			return "", errors.New("there must be an operand or command after the FIND keyword ")
		}

		//Execute the find with the operand
		foundProp, errHandler = Find(objectTree, operand)
		utilities.LogToConsole("(ExecuteCommand) Found the following property : %s\n", foundProp.Property)

		if errHandler != nil {
			//fmt.Printf("ERROR! :: %s", errHandler)
			return "", errHandler
		}

		if len(foundProp.Property) > 0 {
			jsonString, errHandler = foundProp.ConvertString()
			if errHandler != nil {
				fmt.Printf("ERROR! :: %s", errHandler)
			}
		}

		utilities.LogToConsole("(ExecuteCommand) transformed into json string : %s\n", jsonString)
		//var operandWord = operand[0]

		// switch operandWord {
		// case '[':
		// 	// array. Recurse because arrays can hold the same types recursively
		// case '{':
		// 	// object
		// 	Find(objectTree, operand, lexicon.Object)
		// case '"':
		// 	//string
		// case 't':
		// 	if operand == "true" {
		// 		//bool
		// 	}
		// case 'f':
		// 	if operand == "false" {
		// 		//bool
		// 	}
		// case '.':
		// 	//path
		// default: // numbers
		// }
	case SEEK:

	default:
		// freebsd, openbsd,
		// plan9, windows...
		//fmt.Printf("%s.\n", os)
	}

	return jsonString, errHandler
}

func Find(objectTree jsonmarshaller.JSONTreeNode, operand string) (jsonmarshaller.KeyValuePair, error) {
	var errMessage string = ""
	var err error
	var result jsonmarshaller.KeyValuePair

	//TODO Deal with the possibility of having different types of operands in the path, like arrays
	//Paths have to have a dot in them, and should always start with a dot
	if isPath(operand) {
		result, err = GetPropertiesFromTreeWithPath(objectTree, operand)
		utilities.LogToConsole("(Find) Found the following property : %s\n", result.Property)
	} else {
		errMessage = "The first character of a path must be the root, which is noted as a dot. Please put a dot in front of the path."
		return result, fmt.Errorf("%s", errMessage)
	}

	return result, err
}

func GetPropertiesFromTreeWithPath(tree jsonmarshaller.JSONTreeNode, path string) (jsonmarshaller.KeyValuePair, error) {
	var result jsonmarshaller.KeyValuePair
	var isArray bool = false
	var searchingForIndex int
	var searchingForProperty string
	var err error
	//remove the leading dot
	var currentPath string

	//TODO REfactor this block into a function called ParsePath or something -------------
	currentPath = path

	if currentPath[0] == '.' {
		currentPath = path[1:]
	}

	var pathSections = strings.Split(currentPath, ".")
	searchingForProperty = pathSections[0]

	if strings.Contains(pathSections[0], "[") {
		//The path for the current property is an array. The value in the array blocks is the INDEX of the array.
		isArray = true
		var leftBlockIndex = strings.Index(pathSections[0], "[")
		var rightBlockIndex, err = utilities.CharIndexFrom(searchingForProperty, ']', leftBlockIndex)
		var searchingForIndexString = string(pathSections[0][leftBlockIndex+1 : rightBlockIndex])
		searchingForIndex, err = strconv.Atoi(searchingForIndexString)

		if err != nil {
			//TODO This means it probably can't find the ] character. We should try to catch this in a validation step before hand
			panic(err)
		}

		searchingForProperty = string(pathSections[0][:leftBlockIndex])

		if err != nil {
			return result, err
		}
	}
	// DONE WITH refactor TODO Block ----------------------------------------------------------

	//Check the JSONTreeNode for the property in the current path
	for _, prop := range tree.Properties {
		utilities.LogToConsole("(GetPropertiesFromTreeWithPath) Looking for %s in object property %s \n", searchingForProperty, prop.Property)
		if strings.EqualFold(prop.Property, searchingForProperty) {
			//Found it! return the keyValuePair
			result = prop

			if isArray {
				// Because they asked for a specific array element, we are going to return a custom kvp
				var kvp = jsonmarshaller.KeyValuePair{}
				kvp.Property = prop.Property
				kvp.Jtype = Array
				var elements = strings.Split(result.Value.(string), ",")

				if len(pathSections) > 1 && len(result.Property) > 0 && result.Jtype == jsonmarshaller.Object {
					//Here's the hard part, it's possible that the search path looks like this: .Property1.SubProperty[3].AnObjectInTheArraysProperties
					//This means that the operand believes the array in question is an array of objects, and they want
					// a property from a SPECIFIC index. This is likely not going to be a popular use case, so ....
					//Let's do it later!
				} else {
					//Sanity check to make sure the index they asked for is IN this array
					if len(elements) > searchingForIndex {
						//Check to see if its a string
						if strings.Contains(elements[searchingForIndex], "\"") {
							kvp.Value = fmt.Sprintf("\"%s\"", elements[searchingForIndex])
						} else {
							kvp.Value = elements[searchingForIndex]
						}
					} else {
						//The array index is not in the array. We should tell the user that it wasn't found
						err = fmt.Errorf("the provided index '%d' for the array '%s' was not found", searchingForIndex, prop.Property)
						return kvp, err
					}
				}
				return kvp, err
			}
			break
		}
	}

	if len(pathSections) > 1 && len(result.Property) > 0 && result.Jtype == jsonmarshaller.Object {
		// The property was found, and its an object, and the path is requesting a property on this object, so recurse!
		result, err = GetPropertiesFromTreeWithPath(result.Value.(jsonmarshaller.JSONTreeNode), pathSections[1])
	}

	return result, err
}

func isPath(str string) bool {
	if strings.Contains(str, ".") {
		if str[0] == '.' {
			return true
		}

	}
	return false
}
