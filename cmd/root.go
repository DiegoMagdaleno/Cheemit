package cmd

import (
	"fmt"
	"os"

	"github.com/diegomagdaleno/cheemit/lib"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cheemit",
	Short: "Cheemit takes an image and puts internet's favorite dog on top of it!",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			originalImage   = args[0]
			outputImagePath = args[1]
		)

		lib.AddCheems(originalImage, outputImagePath)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
