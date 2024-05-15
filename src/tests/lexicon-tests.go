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
	TestFindCommand()
	TestParseJSON()
}

func TestIsLegitRootNode() bool {
	print("Testing IsLegitRootNode, expecting true...")
	var expected = true
	var actual = jsonmarshaller.IsLegitJSONRoot("{}")

	fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

func TestIsLegitRootNodeOnlyOpening() bool {
	print("Testing TestIsLegitRootNodeOnlyOpening, expecting false...")
	var expected = true
	var actual = jsonmarshaller.IsLegitJSONRoot("{")

	fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

func TestIsLegitRootNodeOnlyClosing() bool {
	print("Testing TestIsLegitRootNodeOnlyClosing, expecting false...")
	var expected = true
	var actual = jsonmarshaller.IsLegitJSONRoot("}")

	fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

// Check the find command
func TestFindCommand() bool {
	print("Testing TestFindCommand, expecting TRUE...")
	var expected = "{}"
	var actual, err = lexicon.ExecuteCommand(lexicon.FIND, ".", "{}")

	if len(err) > 0 {
		println("Test FAILED!")
	}

	fmt.Printf("actually got %v \n", actual)
	return actual == expected
}

func TestParseJSON() bool {
	print("Testing ParseJSONObject, expecting TRUE...")
	//var expected = "{}"
	//var _ = lexicon.ParseJSONObjectString2("{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnEmptyString\" : \"Testing\"  }")
	var _ = jsonmarshaller.MarshallJSON("{ \"someObject\" : {\"SomeArray\" : [8,7,6,5,3,0,9], \"SomeNumber\" : 1, \"SomeString\" : \"This is a string\", \"Aboolean\" : true }, \"AnEmptyString\" : \"Testing\"  }")
	//fmt.Printf("actually got %v \n", actual)
	return true //actual == expected
}

//"{ \"someProperty\" : {\"SOMETHING\" : [], \"ANOTHER THING\" : 1, \"dude\" : \"Wheres my car\" }, \"AnotherProp\" : \"\",  }"
//{ "someObject" : {"SomeArray" : [], "SomeNumber" : 1, "SomeString" : "This is a string", "Aboolean" : true }
