package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version    = "1.0.0"
	poo        string
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of imgaug",
		Long:  `All software has versions. This is imgaug's`,
		Run: func(cmd *cobra.Command, args []string) {
			versionStr := fmt.Sprintf("imgaug v%s", version)
			fmt.Println(versionStr)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
