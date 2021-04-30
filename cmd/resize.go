package cmd

import (
	"github.com/spf13/cobra"
	"github.com/insightercorperation/image-augment/usecase"
)

var (
	parentDirs    []string
	target        string
	size          int
	remainPolygon bool
	resizeCmd     = &cobra.Command{
		Use:   "resize",
		Short: "resize image and label on pack",
	}
)

func init() {

	rootCmd.AddCommand(resizeCmd)

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create resized image and label data",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			for _, parentDir := range parentDirs {
				usecase.Resize(parentDir, target, size, remainPolygon)
			}
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete resized image and label data",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	resizeCmd.PersistentFlags().StringVar(&target, "target", "all", "target to resize. allows [image | label | all]")
	resizeCmd.PersistentFlags().StringArrayVar(&parentDirs, "parentDirs", make([]string, 1), "parent directories")
	resizeCmd.PersistentFlags().IntVar(&size, "size", 416, "size for (width, height)")
	resizeCmd.PersistentFlags().BoolVarP(&remainPolygon, "polygon", "p", false, "Flag not to convert polygon to bbox")

	resizeCmd.AddCommand(createCmd)
	resizeCmd.AddCommand(deleteCmd)
}
