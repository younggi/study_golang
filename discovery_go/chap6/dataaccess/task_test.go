package dataaccess

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func Example_marshalJSON() {
	t := Task{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2015, time.August, 16, 15, 43, 0, 0, time.UTC)),
		0,
		nil,
	}
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output:
	// {"title":"Laundry","status":"DONE","deadline":"2015-08-16T15:43:00Z"}
}

func Example_unmarshalJSON() {
	b := []byte(`{"Title":"Buy Milk","Status":"DONE","Deadline":"2015-08-16T15:43:00Z"}`)
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
	// Output:
	// Buy Milk
	// DONE
	// 2015-08-16 15:43:00 +0000 UTC
}
