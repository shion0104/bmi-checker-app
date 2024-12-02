package input

import (
	"fmt"
)

func GetWeight() (int, error) {
	var weight int
	fmt.Print("体重を入力してください(kg)")
	_, err := fmt.Scanln(&weight)
	if err != nil || weight <= 0 {
		return 0, fmt.Errorf("無効な身長です。正しい値を入力してください")
	}
	return weight, nil
}