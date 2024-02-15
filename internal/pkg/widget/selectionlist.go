package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type SelectionList struct {
	widget.List
}

// NewSelectionList creates and returns a list widget for displaying items in
// a vertical layout with scrolling and caching for performance.
// This list is very similar to widget.List but it has no key listeners by default.
func NewSelectionList(length func() int, createItem func() fyne.CanvasObject, updateItem func(widget.ListItemID, fyne.CanvasObject)) *SelectionList {
	selectionList := &SelectionList{List: widget.List{Length: length, CreateItem: createItem, UpdateItem: updateItem}}
	selectionList.ExtendBaseWidget(selectionList)
	return selectionList
}

func (l *SelectionList) TypedKey(_ *fyne.KeyEvent) {
	// Left empty by intention
	// Removes the default key listen behavior of widget.List
}
