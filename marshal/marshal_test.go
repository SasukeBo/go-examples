package marshal

import (
	"fmt"
	"strings"
	"testing"
)

// User _
type User struct {
	Name string
	Age  int
}

func TestMarshal(t *testing.T) {
	user := User{"sasukebo", 25}

	json, err := Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	result := `{"Name":"sasukebo","Age":25}`
	if strings.TrimSpace(json) != result {
		t.Errorf("json is %s, expect %s", json, result)
	}
}

func TestUnMarshal(t *testing.T) {
	json := `[{"Name":"sasukebo","Age":25}]`

	result, err := Unmarshal(json)
	if err != nil {
		t.Fatal(err)
	}

	users, ok := result.([]interface{})
	if !ok {
		t.Error("result is not a []interface{}")
	}
	fmt.Println(users)

	user, ok := users[0].(map[string]interface{})
	if !ok {
		t.Error("result is not a User")
	}

	if user["Name"] != "sasukebo" {
		t.Errorf("user.Name = %s, expect sasukebo", user["Name"])
	}
}
