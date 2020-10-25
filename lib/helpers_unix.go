// +build !windows

package lib

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

func calculateName(path string) string {

	splitedPath := strings.Split(path, "/")
	return splitedPath[len(splitedPath)-1]
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
