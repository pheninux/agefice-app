package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Pmain struct {
	Grid            *tview.Grid
	Tree            *tview.TreeView
	Pages           *tview.Pages
	tableAll        *tview.Table
	tableDetailForm *tview.Table
	tableDetailDoc  *tview.Table
	tableDetail     *tview.Table
	flex            *tview.Flex
}
type node struct {
	text     string
	expand   bool
	selected func()
	children []*node
}

var rootNode *node

func (gm *GuiManager) initPmain() {
	gm.Pmain.Grid = tview.NewGrid()
	gm.Pmain.Tree = tview.NewTreeView()
	gm.Pmain.Pages = tview.NewPages()
	gm.Pmain.tableAll = tview.NewTable()
	gm.Pmain.tableDetailForm = tview.NewTable()
	gm.Pmain.tableDetailDoc = tview.NewTable()
	gm.Pmain.flex = tview.NewFlex()
}

func (gm *GuiManager) ComposeAndShowPmain() {

	gm.initPmain()
	gm.composeGridPmain()
	gm.manageFocus()
	gm.ComposeAndShowPmanagePersonne()
	gm.App.SetRoot(gm.Pmain.Grid, true).SetFocus(gm.Pmain.Tree).EnableMouse(true)

}

func (gm *GuiManager) composeGridPmain() {

	gm.Pmain.Grid.SetColumns(40, -1)
	gm.Pmain.Grid.SetRows(-1, 3)
	gm.Pmain.Grid.SetBorders(true)
	gm.composeTreePmain()
	gm.composePgaesPmain()
	gm.Pmain.Grid.AddItem(gm.Pmain.Tree, 0, 0, 1, 1, 0, 0, false)
	gm.Pmain.Grid.AddItem(gm.Pmain.Pages, 0, 1, 1, 1, 0, 0, false)
}

func (gm *GuiManager) composePgaesPmain() {
	// form field commentaire
	mfa := tview.NewForm().AddCheckbox("MFA", false, func(checked bool) {
		gm.ComposeTableGetAll("getByFma", checked)
	})
	mfa.SetBorder(true).SetTitle("FILTRE").SetTitleAlign(tview.AlignLeft)

	gm.Pmain.flex.AddItem(mfa, 5, 1, false).
		AddItem(gm.Pmain.tableAll, 0, 1, false).
		AddItem(tview.NewFlex().
			AddItem(gm.Pmain.tableDetailDoc, 0, 1, false).
			AddItem(gm.Pmain.tableDetailForm, 0, 3, false), 0, 1, false).
		SetDirection(tview.FlexRow)

	gm.Pmain.Pages.AddPage("liste-stagiaires-page", gm.Pmain.flex, true, false)

}

func (gm *GuiManager) composeRootNode() {
	rootNode = &node{
		text: "Navigation",
		children: []*node{
			/*{text: "Expand all", selected: func() { gm.Pmain.Tree.GetRoot().ExpandAll() }},
			{text: "Collapse all", selected: func() {
				for _, child := range gm.Pmain.Tree.GetRoot().GetChildren() {
					child.CollapseAll()
				}
			}},*/
			{text: "Stagiaire", expand: true, children: []*node{
				{text: "Liste des stagiaires", selected: func() {
					gm.ComposeTableGetAll("getAll", false)
					gm.Pmain.Pages.SwitchToPage("liste-stagiaires-page")

				}},
				{text: "Gesion des stagiaires", selected: func() {
					//gm.ComposeAndShowPmanagePersonne()
					gm.Pmain.Pages.SwitchToPage("manage-stagiaire-page")
					gm.App.SetFocus(gm.Fpmanager.Fperonne)

				}},
			}},
			{text: "Entreprise", expand: true, children: []*node{
				{text: "Liste des entreprises", selected: func() {

				}},
				{text: "Gestion de l'entreprise", selected: func() {

				}},
			}},
			{text: "Formation", expand: true, children: []*node{
				{text: "Gestion de formation", selected: func() {
					gm.ComposeAndShowPmanageFormation()
					gm.Pmain.Pages.SwitchToPage("page_manage_formation")
				}},
			}},
		}}
}

func (gm *GuiManager) composeTreePmain() {

	gm.composeRootNode()

	gm.Pmain.Tree.SetBorder(true).
		SetTitle("Menu")
	gm.Pmain.Tree.SetAlign(true).SetTopLevel(0).SetGraphics(true).SetPrefixes(nil)
	//gm.GuiManager.Tree.SetAlign(false).SetTopLevel(1).SetGraphics(true).SetPrefixes(nil)
	gm.Pmain.Tree.SetGraphicsColor(tcell.ColorGreenYellow)
	// Add nodes.
	var add func(target *node) *tview.TreeNode
	add = func(target *node) *tview.TreeNode {
		node := tview.NewTreeNode(target.text).
			SetSelectable(target.expand || target.selected != nil).
			SetExpanded(target == rootNode).
			SetReference(target)
		if target.expand {
			node.SetColor(tcell.ColorGreen)
		} else if target.selected != nil {
			node.SetColor(tcell.ColorOlive)
		}
		for _, child := range target.children {
			node.AddChild(add(child))
		}
		return node
	}
	root := add(rootNode)
	gm.Pmain.Tree.SetRoot(root).
		SetCurrentNode(root).
		SetSelectedFunc(func(n *tview.TreeNode) {
			original := n.GetReference().(*node)
			if original.expand {
				n.SetExpanded(!n.IsExpanded())
			} else if original.selected != nil {
				original.selected()
			}
		})

}

func (gm *GuiManager) manageFocus() {
	gm.Pmain.tableAll.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyCtrlF {
			gm.App.SetFocus(gm.Pmain.Tree)

		}
		return event
	})
}
