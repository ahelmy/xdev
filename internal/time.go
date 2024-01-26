package internal

import (
	"fmt"
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

	location, err := time.LoadLocation(*timeZone)
	if err != nil {
		return Time{}, err
	}
	return Time{UTC: t.UTC().Format(format), YourTimezone: t.In(location).Format(format), Epoch: t.Unix()}, nil
}

func ConvertTimeFromEpoch(epoch int64, format string) Time {
	t := time.Unix(epoch, 0)
	return Time{UTC: t.UTC().Format(format), YourTimezone: t.Format(format), Epoch: t.Unix()}
}

func ConvertTimeFromFormat(datetime string, fromFormat string, toFormat string) (Time, error) {
	t, err := time.ParseInLocation(fromFormat, datetime, time.Local)
	if err != nil {
		return Time{}, err
	}
	return Time{UTC: t.UTC().Format(toFormat), YourTimezone: t.Format(toFormat), Epoch: t.Unix()}, nil
}
