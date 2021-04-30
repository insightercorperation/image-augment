package cmd

import (
	"github.com/spf13/cobra"
	"github.com/insightercorperation/image-augment/usecase"
)

var (
	cropCmd = &cobra.Command{
		Use:   "crop",
		Short: "crop image",
	}
)

func init() {
	rootCmd.AddCommand(cropCmd)

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create crop image on resized",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			for _, parentDir := range parentDirs {
				usecase.Crop(parentDir, size)
			}
		},
	}

	cropCmd.PersistentFlags().StringArrayVar(&parentDirs, "parentDirs", make([]string, 2), "parent directories")
	cropCmd.PersistentFlags().IntVar(&size, "size", 416, "size for (width, height)")
	cropCmd.AddCommand(createCmd)
}
