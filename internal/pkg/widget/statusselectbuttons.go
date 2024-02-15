package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"thisisfyne/internal/app/selfie"
)

type StatusSelectButtons struct {
	widget.BaseWidget

	Status     selfie.SelfieSetStatus
	onChangeFn func(status selfie.SelfieSetStatus)

	// private UI components
	buttonNotHandled *widget.Button
	buttonOk         *widget.Button
	buttonSuspicious *widget.Button
	buttonFake       *widget.Button
}

func NewStatusSelectButtons(status selfie.SelfieSetStatus, onChangeFn func(status selfie.SelfieSetStatus)) *StatusSelectButtons {
	statusSelectButtons := &StatusSelectButtons{Status: status, onChangeFn: onChangeFn}
	statusSelectButtons.ExtendBaseWidget(statusSelectButtons)

	return statusSelectButtons
}

//func (li *LabelledImage) Tapped(*fyne.PointEvent) {
//	log.Printf("Clicked labelled image '%s'", li.Text)
//}

func (li *StatusSelectButtons) CreateRenderer() fyne.WidgetRenderer {
	li.buttonNotHandled = widget.NewButton("", func() { li.changeStatus(selfie.SelfieSetStatusNotHandled) })
	li.buttonOk = widget.NewButton("", func() { li.changeStatus(selfie.SelfieSetStatusOk) })
	li.buttonSuspicious = widget.NewButton("", func() { li.changeStatus(selfie.SelfieSetStatusSuspicious) })
	li.buttonFake = widget.NewButton("", func() { li.changeStatus(selfie.SelfieSetStatusFake) })

	li.updateButtonStates()

	c := container.NewGridWithColumns(4, li.buttonNotHandled, li.buttonOk, li.buttonSuspicious, li.buttonFake)

	return widget.NewSimpleRenderer(c)
}

func (li *StatusSelectButtons) Refresh() {
	log.Println("Refresh of select buttons")
	li.updateButtonStates()
	li.BaseWidget.Refresh()
}

func (li *StatusSelectButtons) changeStatus(status selfie.SelfieSetStatus) {
	li.SetStatus(status)

	if li.onChangeFn != nil {
		li.onChangeFn(status)
	}
}

func (li *StatusSelectButtons) SetStatus(status selfie.SelfieSetStatus) {
	li.Status = status
	li.updateButtonStates()
	li.Refresh()
}

func (li *StatusSelectButtons) updateButtonStates() {
	li.buttonNotHandled.Icon, li.buttonNotHandled.Importance, _ = selfie.IconAttributesFromStatus(selfie.SelfieSetStatusNotHandled, li.Status == selfie.SelfieSetStatusNotHandled)
	li.buttonOk.Icon, li.buttonOk.Importance, _ = selfie.IconAttributesFromStatus(selfie.SelfieSetStatusOk, li.Status == selfie.SelfieSetStatusOk)
	li.buttonSuspicious.Icon, li.buttonSuspicious.Importance, _ = selfie.IconAttributesFromStatus(selfie.SelfieSetStatusSuspicious, li.Status == selfie.SelfieSetStatusSuspicious)
	li.buttonFake.Icon, li.buttonFake.Importance, _ = selfie.IconAttributesFromStatus(selfie.SelfieSetStatusFake, li.Status == selfie.SelfieSetStatusFake)
}
