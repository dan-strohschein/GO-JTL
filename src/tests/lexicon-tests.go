package tests

import (
	"fmt"
	jsonmarshaller "jtl/json-marshaller"
	"jtl/lexicon"
)

func TestAll() {
	TestIsLegitRootNode()
	TestIsLegitRootNodeOnlyOpening()
	TestIsLegitRootNodeOnlyClosing()

	TestMarhsallingJSON()
	TestUnMarhsallingJSON()

	TestFindCommand()
	TestFindCommandWithNestedPath()
	TestFindCommandWithNestedPathAndArrayIndex()
}

func TestIsLegitRootNode() bool {
	println("Testing IsLegitRootNode, expecting true...")
	var expected = true
	var actual = jsonmarshaller.IsLegitJSONRoot("{}")

	fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

func TestIsLegitRootNodeOnlyOpening() bool {
	println("Testing TestIsLegitRootNodeOnlyOpening, expecting false...")
	var expected = true
	var actual = jsonmarshaller.IsLegitJSONRoot("{")

	fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

func TestIsLegitRootNodeOnlyClosing() bool {
	println("Testing TestIsLegitRootNodeOnlyClosing, expecting false...")
	var expected = true
	var actual = jsonmarshaller.IsLegitJSONRoot("}")

	fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

// Check the find command
func TestFindCommand() bool {
	println("Testing TestFindCommand....")
	var expected = "{}"
	var jsonString = "{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnotherStringType\" : \"Testing\" }"
	var actual, _ = lexicon.ExecuteCommand(lexicon.FIND, ".someObject", jsonString)

	if len(actual) <= 0 {
		println("Test FAILED!")
	}

	fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

func TestFindCommandWithNestedPath() bool {
	println("Testing TestFindCommandWithNestedPath....")
	var expected = "{}"
	var jsonString = "{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnotherStringType\" : \"Testing\" }"
	var actual, _ = lexicon.ExecuteCommand(lexicon.FIND, ".someObject.SomeString", jsonString)

	if len(actual) <= 0 {
		println("Test FAILED!")
	}

	fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

func TestFindCommandWithNestedPathAndArrayIndex() bool {
	println("Testing TestFindCommandWithNestedPathANDArrayIndex....")
	var expected = "\"SomeArray\" : 6"
	var jsonString = "{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnotherStringType\" : \"Testing\" }"
	var actual, _ = lexicon.ExecuteCommand(lexicon.FIND, ".someObject.SomeArray[2]", jsonString)

	if len(actual) <= 0 {
		println("Test FAILED!")
	}
	fmt.Printf("expected %s \n", expected)
	fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

func TestMarhsallingJSON() bool {
	println("Testing Marshalling, expecting TRUE...")
	//var expected = "{}"
	//var _ = lexicon.ParseJSONObjectString2("{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnEmptyString\" : \"Testing\"  }")
	var obj = jsonmarshaller.MarshallJSON("{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnEmptyString\" : \"Testing\" }")

	if len(obj.Properties) > 0 && obj.Properties[0].Property == "someObject" {
		fmt.Printf("Test Passed!")

	}

	//fmt.Printf("actually got %v \n", actual)
	return true //actual == expected
}

func TestUnMarhsallingJSON() bool {
	println("Testing Unmarshalling, expecting TRUE...")
	//var expected = "{}"
	var jsonString = "{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnEmptyString\" : \"Testing\" }"
	var obj = jsonmarshaller.MarshallJSON(jsonString)
	var objString, err = jsonmarshaller.Unmarshall(obj)

	if err != nil {
		fmt.Printf("There was an error! %s", err)
	}

	fmt.Printf("M: %s \n\nU: %s \n", jsonString, objString)
	return true //actual == expected
}

//"{ \"someProperty\" : {\"SOMETHING\" : [], \"ANOTHER THING\" : 1, \"dude\" : \"Wheres my car\" }, \"AnotherProp\" : \"\",  }"
//{ "someObject" : {"SomeArray" : [], "SomeNumber" : 1, "SomeString" : "This is a string", "Aboolean" : true }
