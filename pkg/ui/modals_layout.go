package ui

import (
	"github.com/rivo/tview"
)

func (gm *GuiManager) ComposeAndShowModalSignin(buttons []string, msg string) {
	modal := tview.NewModal()
	modal.AddButtons(buttons).SetText(msg).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "back" {
			gm.App.SetRoot(gm.Psignin.Grid, true).
				SetFocus(gm.Psignin.Form.GetFormItem(0).(*tview.InputField))
		}
	})
	gm.App.SetRoot(modal, false).SetFocus(modal)
}

func (gm *GuiManager) ComposeAndShowModalPmangeP(buttons []string, msg string) {
	modal := tview.NewModal()
	modal.AddButtons(buttons).SetText(msg).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "back" {
			gm.App.SetRoot(gm.Pmain.Grid, true)
		}
	})
	gm.App.SetRoot(modal, false).SetFocus(modal)
}

func (gm *GuiManager) ComposeAndShowModalPmangePConfirmation(buttons []string, msg string, action string) {
	modal := tview.NewModal()
	modal.AddButtons(buttons).SetText(msg).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "non" {
			gm.App.SetRoot(gm.Pmain.Grid, true)
			gm.App.SetFocus(gm.Pmain.tableAll)
		} else {
			switch action {
			case "delete":
				e := make(chan string, 1)
				gm.Services.Personne.DeletePersonneById(e)
				err := <-e
				if err != "" {
					gm.ComposeAndShowModalPmangeP([]string{"back"}, err)
				} else {
					gm.ComposeTableGetAll("getAll", false)
					gm.App.SetRoot(gm.Pmain.Grid, true)
					gm.App.SetFocus(gm.Pmain.tableAll)
				}
			case "upadate":

			}

		}
	})
	gm.App.SetRoot(modal, false).SetFocus(modal)
}

func (gm *GuiManager) ComposeAndShowModalMessage(buttons []string, msg string) {
	modal := tview.NewModal()
	modal.AddButtons(buttons).SetText(msg).SetDoneFunc(func(buttonIndex int, buttonLabel string) {

	})
	gm.App.SetRoot(modal, false).SetFocus(modal)
}
