// +build windows

package lib

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func calculateName(path string) string {

	splitedPath := strings.Split(path, "\\")
	return splitedPath[len(splitedPath)-1]
}

func getCharacterPath(character string) string {
	var appData = os.Getenv("APPDATA")
	switch character {
	case "cheems":
		return fmt.Sprintf("%s\\Cheemit\\image\\Cheems.png", appData)
	case "doge":
		return fmt.Sprintf("%s\\Cheemit\\image\\Doge.png", appData)
	default:
		log.WithFields(log.Fields{
			"character": character,
		}).Fatal("Not a valid character")
	}
	return ""
}
