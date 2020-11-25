package cmd

import (
	"fmt"

	"github.com/jeanmarcboite/truc/pkg/epub"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// Logger
var Logger *zap.SugaredLogger

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
		logger, _ := zap.NewProduction()
		Logger = logger.Sugar()

		fmt.Println(args)
		ebook, error := epub.OpenReader(args[0])

		if error != nil {
			Logger.Error(error)
		} else {
			Logger.Infof("ebook name: %s", ebook.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
