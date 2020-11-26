package cmd

import (
	"encoding/json"
	"fmt"

	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/jeanmarcboite/librarytruc/pkg/books"
	"github.com/jeanmarcboite/librarytruc/pkg/books/epub"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/net"
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
		net.PrintKey()
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

		log.Debug().Str("args", fmt.Sprint(args)).Msg("lookup")

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

				ISBN, _ := ereader.GetISBN()
				w, _ := books.WorkFromISBN(ISBN)

				log.Info().Str("title", ereader.Container.Rootfiles[0].Metadata.Title).
					Str("ISBN", ISBN).
					Msg("")
				ws, _ := json.MarshalIndent(w, "", "  ")
				fmt.Printf(string(ws))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)
	lookupCmd.Flags().BoolP("debug", "d", false, "Print xml parsed to stdout")
}
