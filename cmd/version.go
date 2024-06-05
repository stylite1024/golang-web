package cmd

import (
	"fmt"
	"iris-web/config"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "print version info",
		Run: func(cmd *cobra.Command, args []string) {
			printVersionInfo()
		},
	}
)

func printVersionInfo() {
	info := fmt.Sprintf(`%s version information:
Version    :  %s
Go version :  %s
OS / Arch  :  %s/%s
`, config.AppName, config.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Println(string(config.LogoContent))
	fmt.Println(info)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
