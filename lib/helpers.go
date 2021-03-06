package lib

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/disintegration/imaging"
)

var prefix string

func ParseCoordinates(input, delimeter string) (int, int) {

	cordArray := strings.Split(input, delimeter)

	x, err := strconv.Atoi(cordArray[0])

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("An error ocurred while converting dimentions of X from string to int.")
	}

	y, err := strconv.Atoi(cordArray[1])

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("An error ocurred while converting dimentions of Y from string to int.")
	}

	return x, y

}

func OpenImage(path string) image.Image {
	src, err := imaging.Open(path)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("An error ocurred while opening the desired image.")
	}
	return src
}

func ResizeImage(img, dimensions string) image.Image {
	width, height := ParseCoordinates(dimensions, "x")
	src := OpenImage(img)
	return imaging.Fit(src, width, height, imaging.CatmullRom)
}

func CalcCheemsPosition(origDimentions, cheemsDimentions image.Point) (int, int) {
	origX := origDimentions.X
	origY := origDimentions.Y
	cheemsX := cheemsDimentions.X
	cheemsY := cheemsDimentions.Y

	return origX - cheemsX, origY - cheemsY
}

func calculateCheemsSize(origImageSize image.Point) string {
	originalX := origImageSize.X
	originalY := origImageSize.Y

	// We want to calculate the 10% of an image to place cheems so we do that with the following:
	// We want to calculate the 30% of an image to place cheems so we do that with the following:
	tenPerOfX := originalX / 100 * 30
	tenPerOfY := originalY / 100 * 30

	s1 := strconv.Itoa(tenPerOfX)
	s2 := strconv.Itoa(tenPerOfY)

	return fmt.Sprintf("%vx%v", s1, s2)
}
