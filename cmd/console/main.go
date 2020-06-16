package main

import (
	models "agefice-cons/adil.net/pkg/model"
	"agefice-cons/adil.net/pkg/service"
	"agefice-cons/adil.net/pkg/ui"
	"github.com/rivo/tview"
)

type ApplicationManager struct {
	App *ui.GuiManager
}

func main() {
	tapp := tview.NewApplication()
	app := &ui.GuiManager{App: tapp,
		Psignin:          &ui.Psignin{},
		Pmain:            &ui.Pmain{},
		Fpmanager:        &ui.PmanagePersonne{},
		PmanageFormation: &ui.PmanageFormation{},
		Services:         service.PServicesManagerImpl{Personne: models.Personne{}}}
	apm := &ApplicationManager{App: app}
	// app.ComposeTableGetAll()
	app.ComposePsignin()
	if err := tapp.SetRoot(apm.App.Psignin.Grid, true).
		EnableMouse(true).
		SetFocus(app.Psignin.Form).Run(); err != nil {
		panic(err)
	}
}
