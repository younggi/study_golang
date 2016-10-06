package baby

import "fmt"

func ExampleNameGenerator() {
	bg := NameGenerator("성정명재경", "준호우훈진")
	for s := bg(); s != ""; s = bg() {
		fmt.Print(s, ",")
	}
	// Output:
	// 성준,성호,성우,성훈,성진,정준,정호,정우,정훈,정진,명준,명호,명우,명훈,명진,재준,재호,재우,재훈,재진,경준,경호,경우,경훈,경진,
}

func ExampleCallBack() {
	CallBack("성정명재경", "준호우훈진", func(s string) {
		fmt.Print(s, ",")
	})
	// Output:
	// 성준,성호,성우,성훈,성진,정준,정호,정우,정훈,정진,명준,명호,명우,명훈,명진,재준,재호,재우,재훈,재진,경준,경호,경우,경훈,경진,
}
