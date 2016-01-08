package utils

import (
	"math"
	"strconv"
	"time"
)

type IUtils interface {
	IntArrayToString(array []int) string
	GetTimeDimensionValues() (year_number,
		month_number,
		month_day_number,
		week_number,
		quarter_number int)
}
type Utils struct {
	IUtils
}

func GetUtils() *Utils {
	var utils = Utils{}
	return &utils
}

// Gets a int array and returns string.
func (utils *Utils) IntArrayToString(array []int) string {
	var str string
	for _, value := range array {
		if str == "" {
			str += strconv.Itoa(value)
		} else {
			str += "," + strconv.Itoa(value)
		}
	}
	return str
}

//returns the year, month, day, week and  quarter for time.Now().Date()
func (utils *Utils) GetTimeDimensionValues() (
	year_number,
	month_number,
	month_day_number,
	week_number,
	quarter_number int,
) {
	year_number, month, month_day_number := time.Now().Date()
	_, week_number = time.Now().ISOWeek()
	month_number = int(month)
	quarter_number = int(math.Ceil(float64(month_number) / 3))

	return year_number, month_number, month_day_number, week_number, quarter_number
}
