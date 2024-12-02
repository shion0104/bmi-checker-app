package bmi

import(
	"fmt"
)

func CalculateBMI(bodyWeight int, stature int) (float64, error) {

	heightInMeters := float64(stature) / 100.0

	if heightInMeters <= 0 {
		return 0, fmt.Errorf("無効な身長です")
	}

	bmi := float64(bodyWeight) / (heightInMeters * heightInMeters)
	return bmi, nil
}