package cmd

import (
	"fmt"

	"errors"

	"github.com/apex/log"
	"github.com/evalphobia/go-timber/timber"
	"github.com/jeanmarcboite/truc/pkg/epub"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func someFunction() {
	ctx := log.WithFields(log.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "tobi",
	})

	ctx.Info("upload")
	ctx.Info("upload complete")
	ctx.Warn("upload retry")
	ctx.WithError(errors.New("unauthorized")).Error("upload failed")
	ctx.Errorf("failed to upload %s", "img.png")

	conf := timber.Config{
		APIKey:       "",
		SourceID:     "",
		Environment:  "development",
		MinimumLevel: timber.LogLevelInfo,
		Sync:         false,
		Debug:        true,
	}

	cli, _ := timber.New(conf)

	cli.Debug("logging...")
	cli.Trace("logging...")
	cli.Info("logging...")
	cli.Warn("logging...")
	cli.Err("logging...")
	cli.Fatal("logging...")
}

// Logger
var Logger *zap.SugaredLogger

func NewDevelopmentConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      false, // do not panic
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func NewDevelopment(options ...zap.Option) (*zap.Logger, error) {
	return NewDevelopmentConfig().Build(options...)
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
		logger, _ := NewDevelopment()
		Logger = logger.Sugar()

		someFunction()
		fmt.Println(args)
		ebook, error := epub.OpenReader(args[0])

		if error != nil {
			Logger.Errorf("%s: %s", args[0], error)
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
