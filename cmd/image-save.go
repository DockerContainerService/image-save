package cmd

import (
	"github.com/DockerContainerService/image-save/pkg/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	version, os, arch, username, password, output string
	debug, insecure                               bool
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
		var osFilter []string
		if os == "" {
			osFilter = []string{}
		} else {
			osFilter = []string{os}
		}
		c.Save(osFilter, []string{arch}, output)
	},
}

func init() {
	rootCmd.SetVersionTemplate("imsave version {{.Version}}\n")
	rootCmd.PersistentFlags().StringVar(&arch, "arch", runtime.GOARCH, "The architecture of the image you want to save")
	rootCmd.PersistentFlags().StringVar(&os, "os", "", "The os of the image you want to save")
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