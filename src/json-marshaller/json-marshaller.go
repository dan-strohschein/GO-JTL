package jsonmarshaller

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type JSONType int
type Command int
type Tokens int

// ----------------------------------- JSON object Tokens -----------------------------------------
const (
	OpenProperty = iota
	CloseProperty
	OpenValue
	CloseValue
	OpenObject
	CloseObject
	Comma
	OpenArray
	CloseArray
	Colon
)

func (t Tokens) String() string {
	return [...]string{"OpenProperty", "CloseProperty", "OpenValue", "CloseValue", "OpenObject", "CloseObject", "Comma", "OpenArray", "CloseArray", "Colon"}[t]
}

func (t Tokens) EnumIndex() int {
	return int(t)
}

// ------------------------------------ JSONType Enums --------------------------------------

const (
	Object = iota
	String = 2
	Number = 3
	Array  = 4
	Null   = 5
	CMD    = 6
	Bool   = 7
)

var (
	JSONTypesMap = map[string]JSONType{
		"Object": Object,
		"String": String,
		"Number": Number,
		"Array":  Array,
		"NULL":   Null,
		"CMD":    CMD,
		"Bool":   Bool,
	}
)

func (jt JSONType) String() string {
	return [...]string{"Object", "String", "Number", "Array", "NULL", "Command", "Bool"}[jt-1]
}

func (jt JSONType) EnumIndex() int {
	return int(jt)
}

func ParseJSONTypeString(str string) (JSONType, bool) {
	c, ok := JSONTypesMap[strings.ToLower(str)]
	return c, ok
}

type KeyValuePair struct {
	Property string
	Value    any
	Jtype    JSONType
}

type JSONTreeNode struct {
	Properties []KeyValuePair
}

// //When given an entire JSON Object as a string, if the first character is a { and the last is } then it is OK
func IsLegitJSONRoot(json string) bool {
	var chars = strings.Split(json, "")
	if len(chars) > 0 {
		if chars[0] == "{" {
			if chars[len(chars)-1] == "}" {
				return true
			}
		}
	}

	return false
}

func MarshallJSON(inputJSON string) JSONTreeNode {

	//Clean the string by killing spaces
	var json, err = CleanJSONString(inputJSON)

	if err != nil {
		panic(err)
	}
	var currentToken Tokens = OpenObject
	var currentValue string
	var startingChar = 0

	if string(json[0]) == "}" {
		startingChar = 1
	}
	rootNode := JSONTreeNode{}
	var properties = make([]KeyValuePair, 0)
	var NewKeyValuePair = KeyValuePair{}

	for i := startingChar; i < len(json); i++ {
		switch c := json[i]; c {
		case '{':
			//Opening an object

			if currentToken == Colon {
				currentToken = OpenObject
				println("Opening up an object")

				// Find the next instance of the object closing tag
				var objectEndIndex, err = CharIndexFrom(json, '}', i)

				if err != nil {
					panic(err)
				}

				if objectEndIndex > -1 {
					//Only pass what we need to pass
					var objectString = json[i+1 : objectEndIndex+1]

					// Recurse into the sub-object
					var finishedObject = MarshallJSON(objectString)

					//Set the values and types from the returned recursion
					NewKeyValuePair.Value = finishedObject
					NewKeyValuePair.Jtype = Object

					//Move our character pointer to the closing bracket
					i = objectEndIndex
					currentToken = CloseObject

				}
			}

		case '}':
			println("Closing object")
			if currentToken == Colon {
				NewKeyValuePair.Jtype = DetermineJSonType(currentValue)
			}

			currentToken = CloseObject
			NewKeyValuePair.Value = currentValue
			properties = append(properties, NewKeyValuePair)
			currentValue = ""
		case '"':
			//Either a property name or a string value
			if currentToken == OpenObject {
				currentToken = OpenProperty
				println("Opening up a Property")
			} else if currentToken == Comma {
				currentToken = OpenProperty
				println("Opening up a Property after comma")
				currentValue = ""
			} else if currentToken == OpenProperty {
				//Set the property name on the objectnode
				println("Closing a property")
				currentToken = CloseProperty
				NewKeyValuePair.Property = currentValue
				currentValue = ""
			} else if currentToken == OpenArray {
				// This is an array of strings, so we don't do anything here
			} else if currentToken == Colon {
				currentToken = OpenValue
				println("OPening a string value")

			} else if currentToken == OpenValue {
				//Store the value in the objectNode
				println("Closing a string value")
				currentToken = CloseValue
				NewKeyValuePair.Value = currentValue
				NewKeyValuePair.Jtype = String
			}
		case '[':
			//an array starts here
			println("Opening an array")
			currentToken = OpenArray

			var arrayEndIndex, err = CharIndexFrom(json, ']', i)

			if err != nil {
				panic(err)
			}

			var arrayString = json[i+1 : arrayEndIndex]
			NewKeyValuePair.Value = arrayString
			NewKeyValuePair.Jtype = Array
			i = arrayEndIndex
			currentToken = CloseArray
		case ']':
			println("Closing an array")
			currentToken = CloseArray
		case ':':
			// The end of a property
			println("Found colon")
			currentToken = Colon
		case ',':
			println("Found comma")

			if currentToken == Colon {
				NewKeyValuePair.Value = currentValue
				NewKeyValuePair.Jtype = DetermineJSonType(currentValue)
			}

			properties = append(properties, NewKeyValuePair)
			currentToken = Comma
			currentValue = ""
		default:
			//Accumulate the characters for properties/values
			currentValue += string(c)
		}
	}

	for j := 0; j < len(properties); j++ {
		fmt.Printf("%s : %s \n", properties[j].Property, properties[j].Value)
	}
	rootNode.Properties = properties

	return rootNode
}

func DetermineJSonType(str string) JSONType {

	var numericCheck = regexp.MustCompile(`^[0-9]+$`)
	if numericCheck.MatchString(str) {
		return Number
	} else if strings.ToLower(str) == "true" || strings.ToLower(str) == "false" {
		return Bool
	} else if strings.Contains(str, "[") {
		return Array
	} else {
		return Null
	}
}

func CleanJSONString(json string) (string, error) {
	if len(json) <= 0 {
		return "", errors.New("the json string must have a length longer than 0")
	}

	//Clean the string by killing spaces
	json = strings.ReplaceAll(json, " : ", ":")
	json = strings.ReplaceAll(json, ", ", ",")
	json = strings.ReplaceAll(json, " }", "}")
	json = strings.ReplaceAll(json, "{ ", "{")

	return json, nil
}

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
