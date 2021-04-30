package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/insightercorperation/image-augment/usecase"
)

var (
	trainingRatio   int
	validationRatio int
	testRatio       int
	hasSample       bool
	outputDir       string
	segmentCmd      = &cobra.Command{
		Use:   "segment",
		Short: "create data segment",
	}
)

func init() {
	rootCmd.AddCommand(segmentCmd)

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create segment format annotation on resized",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			for idx, parentDir := range parentDirs {
				fmt.Println(idx, parentDir)
				usecase.Segment(parentDir, outputDir, trainingRatio, validationRatio, testRatio, hasSample)
			}
		},
	}

	segmentCmd.PersistentFlags().StringArrayVar(&parentDirs, "parentDirs", make([]string, 2), "parent directories")
	segmentCmd.PersistentFlags().StringVar(&outputDir, "outputDir", "./segment", "output directories")
	segmentCmd.PersistentFlags().IntVar(&trainingRatio, "training", 8, "ratio of training data")
	segmentCmd.PersistentFlags().IntVar(&validationRatio, "validation", 1, "ratio of validation data")
	segmentCmd.PersistentFlags().IntVar(&testRatio, "test", 1, "ratio of test data")
	segmentCmd.PersistentFlags().BoolVarP(&hasSample, "sample", "s", false, "Flag to create sample segment")
	segmentCmd.AddCommand(createCmd)
}
