package ui

import (
	"agefice-cons/adil.net/pkg/service"
	"agefice-cons/adil.net/pkg/utils"
	"github.com/rivo/tview"
)

type GuiManager struct {
	App              *tview.Application
	Services         service.PServicesManagerImpl
	Psignin          *Psignin
	Pmain            *Pmain
	Fpmanager        *PmanagePersonne
	PmanageFormation *PmanageFormation
	Utils            *utils.Utils
}
