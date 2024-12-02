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

		bodyWeight, err := input.GetWeight()
		if err != nil {
			fmt.Println("エラー:", err)
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