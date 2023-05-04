package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool // #3

type user struct { // #1
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"` // #6

	// name        string // #1
	// password    string // #1
	// permissions // #3
}

func main() {
	users := []user{ // #2
		{
			Name:        "zap",
			Password:    "1234",
			Permissions: nil,
		},
		{
			Name:        "john",
			Password:    "42",
			Permissions: permissions{"admin": true},
		},
		{
			Name:        "doe",
			Password:    "66",
			Permissions: permissions{"write": true},
		},
	}

	// out, err := json.Marshal(users) // #4
	out, err := json.MarshalIndent(users, "", "\t") // #5
	if err != nil {                                 // #4
		fmt.Println(err)
		return
	}

	fmt.Println(string(out)) // #4
}
