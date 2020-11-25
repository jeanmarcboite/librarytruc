package cmd

import (
	"encoding/json"
	"fmt"

	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/jeanmarcboite/truc/pkg/books/epub"
	"github.com/spf13/cobra"
)

func someFunction() {
}

// serverCmd represents the server command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files in Epub",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

		log.Debug().Str("args", fmt.Sprint(args)).Msg("ls")

		debug, _ := cmd.Flags().GetBool("debug")

		for _, filename := range args {
			ereader, error := epub.OpenReader(filename)

			if error != nil {
				log.Error().Str("file", filename).Msg(error.Error())
			} else {
				ereader.Close()
				log.Debug().Str("file", ereader.Name).Msg("epub open")
				if debug {
					xmlj, _ := json.MarshalIndent(ereader.Container.Rootfiles[0], "", "    ")
					fmt.Println(string(xmlj))
				}
				log.Info().Str("title", ereader.Container.Rootfiles[0].Metadata.Title).
					Str("author", ereader.Container.Rootfiles[0].Metadata.Creator.Text).Msg("")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolP("debug", "d", false, "Print xml parsed to stdout")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
