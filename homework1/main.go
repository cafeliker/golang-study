package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Users struct {
	Users []User `json:"users"`
}

func main() {
	content, err := ioutil.ReadFile("./users.json")
	if err != nil {
		panic("can't open file: " + err.Error())
	}

	users := Users{}
	err = json.Unmarshal(content, &users)
	if err != nil {
		panic("can't unmarshal the content: " + err.Error())
	}

	for _, user := range users.Users {
		fmt.Println("User Type: ", user.Type)
		fmt.Println("User Age: ", user.Age)
		fmt.Println("User Name: ", user.Name)
		fmt.Println("Facebook Url: ", user.Social.Facebook)
	}

	fmt.Println(users)
}
