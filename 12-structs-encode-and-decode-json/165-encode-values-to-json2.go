package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	users := []user{
		{"zap", "1234", nil},
		{"john", "42", permissions{"admin": true}},
		{"doe", "66", permissions{"write": true}},
	}

	out, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", out)
	fmt.Printf("%#v\n", string(out))
}
