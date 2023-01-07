package main

type GUIFactory interface {
	CreateButton() (Button, error)
	CreateDialog() (Dialog, error)
}

type WinFactory struct {
}

func (w WinFactory) CreateButton() (Button, error) {
	return WinButton{}, nil
}

func (w WinFactory) CreateDialog() (Dialog, error) {
	return WinDialog{}, nil
}

type MacFactory struct {
}

func (w MacFactory) CreateButton() (Button, error) {
	return MacButton{}, nil
}

func (w MacFactory) CreateDialog() (Dialog, error) {
	return MacDialog{}, nil
}
