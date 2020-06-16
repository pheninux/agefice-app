package ui

import (
	"github.com/rivo/tview"
)

type PmanageFormation struct {
	Page       *tview.Pages
	Grid       *tview.Grid
	Fformation *tview.Form
}

func (gm *GuiManager) PmfInit() {
	gm.PmanageFormation.Page = tview.NewPages()
	gm.PmanageFormation.Grid = tview.NewGrid()
	gm.PmanageFormation.Fformation = tview.NewForm()

}

func (gm *GuiManager) ComposeAndShowPmanageFormation() {

	gm.PmfInit()
	gm.ComposeFormFormation()
	gm.ComposeGrid()
	gm.ComposePage()

}

func (gm *GuiManager) ComposeFormFormation() {

	gm.PmanageFormation.Fformation.AddInputField("Formation", "", 30, nil, nil).
		AddInputField("Date début", "", 10, nil, nil).
		AddInputField("Date fin", "", 10, nil, nil).
		AddInputField("Nombre H", "", 10, tview.InputFieldInteger, nil).
		AddInputField("Coût", "", 10, tview.InputFieldInteger, nil).
		AddButton("Enregistrer", func() {

		}).AddButton("Back", func() {

	})
	gm.PmanageFormation.Fformation.SetBorder(true).SetTitle("Formulaire formation")

}
func (gm *GuiManager) ComposeGrid() {

	gm.PmanageFormation.Grid.SetRows(-1, 15, -1)
	gm.PmanageFormation.Grid.SetColumns(-1, 30, -1)
	gm.PmanageFormation.Grid.AddItem(gm.PmanageFormation.Fformation, 1, 1, 1, 1, 0, 0, false)

}

func (gm *GuiManager) ComposePage() {
	gm.Pmain.Pages.AddPage("page_manage_formation", gm.PmanageFormation.Grid, true, false)
}
