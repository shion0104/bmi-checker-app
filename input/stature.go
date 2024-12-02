package stature

import (
	"fmt"
)

func GetStature() (int, error) {
	var stature int
	fmt.Print("身長を入力してください(cm)")
	_, err := fmt.Scanln(&stature)
	if err != nil || stature <= 0 {
		return 0, fmt.Errorf("無効な身長です。正しい値を入力してください")
	}
	return stature, nil
}