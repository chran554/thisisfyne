//go:generate fyne bundle --output resources_generated.go --package resources --name thisIsFyneImageResource images/thisisfyne.jpg
//go:generate fyne bundle --output resources_generated.go --package resources --name thisIsFyneIconResource --append images/thisisfyne_icon.png

package resources

var (
	ThisIsFyneImageResource = thisIsFyneImageResource
	ThisIsFyneIconResource  = thisIsFyneIconResource
)
