package lib

import (
	"fmt"
	"image"

	log "github.com/sirupsen/logrus"

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
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("An error ocurred while attemping to save the image")
	}

	log.WithFields(log.Fields{
		"Original Image": origImg,
	}).Info("Cheems was placed on image!")
}

func AddCharacter(origImg, character, desiredText string) {

	baseImgName := calculateName(origImg)
	characterPath := getCharacterPath(character)
	outName := fmt.Sprintf("cheemed-%s", baseImgName)

	src := OpenImage(origImg)

	calculatedDimentions := calculateCheemsSize(src.Bounds().Max)

	cheemsFit := ResizeImage(characterPath, calculatedDimentions)

	bgDimensions := src.Bounds().Max
	markDimensions := cheemsFit.Bounds().Max

	xPos, yPos := CalcCheemsPosition(bgDimensions, markDimensions)

	PlaceImg(outName, origImg, characterPath, calculatedDimentions, fmt.Sprintf("%dx%d", xPos, yPos))

	log.WithFields(log.Fields{
		"Cheems dimentions": calculatedDimentions,
	}).Infof("Image '%s' was cheemified successfully!\n", origImg)
}
