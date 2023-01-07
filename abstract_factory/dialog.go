package main

import "fmt"

type Dialog interface {
	PositiveButtonClick() error
	NegativeButtonClick() error
}

type WinDialog struct{}

func (d WinDialog) PositiveButtonClick() error {
	fmt.Println("Positive button was clicked on the windows dialog")
	return nil
}

func (d WinDialog) NegativeButtonClick() error {
	fmt.Println("Negative button was clicked on the windows dialog")
	return nil
}

type MacDialog struct{}

func (d MacDialog) PositiveButtonClick() error {
	fmt.Println("Positive button was clicked on the mac dialog")
	return nil
}

func (d MacDialog) NegativeButtonClick() error {
	fmt.Println("Negative button was clicked on the mac dialog")
	return nil
}
