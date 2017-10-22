package golang

import (
	"encoding/json"
	"encoding/xml"
	"testing"
)

type langStruct struct{}

func TestJsonStructure(t *testing.T) {
	jsonStr := `{
	"id": 1,
	"lang": "go"		
}`

	type testStruct = langStruct
	var myStruct testStruct

	err := json.Unmarshal([]byte(jsonStr), &myStruct)

	if err != nil {
		t.Error("it should not error when you make the righ struct that represent json string")
	}

	if myStruct.ID != 1 {
		t.Error("it should be 1")
	}

	if myStruct.Lang != "go" {
		t.Error("it should be go")
	}
}

func TestJsonStructures(t *testing.T) {
	jsonStr := `[
	{
		"id": 1,
		"lang": "go"
	},
	{
		"id": 2,
		"lang": "java"
	},
	{
		"id": 3,
		"lang": "python"
	}
]`

	type testStruct = langStruct
	var myStruct []testStruct

	err := json.Unmarshal([]byte(jsonStr), &myStruct)

	if err != nil {
		t.Error("it should not error when you make the righ struct that represent json string")
	}
}

func TestJsonAndXMLStructure(t *testing.T) {
	xmlStr := `<xml>
	<id>1</id>
	<lang>go</lang>
	</xml>
}`

	type testStruct = langStruct
	var myStruct testStruct

	err := xml.Unmarshal([]byte(xmlStr), &myStruct)

	if err != nil {
		t.Error("it should not error when you make the righ struct that represent json string")
	}

}
