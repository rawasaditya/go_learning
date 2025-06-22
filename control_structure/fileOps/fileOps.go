package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func GetFloatFromFile(fileName string) (acntBalance float64, err error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	balanceText := string(data)
	acntBalance, err = strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return
}
