package main

import "fmt"

type Button interface {
	OnClick() error
}

type WinButton struct{}

func (w WinButton) OnClick() error {
	fmt.Println("OnClick windows button")
	return nil
}

type MacButton struct{}

func (m MacButton) OnClick() error {
	fmt.Println("OnClick mac button")
	return nil
}
