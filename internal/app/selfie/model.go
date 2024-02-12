package selfie

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"regexp"
	"sort"
)

type SelfieSetStatus string

const (
	SelfieSetStatusNotHandled = SelfieSetStatus("NOT_HANDLED")
	SelfieSetStatusOk         = SelfieSetStatus("OK")
	SelfieSetStatusSuspicious = SelfieSetStatus("SUSPICIOUS")
	SelfieSetStatusFake       = SelfieSetStatus("FAKE")
)

func (s SelfieSetStatus) String() string {
	return string(s)
}

func IconAttributesFromStatus(status SelfieSetStatus, active bool) (fyne.Resource, widget.Importance, color.Color) {
	var iconResource fyne.Resource
	var importance widget.Importance
	var c color.Color

	switch status {
	case SelfieSetStatusNotHandled:
		iconResource, importance, c = theme.QuestionIcon(), widget.HighImportance, theme.PrimaryColor()
	case SelfieSetStatusOk:
		iconResource, importance, c = theme.ConfirmIcon(), widget.SuccessImportance, theme.SuccessColor()
	case SelfieSetStatusSuspicious:
		iconResource, importance, c = theme.WarningIcon(), widget.WarningImportance, theme.WarningColor()
	case SelfieSetStatusFake:
		iconResource, importance, c = theme.ErrorIcon(), widget.DangerImportance, theme.ErrorColor()
	default:
		iconResource, importance, c = theme.CancelIcon(), widget.LowImportance, color.Transparent
	}

	if active {
		iconResource = theme.NewInvertedThemedResource(iconResource)
	} else {
		importance = widget.MediumImportance
		c = theme.DisabledButtonColor()
	}

	return iconResource, importance, c
}

type SelfieSet struct {
	Name            string          `json:"name"`
	PrimaryImage    *ImageFile      `json:"primary_image"`
	SecondaryImages []*ImageFile    `json:"secondary_images"`
	Status          SelfieSetStatus `json:"status"`
}

func ConvertToSelfieSets(imageFiles []*ImageFile) []*SelfieSet {
	var selfieSets []*SelfieSet

	var compRegEx = regexp.MustCompile("([[:alpha:]]+(\\d+))_(\\d+)\\.jpg") // Pattern for file names like "selfie01_03.jpg"

	selfieSetImagesMap := make(map[string][]*ImageFile)
	for _, imageFile := range imageFiles {
		match := compRegEx.FindStringSubmatch(imageFile.FileName)
		selfieSetImagesMap[match[1]] = append(selfieSetImagesMap[match[1]], imageFile)
	}

	var selfieSetNames []string
	for selfieSetName := range selfieSetImagesMap {
		selfieSetNames = append(selfieSetNames, selfieSetName)
	}
	sort.Strings(selfieSetNames)

	for _, selfieSetName := range selfieSetNames {
		selfieSetImages := selfieSetImagesMap[selfieSetName]
		selfieSet := &SelfieSet{
			Name:            selfieSetName,
			PrimaryImage:    selfieSetImages[0],
			SecondaryImages: selfieSetImages[1:],
			Status:          SelfieSetStatusNotHandled,
		}

		selfieSets = append(selfieSets, selfieSet)
	}

	return selfieSets
}
