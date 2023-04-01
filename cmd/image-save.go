package cmd

import (
	"github.com/DockerContainerService/image-save/pkg/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	version, osFilter, archFilter, username, password, output string
	debug, insecure                                           bool
)

var rootCmd = &cobra.Command{
	Use:   "imsave [image] [flags]",
	Short: "Dockerlessed image save tool",
	Long: `Save docker image to local without docker daemon
	Complete documentation is available at https://github.com/DockerContainerService/image-save`,
	Version: version,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		c := client.NewClient(args[0], username, password, insecure)
		var osFilters []string

		// Avoid empty osFilter
		if osFilter == "" {
			osFilters = []string{}
		} else {
			osFilters = []string{osFilter}
		}

		c.Save(osFilters, []string{archFilter}, output)
	},
}

func init() {
	rootCmd.SetVersionTemplate("imsave version {{.Version}}\n")
	rootCmd.PersistentFlags().StringVar(&archFilter, "arch", runtime.GOARCH, "The architecture of the image you want to save")
	rootCmd.PersistentFlags().StringVar(&osFilter, "os", "", "The osFilter of the image you want to save")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output file")
	rootCmd.PersistentFlags().StringVarP(&username, "user", "u", "", "Username of the registry")
	rootCmd.PersistentFlags().StringVarP(&password, "passwd", "p", "", "Password of the registry")
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "Whether the registry is using http")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug mode")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
