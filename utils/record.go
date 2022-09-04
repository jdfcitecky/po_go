package utils

import (
	"fmt"

	"strconv"
	"time"
)

func GetCurrentDate() int {
	//save as file
	now := time.Now()
	stringTime := fmt.Sprintf("%s%s%s",
		fmt.Sprintf("%04d", now.Year()),
		fmt.Sprintf("%02d", now.Month()),
		fmt.Sprintf("%02d", now.Day()))
	i, err := strconv.Atoi(stringTime)
	if err != nil {
		return -1
	}
	return i
}
