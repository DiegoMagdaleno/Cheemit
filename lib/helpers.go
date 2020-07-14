package lib

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
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
		log.Panic(err)
	}
	return file.Name() + ".png"
}
