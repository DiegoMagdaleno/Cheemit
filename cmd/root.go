package cmd

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/diegomagdaleno/cheemit/lib"
	"github.com/spf13/cobra"
)

var character string

var rootCmd = &cobra.Command{
	Use:   "cheemit",
	Short: "Cheemit takes an image and puts internet's favorite dog on top of it!",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			originalImage = args[0]
		)
		switch character = strings.ToLower(character); character {
		case "cheems":
			lib.AddCharacter(originalImage, "cheems", "hey")
		case "doge":
			lib.AddCharacter(originalImage, "doge", "hey")
		default:
			log.WithFields(log.Fields{
				"character": character,
			}).Fatal("Not a valid character")
		}
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&character, "character", "c", "cheems", "Allows you to set the Doge Lore character")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
