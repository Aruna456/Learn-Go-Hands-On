package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json")
	// encodeJson()
	DecodeJson()

}

func encodeJson() {

	tekiiCourses := []course{
		{"Java", 299, "LearnJavaWithTekki.in", "12345", []string{"programming", "java"}},
		{"ReactJs", 299, "LearnReactWithTekki.in", "12345", []string{"Web-dev", "js"}},
		{"Go", 299, "LearnGoWithTekki.in", "12345", nil}}

	//package this data into Json data

	finalJson, err := json.MarshalIndent(tekiiCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {

	jsonDatafromWeb := []byte(`{
                "courseName": "Java",
                "Price": 299,
                "website": "LearnJavaWithTekki.in",
                "tags": ["programming","java"]
        }`)

	var tekkicourse course

	validity := json.Valid(jsonDatafromWeb)

	if validity {
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonDatafromWeb, &tekkicourse)
		fmt.Printf("The parsed Json is %#v\n", tekkicourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	var myOnlineData map[string]interface{} //can use any if value is of any type
	json.Unmarshal(jsonDatafromWeb, &myOnlineData)
	// fmt.Printf("%#v", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and Value is %v and Type of the data is %T\n", k, v, v)
	}

}
