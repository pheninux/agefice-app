package ui

import (
	models "agefice-cons/adil.net/pkg/model"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strconv"
)

func (gm *GuiManager) ComposeTableGetAll(action string, mfa bool) {

	gm.funcName(action, mfa)

}

func (gm *GuiManager) funcName(action string, mfa bool) {
	//clear table
	gm.Pmain.tableAll.Clear()
	gm.Pmain.tableDetailForm.Clear()
	gm.Pmain.tableDetailDoc.Clear()

	//Create channel
	c := make(chan []models.Personne, 1)
	e := make(chan error, 1)
	if action == "getAll" {
		go gm.Services.Personne.GetAll(c, e, false)
	} else {
		go gm.Services.Personne.GetByMfa(c, e, mfa)
	}

	for {
		select {
		case err, ok := <-e:
			if ok {
				gm.ComposeAndShowModalPmangeP([]string{"back"}, err.Error())
				//gm.Pmain.tableAll.SetCell(0, 0, &tview.TableCell{Color: tcell.ColorIndianRed, Text: fmt.Sprintf("Erreur :%s", err)})
				return
			}
		case p := <-c:

			if len(p) > 0 {

				gm.Pmain.tableAll.SetFixed(1, 1)
				gm.Pmain.tableAll.SetSeparator(tcell.RuneVLine)

				row := 1
				personneMap := make(map[int]models.Personne)
				for _, v := range p {

					gm.createHeadersTableAll()
					gm.Pmain.tableAll.SetCell(row, 0, &tview.TableCell{Color: tcell.ColorDarkCyan, Text: strconv.Itoa(v.Id), NotSelectable: true})
					gm.Pmain.tableAll.SetCell(row, 1, &tview.TableCell{Color: tcell.ColorWhite, Text: v.CreatedAt.Format("2006-01-02")})
					gm.Pmain.tableAll.SetCell(row, 2, &tview.TableCell{Color: tcell.ColorWhite, Text: v.Nom})
					gm.Pmain.tableAll.SetCell(row, 3, &tview.TableCell{Color: tcell.ColorWhite, Text: v.Prenom})
					gm.Pmain.tableAll.SetCell(row, 4, &tview.TableCell{Color: tcell.ColorWhite, Text: strconv.Itoa(v.Age)})
					gm.Pmain.tableAll.SetCell(row, 5, &tview.TableCell{Color: tcell.ColorWhite, Text: v.DateNaissance.Format("2006-01-02")})
					gm.Pmain.tableAll.SetCell(row, 6, &tview.TableCell{Color: tcell.ColorWhite, Text: v.Tel})
					gm.Pmain.tableAll.SetCell(row, 7, &tview.TableCell{Color: tcell.ColorWhite, Text: v.Mail})
					gm.Pmain.tableAll.SetCell(row, 8, &tview.TableCell{Color: tcell.ColorWhite, Text: v.Adresse})
					gm.Pmain.tableAll.SetCell(row, 9, &tview.TableCell{Color: tcell.ColorWhite, Text: v.Entreprise.Nom})

					personneMap[v.Id] = v
					row++
				}

				// event and func table
				gm.Pmain.tableAll.SetSelectionChangedFunc(func(row int, column int) {

					if gm.manageFormationAndDocumentTables(row, personneMap) {
						return
					}

					gm.Pmain.tableAll.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
						if event.Key() == tcell.KeyCtrlS {
							idToDelete, _ := strconv.Atoi(gm.Pmain.tableAll.GetCell(row, 0).Text)
							gm.Services.Personne.Id = idToDelete
							msg := fmt.Sprintf("La personne avec l'ID : %v va étre supprimer !", idToDelete)
							gm.ComposeAndShowModalPmangePConfirmation([]string{"oui", "non"}, msg, "delete")
						}

						if event.Key() == tcell.KeyCtrlJ {

							// change flag create to update
							gm.Fpmanager.Flag = false
							// initialise les form
							gm.ComposeAndShowPmanagePersonne()

							// bind personne with service
							id, _ := strconv.Atoi(gm.Pmain.tableAll.GetCell(row, 0).Text)
							gm.Services.Personne = personneMap[id]
							//bind form personne
							gm.Fpmanager.Fperonne.GetFormItem(0).(*tview.InputField).SetText(gm.Pmain.tableAll.GetCell(row, 2).Text)
							gm.Fpmanager.Fperonne.GetFormItem(1).(*tview.InputField).SetText(gm.Pmain.tableAll.GetCell(row, 3).Text)
							gm.Fpmanager.Fperonne.GetFormItem(2).(*tview.InputField).SetText(gm.Pmain.tableAll.GetCell(row, 4).Text)
							gm.Fpmanager.Fperonne.GetFormItem(3).(*tview.InputField).SetText(gm.Pmain.tableAll.GetCell(row, 5).Text)
							gm.Fpmanager.Fperonne.GetFormItem(4).(*tview.InputField).SetText(gm.Pmain.tableAll.GetCell(row, 6).Text)
							gm.Fpmanager.Fperonne.GetFormItem(5).(*tview.InputField).SetText(gm.Pmain.tableAll.GetCell(row, 7).Text)
							gm.Fpmanager.Fperonne.GetFormItem(6).(*tview.InputField).SetText(gm.Pmain.tableAll.GetCell(row, 8).Text)

							// bind form entreprise
							gm.Fpmanager.Fentreprise.GetFormItem(0).(*tview.InputField).SetText(personneMap[id].Entreprise.Code)
							gm.Fpmanager.Fentreprise.GetFormItem(1).(*tview.InputField).SetText(gm.Pmain.tableAll.GetCell(row, 9).Text)

							//bind form formation
							gm.Fpmanager.Fformation.GetFormItem(1).(*tview.InputField).SetText(personneMap[id].Formation[0].Intitule)
							gm.Fpmanager.Fformation.GetFormItem(2).(*tview.InputField).SetText(personneMap[id].Formation[0].DateDebut.Format("2006-01-02"))
							gm.Fpmanager.Fformation.GetFormItem(3).(*tview.InputField).SetText(personneMap[id].Formation[0].DateFin.Format("2006-01-02"))
							gm.Fpmanager.Fformation.GetFormItem(4).(*tview.InputField).SetText(strconv.Itoa(personneMap[id].Formation[0].NbrHeures))
							gm.Fpmanager.Fformation.GetFormItem(5).(*tview.InputField).SetText(fmt.Sprintf("%.2f", personneMap[id].Formation[0].Cout))

							//bind form document

							docs := personneMap[id].Document
							for _, v := range docs {
								for i := 0; i < gm.Fpmanager.Fdocument.GetFormItemCount(); i++ {
									if gm.Fpmanager.Fdocument.GetFormItem(i).(*tview.Checkbox).GetLabel() == v.Libelle {
										gm.Fpmanager.Fdocument.GetFormItem(i).(*tview.Checkbox).SetChecked(true)
									}
								}
							}
							//idToUpdate, _ := strconv.Atoi(gm.Pmain.tableAll.GetCell(row, 0).Text)
							//msg := fmt.Sprintf("La personne avec l'ID : %v va étre supprimer !", idToDelete)

							gm.Pmain.Pages.SwitchToPage("manage-stagiaire-page")
						}
						return event
					})

				})

				gm.Pmain.tableAll.SetSelectedFunc(func(row, column int) {
					if gm.manageFormationAndDocumentTables(row, personneMap) {
						return
					}
				})

				gm.Pmain.tableAll.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
					if event.Key() == tcell.KeyEnter {
						gm.Pmain.tableAll.SetSelectable(true, false)
					}
					return event
				})

				gm.Pmain.tableAll.SetBorder(true).SetTitle("Liste des stagiares")
				gm.App.SetRoot(gm.Pmain.Grid, true).SetFocus(gm.Pmain.tableAll).EnableMouse(true)
			}
			return
		default:

		}
	}
}

func (gm *GuiManager) manageFormationAndDocumentTables(row int, personneMap map[int]models.Personne) bool {
	if row == -1 {
		return true
	}
	id, _ := strconv.Atoi(gm.Pmain.tableAll.GetCell(row, 0).Text)
	gm.createDetailDocumentsTable(id, personneMap)
	gm.createDetailFormationTable(id, personneMap)
	return false
}

func (gm *GuiManager) createHeadersTableAll() {

	gm.Pmain.tableAll.SetCell(0, 0, &tview.TableCell{Color: tcell.ColorYellow, Text: "ID", NotSelectable: true})
	gm.Pmain.tableAll.SetCell(0, 1, &tview.TableCell{Color: tcell.ColorYellow, Text: "DATE DE CREATION", NotSelectable: true})
	gm.Pmain.tableAll.SetCell(0, 2, &tview.TableCell{Color: tcell.ColorYellow, Text: "NOM", NotSelectable: true})
	gm.Pmain.tableAll.SetCell(0, 3, &tview.TableCell{Color: tcell.ColorYellow, Text: "PRENOM", NotSelectable: true})
	gm.Pmain.tableAll.SetCell(0, 4, &tview.TableCell{Color: tcell.ColorYellow, Text: "AGE", NotSelectable: true})
	gm.Pmain.tableAll.SetCell(0, 5, &tview.TableCell{Color: tcell.ColorYellow, Text: "DATE NAISSANCE", NotSelectable: true})
	gm.Pmain.tableAll.SetCell(0, 6, &tview.TableCell{Color: tcell.ColorYellow, Text: "TEL", NotSelectable: true})
	gm.Pmain.tableAll.SetCell(0, 7, &tview.TableCell{Color: tcell.ColorYellow, Text: "MAIL", NotSelectable: true})
	gm.Pmain.tableAll.SetCell(0, 8, &tview.TableCell{Color: tcell.ColorYellow, Text: "ADRESSE", NotSelectable: true})
	gm.Pmain.tableAll.SetCell(0, 9, &tview.TableCell{Color: tcell.ColorYellow, Text: "ENTREPRISE", NotSelectable: true})
}

func (gm *GuiManager) createDetailDocumentsTable(id int, m map[int]models.Personne) {

	gm.Pmain.tableDetailDoc.SetFixed(1, 1)
	gm.Pmain.tableDetailDoc.SetSelectable(true, false)
	gm.Pmain.tableDetailDoc.SetSeparator(tcell.RuneVLine)
	gm.Pmain.tableDetailDoc.Clear()
	row := 1
	column := 0

	headers := []string{"ID", "LIBELLE"}
	for _, d := range m[id].Document {
		if column == 0 {
			for _, h := range headers {
				gm.Pmain.tableDetailDoc.SetCell(0, column, &tview.TableCell{Color: tcell.ColorYellow, Text: h, NotSelectable: true})
				column++
			}
		}

		gm.Pmain.tableDetailDoc.SetCell(row, 0, &tview.TableCell{Color: tcell.ColorDarkCyan, Text: strconv.Itoa(d.Id), NotSelectable: true})
		gm.Pmain.tableDetailDoc.SetCell(row, 1, &tview.TableCell{Color: tcell.ColorWhite, Text: d.Libelle})
		row++
	}
	if len(m[id].Document) == 0 {
		gm.Pmain.tableDetailDoc.SetCell(0, 0, &tview.TableCell{Text: "PAS DE DOCUMENTS DISPONIBLES", Color: tcell.ColorIndianRed})
	}

	gm.Pmain.tableDetailDoc.SetBorder(true).SetTitle("Documents")
}

func (gm *GuiManager) createDetailFormationTable(id int, m map[int]models.Personne) {
	gm.Pmain.tableDetailForm.SetFixed(1, 1)
	gm.Pmain.tableDetailForm.SetSelectable(true, false)
	gm.Pmain.tableDetailForm.SetSeparator(tcell.RuneVLine)
	gm.Pmain.tableDetailForm.Clear()
	row := 1
	column := 0

	headers := []string{"ID", "INTITULE", "DATE DEBUT", "DATE FIN", "COUT", "NBR H"}
	for _, f := range m[id].Formation {
		if column == 0 {
			for _, h := range headers {
				gm.Pmain.tableDetailForm.SetCell(0, column, &tview.TableCell{Color: tcell.ColorYellow, Text: h, NotSelectable: true})
				column++
			}
		}

		gm.Pmain.tableDetailForm.SetCell(row, 0, &tview.TableCell{Color: tcell.ColorDarkCyan, Text: strconv.Itoa(f.Id), NotSelectable: true})
		gm.Pmain.tableDetailForm.SetCell(row, 1, &tview.TableCell{Color: tcell.ColorWhite, Text: f.Intitule})
		gm.Pmain.tableDetailForm.SetCell(row, 2, &tview.TableCell{Color: tcell.ColorWhite, Text: f.DateDebut.Format("02/01/2006")})
		gm.Pmain.tableDetailForm.SetCell(row, 3, &tview.TableCell{Color: tcell.ColorWhite, Text: f.DateFin.Format("02/01/2006")})
		gm.Pmain.tableDetailForm.SetCell(row, 4, &tview.TableCell{Color: tcell.ColorWhite, Text: fmt.Sprintf("%.2f", f.Cout), Align: tview.AlignRight})
		gm.Pmain.tableDetailForm.SetCell(row, 5, &tview.TableCell{Color: tcell.ColorWhite, Text: strconv.Itoa(f.NbrHeures), Align: tview.AlignRight})
		row++
	}

	if len(m[id].Formation) == 0 {
		gm.Pmain.tableDetailForm.SetCell(0, 0, &tview.TableCell{Text: "PAS DE FORMATIONS DISPONIBLES", Color: tcell.ColorIndianRed})
	}

	gm.Pmain.tableDetailForm.SetBorder(true).SetTitle("Formations")
}
