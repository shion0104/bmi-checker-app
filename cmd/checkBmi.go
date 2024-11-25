package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var checkBMICmd = &cobra.Command{
	Use:   "check_bmi", 
	Short: "Measuring BMI",
	Long:  "Input height and weight and output BMI",
	Run: func(cmd *cobra.Command, args []string) {
		var stature int
		fmt.Print("身長を入力してください(cm): ")
		_, err := fmt.Scanln(&stature)
		if err != nil {
			fmt.Println("身長を正しく入力してください")
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