package sample

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Response struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func UnMarshal() {
	js := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response{}
	json.Unmarshal([]byte(js), &res)
	fmt.Printf("%v, %v\n", res.Page, res.Fruits)
}

func UnstructuredData() {
	js := `{"page": 1, "fruits": ["apple", "peach"]}`
	var result map[string]interface{}
	json.Unmarshal([]byte(js), &result)

	for key, value := range result {
		vt := reflect.TypeOf(value)
		fmt.Printf("%v, %v, %v\n", key, value, vt.Kind())
	}
}

func TA() {
	UnstructuredData()
}
