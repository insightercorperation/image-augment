package cmd

import (
	"github.com/spf13/cobra"
	"github.com/insightercorperation/image-augment/usecase"
)

var (
	polygon    bool
	categories []string
	cocoCmd    = &cobra.Command{
		Use:   "coco",
		Short: "create coco format annotation",
	}
)

func init() {
	rootCmd.AddCommand(cocoCmd)

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create coco format annotation on resized",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			for idx, parentDir := range parentDirs {
				usecase.Coco(parentDir, size, categories[idx], polygon)
			}
		},
	}
	cocoCmd.PersistentFlags().StringArrayVar(&parentDirs, "parentDirs", make([]string, 2), "parent directories")
	cocoCmd.PersistentFlags().StringArrayVar(&categories, "categories", make([]string, 2), "coco categories")
	cocoCmd.PersistentFlags().IntVar(&size, "size", 416, "size for (width, height)")
	cocoCmd.PersistentFlags().BoolVarP(&polygon, "polygon", "p", false, "Flag to create coco annotation on polygon")
	cocoCmd.AddCommand(createCmd)

}
