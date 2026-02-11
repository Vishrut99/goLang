package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// All datatypes with JSON tags
type Data struct {
	Int       int            `json:"int"`
	Float     float64        `json:"float"`
	String    string         `json:"string"`
	Bool      bool           `json:"bool"`
	Slice     []int          `json:"slice"`
	Array     [2]string      `json:"array"`
	Map       map[string]int `json:"map"`
	Interface interface{}    `json:"interface"`
	Pointer   *int           `json:"pointer"`
	Ignore    string         `json:"-"`
	Zero      string         `json:"zero,omitempty"`
}

func main() {
	// MARSHAL - Object to JSON
	obj := Data{
		Int:       42,
		Float:     3.14,
		String:    "hello",
		Bool:      true,
		Slice:     []int{1, 2, 3},
		Array:     [2]string{"a", "b"},
		Map:       map[string]int{"key": 100},
		Interface: "any value",
		Pointer:   intPtr(99),
		Ignore:    "ignored",
		Zero:      "",
	}

	fileData, err := os.ReadFile("read.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var file Data
	err = json.Unmarshal(fileData, &file)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Printf("Printing from file \n%v\n", file)


	_ = fileData // Use fileData if needed
	jsonBytes, _ := json.Marshal(obj)
	fmt.Println("MARSHAL:\n", string(jsonBytes))

	// UNMARSHAL - JSON to Object
	jsonStr := `{
		"int":42,
		"float":3.14,
		"string":"hello",
		"bool":true,
		"slice":[1,2,3],
		"array":["a","b"],
		"map":{"key":100},
		"interface":"any value",
		"pointer":99
	}`

	var result Data
	json.Unmarshal([]byte(jsonStr), &result)
	fmt.Println("\nUNMARSHAL:")
	fmt.Printf("%+v\n\n", result)

	fmt.Println("////////////////////////////////////////////////////////////////////////")

	//  Zero values
	empty := Data{}
	j, _ := json.Marshal(empty)
	fmt.Println(" Zero values:", string(j))

	//  Null pointer
	var ptr *int
	obj2 := Data{Pointer: ptr}
	j, _ = json.Marshal(obj2)
	fmt.Printf(" Nil pointer: %+v\n", string(j))

	//  Extra fields in JSON ignored
	var obj4 Data
	json.Unmarshal([]byte(`{"int":5,"unknown":"ignored"}`), &obj4)
	fmt.Printf(" Extra fields ignored: %+v\n", obj4)

	//  Wrong type fails silently (keeps zero value)
	var obj5 Data
	json.Unmarshal([]byte(`{"int":"not-a-number"}`), &obj5)
	fmt.Printf(" Wrong type: %+v\n", obj5)

	//  Interface{} accepts any type
	var obj6 Data
	json.Unmarshal([]byte(`{"interface":123}`), &obj6)
	fmt.Printf(" Interface accepts int: %+v\n", obj6)
}

func intPtr(i int) *int {
	return &i
}
