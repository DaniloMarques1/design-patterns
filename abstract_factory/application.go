package main

type Application struct {
	factory GUIFactory
}

// will receive the winfactory or the macfactory in here
func NewApplication(factory GUIFactory) *Application {
	return &Application{factory: factory}
}

func (a *Application) CreateUI() error {
	button, err := a.factory.CreateButton()
	if err != nil {
		return err
	}
	dialog, err := a.factory.CreateDialog()
	if err != nil {
		return err
	}

	if err := button.OnClick(); err != nil {
		return err
	}

	if err := dialog.PositiveButtonClick(); err != nil {
		return err
	}

	return nil
}
