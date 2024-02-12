package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type DoubleLabel struct {
	widget.BaseWidget

	text1 *widget.Label
	text2 *widget.Label
}

func NewDoubleLabel(text1, text2 string) *DoubleLabel {
	item := &DoubleLabel{
		text1: widget.NewLabel(text1),
		text2: widget.NewLabel(text2),
	}

	item.ExtendBaseWidget(item)

	return item
}

func (item *DoubleLabel) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewVBox(item.text1, item.text2)
	return widget.NewSimpleRenderer(c)
}
