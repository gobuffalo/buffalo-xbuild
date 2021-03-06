package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gobuffalo/buffalo-xbuild/build"
	"github.com/markbates/sigtx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var options = build.Options{}
var tags = ""
var debug bool

var xbuildCmd = &cobra.Command{
	Use:   "xbuild",
	Short: "Builds a Buffalo binary, including bundling of assets (packr & webpack) - experimental",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := sigtx.WithCancel(context.Background(), os.Interrupt)
		defer cancel()

		if options.Debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		b := build.New(ctx, options)
		if tags != "" {
			b.Tags = append(b.Tags, tags)
		}

		go func() {
			<-ctx.Done()
			if ctx.Err() == context.Canceled {
				fmt.Println("~~~BUILD CANCELLED ~~~")
				err := b.Cleanup()
				if err != nil {
					logrus.Fatal(err)
				}
			}
		}()

		err := b.Run()
		if err != nil {
			return errors.WithStack(err)
		}

		fmt.Printf("\nYou application was successfully built at %s\n", filepath.Join(b.Root, b.BinName))

		return nil
	},
}

func init() {
	RootCmd.AddCommand(xbuildCmd)

	pwd, _ := os.Getwd()
	options.Root = pwd
	output := filepath.Join("bin", filepath.Base(pwd))

	if runtime.GOOS == "windows" {
		output += ".exe"
	}

	xbuildCmd.Flags().StringVarP(&options.BinName, "output", "o", output, "set the name of the binary")
	xbuildCmd.Flags().StringVarP(&tags, "tags", "t", "", "compile with specific build tags")
	xbuildCmd.Flags().BoolVarP(&options.ExtractAssets, "extract-assets", "e", false, "extract the assets and put them in a distinct archive")
	xbuildCmd.Flags().BoolVarP(&options.Static, "static", "s", false, "build a static binary using  --ldflags '-linkmode external -extldflags \"-static\"'")
	xbuildCmd.Flags().StringVar(&options.LDFlags, "ldflags", "", "set any ldflags to be passed to the go build")
	xbuildCmd.Flags().BoolVarP(&options.Debug, "debug", "d", false, "print debugging information")
	xbuildCmd.Flags().BoolVarP(&options.Compress, "compress", "c", true, "compress static files in the binary")

}
