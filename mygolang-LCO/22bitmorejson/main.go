package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"cousename"` //creating alias
	Price    int
	Platform string   `json:"website"`
	Pass     string   `json:"-"`
	Tag      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Bit more json")
	//encodeJson()
	decodeJsaonFromWeb()
}

func encodeJson() {
	lcoCourses := []course{
		{"Java Script", 12, "lco", "AbS", []string{"web development", "front end"}},
		{"CSS", 129, "lco", "hit123", []string{"web dev", "front end after html"}},
		{"HTML", 299, "lco", "bcd345", nil},
	}
	// create json from avil data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func decodeJsaonFromWeb() {
	jsonResponsefromWeb := []byte(`
  {
	"cousename": "Java Script",
	"Price": 12,
	"website": "lco",
	"tags": ["web development","front end"]
  }
`)

	var lcoCourse course

	checkValid := json.Valid(jsonResponsefromWeb)

	if checkValid == true {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonResponsefromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}
	// Use of interface method to consume json data
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonResponsefromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("For every Key %v, Value is %v and Datatype of Value is %T\n", k, v, v)
	}
}
