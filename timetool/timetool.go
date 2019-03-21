package timetool

import (
	"errors"
	"fmt"
	"github.com/fedir/toolkit/vartool"
	"strconv"
	"time"
)

func UnixTimestampToDate(timestamp interface{}) (string, error) {
	var err error

	var timestampInt int
	var timestampFormatted string
	timestampType := vartool.Type(timestamp)

	if timestampType != "int" && timestampType != "string" {
		return "", errors.New("At the moment only int and string unixtimestamps are supported")
	}

	if timestampType == "string" {
		timestampInt, err = strconv.Atoi(timestamp.(string))
		if err != nil {
			return "", errors.New("Sorry, couldn't understand Your string")
		}
	} else if timestampType == "int" {
		timestampInt = timestamp.(int)
	}

	unixTimeUTC := time.Unix(int64(timestampInt), 0) //gives unix time stamp in utc

	timestampFormatted = fmt.Sprintf("%d/%02d/%02d", unixTimeUTC.Year(), unixTimeUTC.Month(), unixTimeUTC.Day())

	return timestampFormatted, err
}
