package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"thisisfyne/pkg/layout"
)

type InfoImage struct {
	widget.BaseWidget

	Image *image.Image

	Text1 string
	Text2 string
	Text3 string
	Text4 string
	Text5 string
	Text6 string

	ImageSize int

	// private UI components
	imageUI *canvas.Image

	textUI1 *canvas.Text
	textUI2 *canvas.Text
	textUI3 *canvas.Text
	textUI4 *canvas.Text
	textUI5 *canvas.Text
	textUI6 *canvas.Text
}

func NewInfoImage(image *image.Image, size int, text1, text2, text3, text4, text5, text6 string) *InfoImage {
	infoImage := &InfoImage{
		Text1: text1,
		Text2: text2,
		Text3: text3,
		Text4: text4,
		Text5: text5,
		Text6: text6,

		Image:     image,
		ImageSize: size,
	}

	infoImage.ExtendBaseWidget(infoImage)

	return infoImage
}

//func (li *LabelledImage) Tapped(*fyne.PointEvent) {
//	log.Printf("Clicked labelled image '%s'", li.Text)
//}

func (li *InfoImage) CreateRenderer() fyne.WidgetRenderer {
	text1 := infoText(li.Text1, colornames.Gray, 10)
	text2 := infoText(li.Text2, colornames.Gray, 10)
	text3 := infoText(li.Text3, colornames.Gray, 10)
	text4 := infoText(li.Text4, colornames.Gray, 10)
	text5 := infoText(li.Text5, colornames.Gray, 10)
	text6 := infoText(li.Text6, colornames.Gray, 10)

	li.textUI1 = text1
	li.textUI2 = text2
	li.textUI3 = text3
	li.textUI4 = text4
	li.textUI5 = text5
	li.textUI6 = text6

	img := canvas.NewImageFromImage(*li.Image)
	img.SetMinSize(fyne.NewSquareSize(float32(li.ImageSize)))
	img.ScaleMode = canvas.ImageScaleSmooth
	img.FillMode = canvas.ImageFillContain
	//canvasImage.FillMode = canvas.ImageFillStretch

	li.imageUI = img

	grid := container.New(layout.NewTextFormLayout(),
		infoText("Filename:", colornames.White, 10), text1,
		infoText("Filesize:", colornames.Lightgray, 10), text2,
		infoText("Modified:", colornames.Lightgray, 10), text3,
		infoText("Image index:", colornames.Lightgray, 10), text4,
		infoText("Path:", colornames.Lightgray, 10), text5,
		infoText("info:", colornames.Lightgray, 10), text6,
	)

	c := container.NewVBox(img, grid)
	//c = container.NewPadded(c)

	return widget.NewSimpleRenderer(c)
}

func infoText(text string, color color.RGBA, size int) *canvas.Text {
	canvasText := canvas.NewText(text, color)
	canvasText.TextSize = float32(size)
	return canvasText
}

func (li *InfoImage) SetImage(img *image.Image) {
	li.Image = img
	li.imageUI.Image = *img
	li.Refresh()
}

func (li *InfoImage) SetText(text1, text2, text3, text4, text5, text6 string) {
	li.Text1 = text1
	li.Text2 = text2
	li.Text3 = text3
	li.Text4 = text4
	li.Text5 = text5
	li.Text6 = text6

	li.textUI1.Text = text1
	li.textUI2.Text = text2
	li.textUI3.Text = text3
	li.textUI4.Text = text4
	li.textUI5.Text = text5
	li.textUI6.Text = text6

	li.Refresh()
}
