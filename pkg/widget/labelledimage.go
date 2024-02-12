package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
	"image"
)

type LabelledImage struct {
	widget.BaseWidget

	Text  string
	Image *image.Image

	ImageSize int

	// private UI components
	textUI  *canvas.Text
	imageUI *canvas.Image
}

func NewLabelledImage(text string, image *image.Image, size int) *LabelledImage {
	labelledImage := &LabelledImage{
		Text:      text,
		Image:     image,
		ImageSize: size,
	}

	labelledImage.ExtendBaseWidget(labelledImage)

	return labelledImage
}

//func (li *LabelledImage) Tapped(*fyne.PointEvent) {
//	log.Printf("Clicked labelled image '%s'", li.Text)
//}

func (li *LabelledImage) CreateRenderer() fyne.WidgetRenderer {
	text := canvas.NewText(li.Text, colornames.Gray)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 10

	img := canvas.NewImageFromImage(*li.Image)
	img.SetMinSize(fyne.NewSquareSize(float32(li.ImageSize)))
	img.ScaleMode = canvas.ImageScaleSmooth
	img.FillMode = canvas.ImageFillContain
	//canvasImage.FillMode = canvas.ImageFillStretch

	li.textUI = text
	li.imageUI = img

	c := container.NewVBox(img, text)
	c = container.NewPadded(c)

	return widget.NewSimpleRenderer(c)
}

func (li *LabelledImage) SetImage(img *image.Image) {
	li.Image = img
	li.imageUI.Image = *img
	li.Refresh()
}

func (li *LabelledImage) SetText(text string) {
	li.Text = text
	li.textUI.Text = text
	li.Refresh()
}
