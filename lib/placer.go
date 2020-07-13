package lib

import (
	"fmt"
	"image"
	"log"
	"math"

	"github.com/disintegration/imaging"
)

// Thanks to this blog post, it helped a lot, check it out! https://dev.to/ope__o/how-to-make-a-logo-watermark-tool-in-golang-on-any-unix-distribution-1fae

func PlaceImg(outName, origImg, cheemsImg, cheemsDimentions, locationDimentions string) {
	locaX, locaY := ParseCoordinates(locationDimentions, "x")

	src := OpenImage(origImg)

	cheemsFit := ResizeImage(cheemsImg, cheemsDimentions)

	dst := imaging.Overlay(src, cheemsFit, image.Pt(locaX, locaY), 1.0)

	err := imaging.Save(dst, outName)

	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Placed cheems on image '%s\n' ", origImg)
}

func AddCheems(origImg, cheems string) {
	outName := fmt.Sprintf("cheemed-%s", "test.png")

	src := OpenImage(origImg)

	cheemsFit := ResizeImage(cheems, "754x502")

	bgDimensions := src.Bounds().Max
	markDimensions := cheemsFit.Bounds().Max

	bgAspectRatio := math.Round(float64(bgDimensions.X) / float64(bgDimensions.Y))

	xPos, yPos := CalcCheemsPosition(bgDimensions, markDimensions, bgAspectRatio)

	PlaceImg(outName, origImg, cheems, "754x502", fmt.Sprintf("%dx%d", xPos, yPos))

	fmt.Printf("Cheemified '%s'  image '%s' with Cheems dimensions %s.\n", cheems, origImg, "754x502")
}
