package ui

import (
	"github.com/rivo/tview"
)

type Psignin struct {
	Form *tview.Form
	Grid *tview.Grid
}

func (gm *GuiManager) initPsignin() {
	gm.Psignin.Form = tview.NewForm()
	gm.Psignin.Grid = tview.NewGrid()
}

func (gm *GuiManager) ComposePsignin() {
	gm.initPsignin()
	gm.composeGrid()
}

func (gm *GuiManager) composeGrid() {
	gm.Psignin.Grid.SetColumns(-1, 30, -1)
	gm.Psignin.Grid.SetRows(-1, 7, -1)
	gm.composeForm()
	gm.Psignin.Grid.AddItem(gm.Psignin.Form, 1, 1, 1, 1, 0, 0, false)
}

func (gm *GuiManager) composeForm() {

	gm.Psignin.Form.AddInputField("login", "", 20, nil, nil).SetFocus(0).
		AddPasswordField("password", "", 20, 0, nil).
		AddButton("sign in", func() {
			buttons := []string{"back"}
			if !gm.Utils.CheckAuthentification(gm.Psignin.Form.GetFormItem(0).(*tview.InputField).GetText(),
				gm.Psignin.Form.GetFormItem(1).(*tview.InputField).GetText()) {
				gm.ComposeAndShowModalSignin(buttons, "login or password incorrect")
			} else {
				//gm.ComposeAndShowModalSignin(buttons,"lauthentificated :)")
				gm.ComposeAndShowPmain()

			}
		})
}
