package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"thisisfyne/internal/app/selfie"
)

type SelfieSetStatusButtons struct {
	widget.BaseWidget

	status selfie.SelfieSetStatus

	tapped func(status selfie.SelfieSetStatus)

	buttonNotHandled *widget.Button
	buttonOk         *widget.Button
	buttonSuspicious *widget.Button
	buttonFake       *widget.Button
}

func NewSelfieSetStatusButtons(tapped func(status selfie.SelfieSetStatus)) *SelfieSetStatusButtons {
	statusSelectButtons := &SelfieSetStatusButtons{tapped: tapped}

	statusSelectButtons.ExtendBaseWidget(statusSelectButtons)

	return statusSelectButtons
}

func (b *SelfieSetStatusButtons) CreateRenderer() fyne.WidgetRenderer {
	b.buttonNotHandled = widget.NewButton("", func() { b.setSelfieSetStatus(selfie.SelfieSetStatusNotHandled) })
	b.buttonOk = widget.NewButton("", func() { b.setSelfieSetStatus(selfie.SelfieSetStatusOk) })
	b.buttonSuspicious = widget.NewButton("", func() { b.setSelfieSetStatus(selfie.SelfieSetStatusSuspicious) })
	b.buttonFake = widget.NewButton("", func() { b.setSelfieSetStatus(selfie.SelfieSetStatusFake) })

	b.updateUI()

	c := container.NewGridWithColumns(4, b.buttonNotHandled, b.buttonOk, b.buttonSuspicious, b.buttonFake)

	return widget.NewSimpleRenderer(c)
}

func (b *SelfieSetStatusButtons) setSelfieSetStatus(status selfie.SelfieSetStatus) {
	b.SetSelfieSetStatus(status)
	if b.tapped != nil {
		b.tapped(b.status)
	}
}

func (b *SelfieSetStatusButtons) SetSelfieSetStatus(status selfie.SelfieSetStatus) {
	b.status = status
	b.updateUI()
	b.Refresh()
}

func (b *SelfieSetStatusButtons) updateUI() {
	if b.buttonNotHandled != nil {
		b.buttonNotHandled.Icon, b.buttonNotHandled.Importance, _ = selfie.IconAttributesFromStatus(selfie.SelfieSetStatusNotHandled, b.getStatus() == selfie.SelfieSetStatusNotHandled)
		b.buttonOk.Icon, b.buttonOk.Importance, _ = selfie.IconAttributesFromStatus(selfie.SelfieSetStatusOk, b.getStatus() == selfie.SelfieSetStatusOk)
		b.buttonSuspicious.Icon, b.buttonSuspicious.Importance, _ = selfie.IconAttributesFromStatus(selfie.SelfieSetStatusSuspicious, b.getStatus() == selfie.SelfieSetStatusSuspicious)
		b.buttonFake.Icon, b.buttonFake.Importance, _ = selfie.IconAttributesFromStatus(selfie.SelfieSetStatusFake, b.getStatus() == selfie.SelfieSetStatusFake)
	}
}

func (b *SelfieSetStatusButtons) getStatus() selfie.SelfieSetStatus {
	return b.status
}
