package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "imgaug",
	Short: "pack 데이터를 바탕으로 이미지 augmentation 생성",
	Long: `이미지 사이즈, 반전, 변환 등을 쉽게 할 수 있는 CLI 툴입니다.
이 툴을 통해 다양한 이미지 증강(augment)가 가능합니다.`,

	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
