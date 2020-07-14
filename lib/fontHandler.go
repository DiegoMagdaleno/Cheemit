package lib

import (
	"log"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

func CreateText(text, textDimentions string) {
	splitDimentions := strings.Split(textDimentions, "x")
	dimenX, err := strconv.Atoi(splitDimentions[0])
	if err != nil {
		log.Panic("Error while parsing the text dimentions to int")
	}
	dimenY, err := strconv.Atoi(splitDimentions[1])
	if err != nil {
		log.Panic("Error while parsing the text dimentions to int")
	}
	imageSession := gg.NewContext(dimenX, dimenY)
	imageSession.SetRGBA(0, 0, 0, 0.1)
	imageSession.Clear()
	fontHeight := calculateFontHeight(dimenX, dimenY)
	if err = imageSession.LoadFontFace("/Users/me/Library/Fonts/Anton-Regular.ttf", fontHeight); err != nil {
		log.Panic("Error while loading fontface, is the font face installed?")
	}
	imageSession.SetRGB(0, 0, 0)
	stroke := 4
	for imageSessionY := -stroke; imageSessionY <= stroke; imageSessionY++ {
		for imageSessionX := -stroke; imageSessionX <= stroke; imageSessionX++ {
			if imageSessionX*imageSessionX+imageSessionY*imageSessionY >= stroke*stroke {
				continue
			}
			x := float64(dimenX/2) + float64(imageSessionX)
			y := float64(dimenY/2) + float64(imageSessionY)
			imageSession.DrawStringAnchored(text, x, y, 0.5, 0.5)
		}
	}
	imageSession.SetRGB(1, 1, 1)
	imageSession.DrawStringAnchored(text, float64(dimenX/2), float64(dimenY/2), 0.5, 0.5)
	imageSession.SavePNG("test.png")
}
