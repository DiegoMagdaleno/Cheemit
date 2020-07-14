package lib

import (
	"fmt"
	"image"
	"log"

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

	baseImgName := calculateName(cheems)
	outName := fmt.Sprintf("cheemed-%s", baseImgName)

	src := OpenImage(origImg)

	calculatedDimentions := calculateCheemsSize(src.Bounds().Max)

	cheemsFit := ResizeImage(cheems, calculatedDimentions)

	bgDimensions := src.Bounds().Max
	markDimensions := cheemsFit.Bounds().Max

	xPos, yPos := CalcCheemsPosition(bgDimensions, markDimensions)

	PlaceImg(outName, origImg, cheems, calculatedDimentions, fmt.Sprintf("%dx%d", xPos, yPos))

	fmt.Printf("Cheemified '%s'  image '%s' with Cheems dimensions %s.\n", cheems, origImg, calculatedDimentions)
}
