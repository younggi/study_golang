package task

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

func (d *Deadline) OverDue() bool {
	return d != nil && d.Before(time.Now())
}

type status int

const (
	UNKNOWN status = iota // 0
	TODO                  // 1
	DONE                  // 2
)

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

func (t Task) OverDue() bool {
	return t.Deadline.OverDue()
}

func ExampleDeadline_OverDue() {
	d1 := Deadline{time.Now().Add(-4 * time.Hour)}
	d2 := Deadline{time.Now().Add(4 * time.Hour)}
	fmt.Println(d1.OverDue())
	fmt.Println(d2.OverDue())
	// Output:
	// true
	// false
}

func Example_taskTestAll() {
	d1 := Deadline{time.Now().Add(-4 * time.Hour)}
	d2 := Deadline{time.Now().Add(4 * time.Hour)}
	t1 := Task{"4h ago", TODO, &d1, 0, nil}
	t2 := Task{"4h later", TODO, &d2, 0, nil}
	t3 := Task{"no due", TODO, nil, 0, nil}
	fmt.Println(t1.OverDue())
	fmt.Println(t2.OverDue())
	fmt.Println(t3.OverDue())
	// Output:
	// true
	// false
	// false
}

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
	// {"title":"Laundry","status":2,"deadline":1439739780}
}

func Example_unmarshalJSON() {
	b := []byte(`{"Title":"Buy Milk","Status":2,"Deadline":1439739780}`)
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
	// 2
	// 2015-08-16 15:43:00 +0000 UTC
}

func (d Deadline) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, d.Unix(), 10), nil
}

func (d *Deadline) UnmarshalJSON(data []byte) error {
	unix, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.Unix(unix, 0)
	return nil
}

func Example_mapMarshalJSON() {
	b, _ := json.Marshal(map[string]string{
		"Name": "John",
		"Age":  "16",
	})
	fmt.Println(string(b))
	// Output:
	// {"Age":"16","Name":"John"}
}

func Example_mapMarshalJSON2() {
	b, _ := json.Marshal(map[string]interface{}{
		"Name": "John",
		"Age":  16,
	})
	fmt.Println(string(b))
	// Output:
	// {"Age":16,"Name":"John"}
}

type Fields struct {
	VisibleField   string
	InvisibleField string
}

func ExampleOmitFields() {
	f := &Fields{"a", "b"}
	b, _ := json.Marshal(struct {
		*Fields
		InvisibleField string `json:",omitempty"`
		Additional     string
	}{Fields: f, Additional: "c"})
	fmt.Println(string(b))
	// Output:
	// {"VisibleField":"a","Additional":"c"}
}

func Example_gob() {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	data := map[string]string{"N": "J"}
	if err := enc.Encode(data); err != nil {
		fmt.Println(err)
	}

	const width = 16
	for start := 0; start < len(b.Bytes()); start += width {
		end := start + width
		if end > len(b.Bytes()) {
			end = len(b.Bytes())
		}
		fmt.Printf("% x\n", b.Bytes()[start:end])
	}
	dec := gob.NewDecoder(&b)
	var restored map[string]string
	if err := dec.Decode(&restored); err != nil {
		fmt.Println(err)
	}
	fmt.Println(restored)
	// Output:
	// 0e ff 81 04 01 02 ff 82 00 01 0c 01 0c 00 00 08
	// ff 82 00 01 01 4e 01 4a
	// map[N:J]
}

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

func ExampleTask_String() {
	fmt.Println(Task{"Laundry", DONE, nil, 0, nil})
	// Output:
	// [v] Laundry <nil>
}

type IncludeSubTasks Task

func (t IncludeSubTasks) indentedString(prefix string) string {
	str := prefix + Task(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + IncludeSubTasks(st).indentedString(prefix+"  ")
	}
	return str
}

func (t IncludeSubTasks) String() string {
	return t.indentedString("")
}

func ExampleIncludeSubTasks_String() {
	fmt.Println(IncludeSubTasks(Task{
		Title:    "Laundry",
		Status:   TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task{{
			Title:    "Wash",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: []Task{
				{"Put", DONE, nil, 2, nil},
				{"Detergent", TODO, nil, 2, nil},
			},
		}, {
			Title:    "Dry",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}, {
			Title:    "Fold",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	}))
	// Output:
	// [ ] Laundry <nil>
	//   [ ] Wash <nil>
	//     [v] Put <nil>
	//     [ ] Detergent <nil>
	//   [ ] Dry <nil>
	//   [ ] Fold <nil>
}

func Join(sep string, a ...interface{}) string {
	if len(a) == 0 {
		return ""
	}
	t := make([]string, len(a))
	for i := range a {
		switch x := a[i].(type) {
		case string:
			t[i] = x
		case int:
			t[i] = strconv.Itoa(x)
		case fmt.Stringer:
			t[i] = x.String()
		}
	}
	return strings.Join(t, sep)
}

func ExampleJoin() {
	t := Task{
		Title:    "Laundry",
		Status:   DONE,
		Deadline: nil,
	}
	fmt.Println(Join(",", 1, "two", 3, t))
	// Output:
	// 1,two,3,[v] Laundry <nil>
}

func (t *Task) MarkDone() {
	t.Status = DONE
	for i := 0; i < len(t.SubTasks); i++ {
		t.SubTasks[i].MarkDone()
	}
}

func ExampleMarkDone() {
	t := &(Task{
		Title:    "Laundry",
		Status:   TODO,
		Deadline: nil,
		Priority: 2,
		SubTasks: []Task{{
			Title:    "Wash",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: []Task{
				{"Put", DONE, nil, 2, nil},
				{"Detergent", TODO, nil, 2, nil},
			},
		}, {
			Title:    "Dry",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}, {
			Title:    "Fold",
			Status:   TODO,
			Deadline: nil,
			Priority: 2,
			SubTasks: nil,
		}},
	})
	t.MarkDone()
	fmt.Println(IncludeSubTasks(*t))
	// Output:
	// [v] Laundry <nil>
	//   [v] Wash <nil>
	//     [v] Put <nil>
	//     [v] Detergent <nil>
	//   [v] Dry <nil>
	//   [v] Fold <nil>
}
