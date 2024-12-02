package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-qr-app/bmi"
	"go-qr-app/input"
	"os"
)

var checkBMICmd = &cobra.Command{
	Use:   "check_bmi", 
	Short: "Measuring BMI",
	Long:  "Input height and weight and output BMI",
	Run: func(cmd *cobra.Command, args []string) {

		stature, err := input.GetStature()
		error_check(err)

		bodyWeight, err := input.GetWeight()
		error_check(err)

		bmi, err := bmi.CalculateBMI(bodyWeight, stature)
		error_check(err)

		fmt.Printf("あなたのBMIは %.2f です。\n", bmi)
	},
}

func init() {
	rootCmd.AddCommand(checkBMICmd)
}

func error_check(err error) {
	if err != nil {
			fmt.Println(err)
			os.Exit(1) 
		 }
}