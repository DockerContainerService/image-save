package cmd

import (
	"github.com/DockerContainerService/image-save/pkg/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	version, osFilter, archFilter, username, password, output, mirror string
	debug, insecure                                                   bool
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

		c := client.NewClient(args[0], username, password, mirror, insecure)
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
	rootCmd.PersistentFlags().StringVar(&archFilter, "arch", runtime.GOARCH, "the architecture of the image you want to save")
	rootCmd.PersistentFlags().StringVar(&osFilter, "os", "", "the osFilter of the image you want to save")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "output file")
	rootCmd.PersistentFlags().StringVarP(&username, "user", "u", "", "username of the registry")
	rootCmd.PersistentFlags().StringVarP(&password, "passwd", "p", "", "password of the registry")
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "i", false, "whether the registry is using http")
	rootCmd.PersistentFlags().StringVarP(&mirror, "mirror", "m", "registry.hub.docker.com", "use a mirror repository")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
