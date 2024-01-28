package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	DateFormat   = "02-01-2006 15:04:05"
	ToDateFormat = "Monday, January 02, 2006 3:04:05 PM"
)

func ParseFormat(format string) string {
	if format == "" {
		return DateFormat
	}
	if format == DateFormat {
		return format
	}
	format = strings.ReplaceAll(format, "dd", "02")
	format = strings.ReplaceAll(format, "MM", "01")
	format = strings.ReplaceAll(format, "yyyy", "2006")
	format = strings.ReplaceAll(format, "HH", "15")
	format = strings.ReplaceAll(format, "mm", "04")
	format = strings.ReplaceAll(format, "ss", "05")

	format = strings.ReplaceAll(format, "d", "2")
	format = strings.ReplaceAll(format, "M", "1")
	format = strings.ReplaceAll(format, "H", "3")
	format = strings.ReplaceAll(format, "m", "4")
	format = strings.ReplaceAll(format, "s", "5")
	format = strings.ReplaceAll(format, "yy", "06")
	format = strings.ReplaceAll(format, "Z", "MST")

	return format
}

type Time struct {
	UTC          string `json:"utc"`
	YourTimezone string `json:"your_timezone"`
	Epoch        int64  `json:"epoch"`
}

func (t Time) String() string {
	return fmt.Sprintf("UTC: %s\nYour Timezone: %s\nEpoch: %d", t.UTC, t.YourTimezone, t.Epoch)
}

func Now(format string, timeZone *string) (Time, error) {
	t := time.Now()

	if timeZone == nil {
		return Time{UTC: t.UTC().Format(format), YourTimezone: t.Format(format), Epoch: t.Unix()}, nil
	}

	var yourTZ = getCurrentTimeFromUTCDifferent(t, *timeZone).Format(format)

	/*	location, err := time.LoadLocation(*timeZone)
		if err != nil {
			yourTZ = t.Format(format)
		} else {
			yourTZ = t.In(location).Format(format)
		}*/

	return Time{UTC: t.UTC().Format(format), YourTimezone: yourTZ, Epoch: t.Unix()}, nil
}

func ConvertTimeFromEpoch(epoch int64, format string, timeZone *string) Time {
	t := time.Unix(epoch, 0)

	if timeZone == nil {
		return Time{UTC: t.UTC().Format(format), YourTimezone: t.Format(format), Epoch: t.Unix()}
	}

	var yourTZ = getCurrentTimeFromUTCDifferent(t, *timeZone).Format(format)

	/*	location, err := time.LoadLocation(*timeZone)
		if err != nil {
			yourTZ = t.Format(format)
		} else {
			yourTZ = t.In(location).Format(format)
		}*/

	return Time{UTC: t.UTC().Format(format), YourTimezone: yourTZ, Epoch: t.Unix()}
}

func ConvertTimeFromFormat(datetime string, fromFormat string, toFormat string, timeZone *string) (Time, error) {
	t, err := time.ParseInLocation(fromFormat, datetime, time.Local)
	if err != nil {
		return Time{}, err
	}

	if timeZone == nil {
		return Time{UTC: t.UTC().Format(toFormat), YourTimezone: t.Format(toFormat), Epoch: t.Unix()}, nil
	}

	var yourTZ = getCurrentTimeFromUTCDifferent(t, *timeZone).Format(toFormat)

	/*	location, err := time.LoadLocation(*timeZone)
		if err != nil {
			yourTZ = t.Format(toFormat)
		} else {
			yourTZ = t.In(location).Format(toFormat)
		}*/

	return Time{UTC: t.UTC().Format(toFormat), YourTimezone: yourTZ, Epoch: t.Unix()}, nil
}

func getCurrentTimeFromUTCDifferent(utcTime time.Time, timeZone string) time.Time {
	timeZone = timeZone[3:len(timeZone)] // remove UTC From String
	var operator = timeZone[:1]
	var hours, _ = strconv.Atoi(timeZone[1:3])
	var minutes, _ = strconv.Atoi(timeZone[4:6])

	if operator == "+" {
		utcTime = utcTime.UTC().Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(minutes))
	} else {
		utcTime = utcTime.UTC().Add((time.Hour*time.Duration(hours) + time.Minute*time.Duration(minutes)) * -1)
	}

	return utcTime
}
