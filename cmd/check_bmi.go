package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-qr-app/input"
)

var checkBMICmd = &cobra.Command{
	Use:   "check_bmi", 
	Short: "Measuring BMI",
	Long:  "Input height and weight and output BMI",
	Run: func(cmd *cobra.Command, args []string) {
		stature, err := input.GetStature()
		if err != nil {
			fmt.Println("エラー:", err)
			return
		}

		var bodyWeight int
		fmt.Print("体重を入力してください(kg): ")
		_, err = fmt.Scanln(&bodyWeight)
		if err != nil {
			fmt.Println("体重を正しく入力してください。")
			return
		}

		heightInMeters := float64(stature) / 100.0

		bmi := float64(bodyWeight) / (heightInMeters * heightInMeters)

		fmt.Printf("あなたのBMIは %.2f です。\n", bmi)
	},
}

func init() {
	rootCmd.AddCommand(checkBMICmd)
}