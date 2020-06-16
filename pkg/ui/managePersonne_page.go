package ui

import (
	models "agefice-cons/adil.net/pkg/model"
	"fmt"
	"github.com/badoux/checkmail"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strconv"
	"time"
)

type PmanagePersonne struct {
	Fperonne    *tview.Form
	Fformation  *tview.Form
	Fdocument   *tview.Form
	Fentreprise *tview.Form
	Flex        *tview.Flex
	Flag        bool
}

// map pour rassembler les id et intitulé de la formations
var mIdf map[string]int

// map pour rassembler les id et intitulé des documents
var mIdd map[string]int

func (gm *GuiManager) initPmPersonne() {
	gm.Fpmanager.Fperonne = tview.NewForm()
	gm.Fpmanager.Fdocument = tview.NewForm()
	gm.Fpmanager.Fformation = tview.NewForm()
	gm.Fpmanager.Fentreprise = tview.NewForm()
	gm.Fpmanager.Flex = tview.NewFlex()
	gm.Fpmanager.Flag = true
}

func (gm *GuiManager) ComposeAndShowPmanagePersonne() {

	gm.initPmPersonne()
	gm.composeFormManagePersonne()
	gm.composeFormManageEntreprise()
	gm.composeFormManageFormation()
	gm.composeFormManageDocument()
	gm.composeActions()
	gm.Pmain.Pages.AddPage("manage-stagiaire-page", gm.Fpmanager.Flex, true, false)

}

func (gm *GuiManager) composeFormManagePersonne() {

	gm.Fpmanager.Fperonne.
		AddCheckbox("MFA", false, nil).
		AddInputField("Nom", "", 30, nil, nil).
		AddInputField("Prenom", "", 30, nil, nil).
		AddInputField("age", "", 2, tview.InputFieldInteger, nil).
		AddInputField("Date naissance", "", 10, nil, nil).
		AddInputField("Tel", "", 10, tview.InputFieldInteger, nil).
		AddInputField("Mail", "", 30, nil, nil).
		AddInputField("Adresse", "", 30, nil, nil)
	gm.Fpmanager.Fperonne.SetBorder(true).SetTitle("Formulaire stagiaire")

	for i := 1; i < gm.Fpmanager.Fperonne.GetFormItemCount(); i++ {
		input := gm.Fpmanager.Fperonne.GetFormItem(i).(*tview.InputField)
		input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyCtrlA {
				gm.App.SetFocus(gm.Fpmanager.Fentreprise.GetFormItem(0))
			}
			return event
		})

		if i == 6 {
			input.SetDoneFunc(func(key tcell.Key) {
				err := checkmail.ValidateFormat(gm.Fpmanager.Fperonne.GetFormItem(6).(*tview.InputField).GetText())
				if err != nil {
					gm.ComposeAndShowModalPmangeP([]string{"back"}, err.Error())
					return
				}

				err = checkmail.ValidateHost(gm.Fpmanager.Fperonne.GetFormItem(6).(*tview.InputField).GetText())
				if err != nil {
					gm.ComposeAndShowModalPmangeP([]string{"back"}, err.Error())
					return
				}

				err = checkmail.ValidateHost(gm.Fpmanager.Fperonne.GetFormItem(6).(*tview.InputField).GetText())
				if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
					gm.ComposeAndShowModalPmangeP([]string{"back"}, fmt.Sprintf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr))
					return
				}
			})
		}

		if i == 4 {
			input.SetDoneFunc(func(key tcell.Key) {
				age, _ := strconv.Atoi(gm.Fpmanager.Fperonne.GetFormItem(3).(*tview.InputField).GetText())
				t, err := time.Parse("2006-01-02", gm.Fpmanager.Fperonne.GetFormItem(4).(*tview.InputField).GetText())
				if err != nil {
					gm.ComposeAndShowModalPmangeP([]string{"back"}, err.Error())
				}
				years := int(time.Since(t).Hours() / 8760)
				if years != age {
					gm.ComposeAndShowModalPmangeP([]string{"back"}, "Date naissance n'est pas compatible avec l'age")
				}
			})
		}

	}
	/* gm.Fpmanager.Fperonne.GetFormItem(0).(*tview.InputField).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlA {
			gm.App.SetFocus(gm.Fpmanager.Fentreprise.GetFormItem(0))
		}
		return event
	})*/
}

func (gm *GuiManager) composeFormManageDocument() {
	mIdd = map[string]int{}
	c := make(chan []models.Document, 1)
	e := make(chan error, 1)
	go gm.Services.Docuement.GetDocuments(c, e)
	docs := <-c
	err := <-e
	if err != nil {
		gm.ComposeAndShowModalPmangeP([]string{"back"}, err.Error())
	}
	mIdd = make(map[string]int)
	if len(docs) > 0 {
		for _, v := range docs {
			mIdd[v.Libelle] = v.Id
			gm.Fpmanager.Fdocument.AddCheckbox(v.Libelle, false, nil)
		}
	}
	gm.Fpmanager.Fdocument.SetBorder(true).SetTitle("Formulaire Documents")

	for i := 0; i < gm.Fpmanager.Fdocument.GetFormItemCount(); i++ {
		check := gm.Fpmanager.Fdocument.GetFormItem(i).(*tview.Checkbox)
		check.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyCtrlA {
				gm.App.SetFocus(gm.Fpmanager.Fperonne.GetFormItem(0))
			}
			return event
		})

	}

	/*gm.Fpmanager.Fdocument.GetFormItem(0).(*tview.Checkbox).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlA {
			gm.App.SetFocus(gm.Fpmanager.Fperonne.GetFormItem(0))
		}
		return event
	})*/
}

func (gm *GuiManager) composeFormManageFormation() {
	mIdf = map[string]int{}
	c := make(chan []models.Formation, 1)
	e := make(chan error, 1)
	go gm.Services.Formation.GetFormations(c, e)
	err := <-e
	if err != nil {
		gm.ComposeAndShowModalPmangeP([]string{"back"}, err.Error())
	}
	frms := <-c
	var lf []string
	if len(frms) > 0 {
		for _, v := range frms {
			lf = append(lf, v.Intitule)
			mIdf[v.Intitule] = v.Id
		}
	}
	dd := tview.NewDropDown()
	dd.SetLabel("Intitulé")
	dd.SetOptions(lf, nil)
	dd.SetFieldWidth(30)
	dd.SetTextOptions("", "", "", "", "")
	gm.Fpmanager.Fformation.AddFormItem(dd).
		AddInputField("", "", 30, nil, nil).
		AddInputField("Date début", "", 10, nil, nil).
		AddInputField("Date fin", "", 10, nil, nil).
		AddInputField("Nombre H", "", 10, tview.InputFieldInteger, nil).
		AddInputField("Coût", "", 10, tview.InputFieldInteger, nil)

	gm.Fpmanager.Fformation.SetBorder(true).SetTitle("Formulaire Formation")

	gm.Fpmanager.Fformation.GetFormItem(0).(*tview.DropDown).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlA {
			gm.App.SetFocus(gm.Fpmanager.Fdocument.GetFormItem(0))
		}
		return event
	})

	gm.Fpmanager.Fformation.GetFormItem(3).(*tview.InputField).SetDoneFunc(func(key tcell.Key) {

		dateDeb, _ := time.Parse("2006-01-02", gm.Fpmanager.Fformation.GetFormItem(2).(*tview.InputField).GetText())
		dateFin, err := time.Parse("2006-01-02", gm.Fpmanager.Fformation.GetFormItem(3).(*tview.InputField).GetText())
		if err != nil {
			gm.ComposeAndShowModalPmangeP([]string{"back"}, "La date saisie est incorrect . saissisez la date comme ce format (2006-01-02)")
		}

		if int(time.Since(dateDeb).Hours()) < int(time.Since(dateFin).Hours()) {
			gm.ComposeAndShowModalPmangeP([]string{"back"}, "Incoérance entre les dates de formation")
		}
	})

}

func (gm *GuiManager) composeFormManageEntreprise() {
	gm.Fpmanager.Fentreprise.AddInputField("Code", "", 15, nil, nil).
		AddInputField("Nom", "", 30, nil, nil)

	gm.Fpmanager.Fentreprise.SetBorder(true).SetTitle("Formulaire Entreprise")

	for i := 0; i < gm.Fpmanager.Fentreprise.GetFormItemCount(); i++ {
		input := gm.Fpmanager.Fentreprise.GetFormItem(i).(*tview.InputField)
		input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyCtrlA {
				gm.App.SetFocus(gm.Fpmanager.Fformation.GetFormItem(0))
			}
			return event
		})

	}

	/*gm.Fpmanager.Fentreprise.GetFormItem(0).(*tview.InputField).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlA {
			gm.App.SetFocus(gm.Fpmanager.Fformation.GetFormItem(0))
		}
		return event
	})*/
}

func (gm *GuiManager) composeActions() {

	// form field commentaire
	com := tview.NewForm().AddInputField("", "", 0, nil, nil)
	//form actions buttons
	actions := tview.NewForm().AddButton("Enregister", func() {

		gm.createModelPersonneFromFrms()

	}).AddButton("back", func() {

	})

	com.SetBorder(true).SetTitle("Commentaire").SetTitleAlign(tview.AlignLeft)
	gm.Fpmanager.Flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().AddItem(gm.Fpmanager.Fperonne, 0, 1, false).
			AddItem(gm.Fpmanager.Fentreprise, 0, 1, false).
			AddItem(gm.Fpmanager.Fformation, 0, 1, false).
			AddItem(gm.Fpmanager.Fdocument, 0, 1, false), 0, 1, false).
		AddItem(com, 0, 1, false).
		AddItem(actions.SetButtonsAlign(tview.AlignRight), 3, 1, false)

}

func (gm *GuiManager) createModelPersonneFromFrms() {
	const ISO = "2006-01-02"
	var intitule string
	frmPersonne := gm.Fpmanager.Fperonne
	frmFormation := gm.Fpmanager.Fformation
	frmEntreprise := gm.Fpmanager.Fentreprise
	frmDocument := gm.Fpmanager.Fdocument

	// create formation model
	var frm []models.Formation
	//check if the new formation was writed
	if gm.Fpmanager.Fformation.GetFormItem(1).(*tview.InputField).GetText() != "" {
		intitule = gm.Fpmanager.Fformation.GetFormItem(1).(*tview.InputField).GetText()
	} else {
		_, v := frmFormation.GetFormItem(0).(*tview.DropDown).GetCurrentOption()
		intitule = v
	}

	dateDebut, _ := time.Parse(ISO, frmFormation.GetFormItem(2).(*tview.InputField).GetText())
	dateFin, _ := time.Parse(ISO, frmFormation.GetFormItem(3).(*tview.InputField).GetText())
	nbrHeure, _ := strconv.Atoi(frmFormation.GetFormItem(4).(*tview.InputField).GetText())
	cout, _ := strconv.ParseFloat(frmFormation.GetFormItem(5).(*tview.InputField).GetText(), 8)

	frm = append(frm, models.Formation{CreatedAt: time.Now(), Intitule: intitule,
		DateDebut: dateDebut, DateFin: dateFin,
		NbrHeures: nbrHeure,
		Cout:      cout})

	// create document model
	var docs []models.Document
	for i := 0; i < frmDocument.GetFormItemCount(); i++ {
		if frmDocument.GetFormItem(i).(*tview.Checkbox).IsChecked() {
			idDoc := mIdd[frmDocument.GetFormItem(i).(*tview.Checkbox).GetLabel()]
			docs = append(docs, models.Document{Id: idDoc, Libelle: frmDocument.GetFormItem(i).(*tview.Checkbox).GetLabel()})

		}
	}

	age, _ := strconv.Atoi(frmPersonne.GetFormItem(3).(*tview.InputField).GetText())
	dateNanissance, _ := time.Parse(ISO, frmPersonne.GetFormItem(4).(*tview.InputField).GetText())

	// create personne model
	mfa := false

	if gm.Fpmanager.Fperonne.GetFormItem(0).(*tview.Checkbox).IsChecked() {
		mfa = true
	}
	personne := &models.Personne{CreatedAt: time.Now(), Mfa: mfa, Nom: frmPersonne.GetFormItem(1).(*tview.InputField).GetText(),
		Prenom:        frmPersonne.GetFormItem(2).(*tview.InputField).GetText(),
		Age:           age,
		DateNaissance: dateNanissance,
		Tel:           frmPersonne.GetFormItem(5).(*tview.InputField).GetText(),
		Mail:          frmPersonne.GetFormItem(6).(*tview.InputField).GetText(),
		Adresse:       frmPersonne.GetFormItem(7).(*tview.InputField).GetText(),
		Entreprise: models.Entreprise{Code: frmEntreprise.GetFormItem(0).(*tview.InputField).GetText(),
			Nom: frmEntreprise.GetFormItem(1).(*tview.InputField).GetText(),
		},
		Formation: frm,
		Document:  docs}

	fmt.Println(gm.Fpmanager.Flag)

	// check the flag
	if gm.Fpmanager.Flag == false {
		personne.Id = gm.Services.Personne.Id
		personne.Formation[0].Id = gm.Services.Personne.Formation[0].Id
		personne.Entreprise.Id = gm.Services.Personne.Entreprise.Id
	}

	gm.Services.Personne = *personne
	c := make(chan string, 1)
	if gm.Fpmanager.Flag == true {
		go gm.Services.Personne.SavePersonne(c)
		err := <-c
		if err != "" {
			gm.ComposeAndShowModalPmangeP([]string{"back"}, err)
		} else {
			gm.ComposeAndShowModalPmangeP([]string{"back"}, "personne saved successufuly")
		}
	} else {
		go gm.Services.Personne.UpdatePersonne(c)
		err := <-c
		if err != "" {
			gm.ComposeAndShowModalPmangeP([]string{"back"}, err)
		} else {
			gm.ComposeAndShowModalPmangeP([]string{"back"}, "personne updated successufuly")
		}
	}

	gm.Fpmanager.Flag = true

}
