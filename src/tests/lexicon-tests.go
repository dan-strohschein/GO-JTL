package tests

import (
	"fmt"
	jsonmarshaller "jtl/json-marshaller"
	"jtl/lexicon"
	"strings"
)

func TestAll() {
	println(evalTestResult(TestIsLegitRootNode()))
	println(evalTestResult(TestIsLegitRootNodeOnlyOpening()))
	println(evalTestResult(TestIsLegitRootNodeOnlyClosing()))

	println(evalTestResult(TestMarhsallingJSON()))
	println(evalTestResult(TestUnMarhsallingJSON()))

	println(evalTestResult(TestFindCommand()))
	println(evalTestResult(TestFindCommandWithNestedPath()))
	println(evalTestResult(TestFindCommandWithNestedPathAndArrayIndex()))
}

func TestIsLegitRootNode() bool {
	println("Testing IsLegitRootNode...")
	var expected = true
	var actual = jsonmarshaller.IsLegitJSONRoot("{}")

	//fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

func TestIsLegitRootNodeOnlyOpening() bool {
	println("Testing TestIsLegitRootNodeOnlyOpening...")
	var expected = false
	var actual = jsonmarshaller.IsLegitJSONRoot("{")

	//fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

func TestIsLegitRootNodeOnlyClosing() bool {
	println("Testing TestIsLegitRootNodeOnlyClosing...")
	var expected = false
	var actual = jsonmarshaller.IsLegitJSONRoot("}")

	//fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

// Check the find command
func TestFindCommand() bool {
	println("Testing TestFindCommand....")
	var expected = "\"someObject\":{\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }"
	var jsonString = "{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnotherStringType\" : \"Testing\" }"
	var actual, err = lexicon.ExecuteCommand(lexicon.FIND, ".someObject", jsonString)

	if err != nil {
		fmt.Printf("ERROR :: %s", err)
		return false
	}

	expected = strings.ReplaceAll(expected, " ", "")
	actual = strings.ReplaceAll(actual, " ", "")

	// fmt.Printf("expected %s \n", expected)
	// fmt.Printf("actually %v \n", actual)
	return actual == expected
}

func TestFindCommandWithNestedPath() bool {
	println("Testing TestFindCommandWithNestedPath....")
	var expected = "\"SomeString\" : \"This is a string\""
	var jsonString = "{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnotherStringType\" : \"Testing\" }"
	var actual, err = lexicon.ExecuteCommand(lexicon.FIND, ".someObject.SomeString", jsonString)

	if err != nil {
		fmt.Printf("ERROR :: %s", err)
		return false
	}

	expected = strings.ReplaceAll(expected, " ", "")
	actual = strings.ReplaceAll(actual, " ", "")

	//fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

func TestFindCommandWithNestedPathAndArrayIndex() bool {
	println("Testing TestFindCommandWithNestedPathANDArrayIndex....")
	var expected = "\"SomeArray\" : 6"
	var jsonString = "{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnotherStringType\" : \"Testing\" }"
	var actual, err = lexicon.ExecuteCommand(lexicon.FIND, ".someObject.SomeArray[2]", jsonString)

	if err != nil {
		fmt.Printf("ERROR :: %s", err)
		return false
	}

	expected = strings.ReplaceAll(expected, " ", "")
	actual = strings.ReplaceAll(actual, " ", "")

	// fmt.Printf("expected %s \n", expected)
	// fmt.Printf("actually %v \n", actual)
	return actual == expected
}

func TestMarhsallingJSON() bool {
	println("Testing Marshalling...")

	var obj = jsonmarshaller.MarshallJSON("{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnEmptyString\" : \"Testing\" }")

	if len(obj.Properties) > 0 && obj.Properties[0].Property == "someObject" {
		//fmt.Printf("Test Passed!\n")
		return true
	}

	return false
}

func TestUnMarhsallingJSON() bool {
	println("Testing Unmarshalling...")

	var jsonString = "{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnEmptyString\" : \"Testing\" }"
	var obj = jsonmarshaller.MarshallJSON(jsonString)
	var objString, err = jsonmarshaller.Unmarshall(obj)

	if err != nil {
		fmt.Printf("There was an error! %s", err)
		return false
	}

	var expected = strings.ReplaceAll(jsonString, " ", "")
	var actual = strings.ReplaceAll(objString, " ", "")

	//fmt.Printf("E: %s \n\nA: %s \n", expected, actual)
	return strings.EqualFold(actual, expected)
}

func evalTestResult(rst bool) string {
	if rst {
		return "PASSED"
	} else {
		return "FAILED"
	}
}
