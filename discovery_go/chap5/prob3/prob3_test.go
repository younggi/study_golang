package prob3

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type ID int64

type User struct {
	Name string
	ID   ID
}

func (i *ID) UnmarshalJSON(data []byte) error {
	var f interface{}
	err := json.Unmarshal(data, &f)
	if err != nil {
		return err
	}
	switch v := f.(type) {
	case string:
		tmp, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*i = ID(tmp)
	case float64:
		*i = ID(v)
	default:
		return errors.New("Invalid Type")
	}
	return nil
}

func ExampleUserJSON() {
	b := []byte(`{"Name":"Younggi","ID":12345}`)
	u := User{}
	err := json.Unmarshal(b, &u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(u.Name)
	fmt.Println(u.ID)
	// Output:
	// Younggi
	// 12345
}

func ExampleUserJSON_Old() {
	b := []byte(`{"Name":"Younggi","ID":"12345"}`)
	u := User{}
	err := json.Unmarshal(b, &u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(u.Name)
	fmt.Println(u.ID)
	// Output:
	// Younggi
	// 12345
}
