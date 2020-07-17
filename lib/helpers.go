package lib

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
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

func calculateName(path string) string {
	splitedPath := strings.Split(path, "/")
	return splitedPath[len(splitedPath)-1]
}

func calculateCheemsSize(origImageSize image.Point) string {
	originalX := origImageSize.X
	originalY := origImageSize.Y

	// We want to calculate the 30% of an image to place cheems so we do that with the following:
	tenPerOfX := originalX / 100 * 30
	tenPerOfY := originalY / 100 * 30

	s1 := strconv.Itoa(tenPerOfX)
	s2 := strconv.Itoa(tenPerOfY)

	return fmt.Sprintf("%vx%v", s1, s2)
}

func calculateFontHeight(dimentionX, dimentionY int) float64 {
	return float64(((dimentionX + dimentionY) / 2) / 100 * 10)
}

func calculateTempDir() string {
	/* Keep windows compatibility originally it was planned to load this to memory but syscalls where not
	supported on Windows :/ */
	file, err := ioutil.TempFile(os.TempDir(), "cheemit-")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("An error ocurred while trying to process a temporal file.")
	}
	return file.Name() + ".png"
}

func getCharacterPath(character string) string {
	switch character {
	case "cheems":
		return fmt.Sprintf("%s/share/cheemit/image/Cheems.png", prefix)
	case "doge":
		return fmt.Sprintf("%s/share/cheemit/image/Doge.png", prefix)
	default:
		log.WithFields(log.Fields{
			"character": character,
		}).Fatal("Not a valid character")
	}
	return ""
}

func whereIsTheFont() string {
	return fmt.Sprintf("%s/share/cheemit/font/Anton-Regular.ttf", prefix)
}
