package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"thisisfyne/internal/app/selfie"
)

type StatusIcon struct {
	widget.BaseWidget

	status selfie.SelfieSetStatus

	icon       *widget.Icon
	background *canvas.Rectangle
}

func NewStatusIcon() *StatusIcon {
	statusIcon := &StatusIcon{}

	statusIcon.ExtendBaseWidget(statusIcon)

	return statusIcon
}

func (item *StatusIcon) CreateRenderer() fyne.WidgetRenderer {
	item.icon = widget.NewIcon(nil)

	item.background = canvas.NewRectangle(color.Transparent)
	item.background.CornerRadius = theme.SelectionRadiusSize()

	item.updateUI()

	c := container.NewStack(item.background, container.NewPadded(item.icon))

	return widget.NewSimpleRenderer(c)
}

func (item *StatusIcon) SetSelfieSetStatus(status selfie.SelfieSetStatus) {
	item.status = status
	item.updateUI()
	item.Refresh()
}

func (item *StatusIcon) updateUI() {
	iconResource, _, c := selfie.IconAttributesFromStatus(item.status, true)
	item.icon.Resource = iconResource
	item.background.FillColor = c
}
