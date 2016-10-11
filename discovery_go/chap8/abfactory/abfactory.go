package main

import "fmt"

// Button is interface of button component
type Button interface {
	Paint()
	OnClick()
}

// Label is interface of painting label
type Label interface {
	Paint()
}

// WinButton is a Button implementation for Windows
type WinButton struct{}

// Paint is a method implementation of Button interface
func (WinButton) Paint() { fmt.Println("win button paint") }

// OnClick is a method implementation of Button interface
func (WinButton) OnClick() { fmt.Println("win button click") }

// WinLabel is a Label implementation for Windows.
type WinLabel struct{}

// Paint is a method implementation of Label interface
func (WinLabel) Paint() { fmt.Println("win label paint") }

// MacButton is a Button implementation for Mac.
type MacButton struct{}

// Paint a method implementation of Mac Button interface
func (MacButton) Paint() { fmt.Println("mac button paint") }

// OnClick is a method implementation of Mac Button interface
func (MacButton) OnClick() { fmt.Println("mac button click") }

// MacLabel is a label implementation for Mac.
type MacLabel struct{}

// Paint is a method implementation of Mac Label interface
func (MacLabel) Paint() { fmt.Println("mac label paint") }

// UIFactory can create buttons and labels.
type UIFactory interface {
	CreateButton() Button
	CreateLabel() Label
}

// WinFactory is a UI factory that can create Windows UI elements.
type WinFactory struct{}

// CreateButton is a method implementation of WinFactory
func (WinFactory) CreateButton() Button {
	return WinButton{}
}

// CreateLabel is a method implementation of WinFactory
func (WinFactory) CreateLabel() Label {
	return WinLabel{}
}

// MacFactory is a UI factory that can create Mac UI elements
type MacFactory struct{}

// CreateButton is a method implementation of MacFactory
func (MacFactory) CreateButton() Button {
	return MacButton{}
}

// CreateLabel is a method implementation of MacFactory
func (MacFactory) CreateLabel() Label {
	return MacLabel{}
}

// CreateFactory returns a UFactory of the given os.
func CreateFactory(os string) UIFactory {
	if os == "win" {
		return WinFactory{}
	}
	return MacFactory{}
}

// Run create Button and Label, run it
func Run(f UIFactory) {
	button := f.CreateButton()
	button.Paint()
	button.OnClick()
	label := f.CreateLabel()
	label.Paint()
}

func main() {
	f1 := CreateFactory("win")
	Run(f1)
	f2 := CreateFactory("mac")
	Run(f2)
}
