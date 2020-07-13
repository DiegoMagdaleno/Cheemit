package lib

import (
	"image"
	"log"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

func ParseCoordinates(input, delimeter string) (int, int) {

	cordArray := strings.Split(input, delimeter)

	x, err := strconv.Atoi(cordArray[0])

	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(cordArray[1])

	if err != nil {
		log.Fatal(err)
	}

	return x, y

}

func OpenImage(path string) image.Image {
	src, err := imaging.Open(path)
	if err != nil {
		log.Fatal(err)
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
