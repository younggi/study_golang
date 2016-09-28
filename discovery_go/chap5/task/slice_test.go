package task

import "fmt"

type Task2 struct {
	Value   int
	SubTask []Task2
}

func (t Task2) String() string {
	str := fmt.Sprintf("[%d]\n", t.Value)
	for i := 0; i < len(t.SubTask); i++ {
		str += "-" + t.SubTask[i].String()
	}
	return str
}

func (t *Task2) Reset() {
	t.Value = 0
	for i := 0; i < len(t.SubTask); i++ {
		t.SubTask[i].Reset()
	}
}

func ExampleTask2String() {
	t := Task2{
		Value: 1,
		SubTask: []Task2{
			{2, nil},
			{3, nil},
		},
	}
	t.Reset()
	fmt.Println(t)
	// Output:
	// [0]
	// -[0]
	// -[0]
}
