package cmd

import (
	"encoding/json"
	"fmt"

	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/jeanmarcboite/librarytruc/pkg/books"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "List files in Epub",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

		log.Debug().Str("args", fmt.Sprint(args)).Msg("lookup")

		// debug, _ := cmd.Flags().GetBool("debug")

		for _, filename := range args {
			work, error := books.WorkFromFilename(filename)

			if error != nil {
				log.Error().Str("file", filename).Msg(error.Error())
			} else {
				xmlj, _ := json.MarshalIndent(work, "", "    ")
				fmt.Println(string(xmlj))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)
	lookupCmd.Flags().BoolP("debug", "d", false, "Print xml parsed to stdout")
}
