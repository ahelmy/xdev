package internal

import (
	"fmt"
	"strings"
	"time"
)

const (
	DateFormat = "02-01-2006 15:04:05"
)

func parseFormat(format string) string {
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
	UTC          string
	YourTimezone string
	Epoch        int64
}

func (t Time) String() string {
	return fmt.Sprintf("UTC: %s\nYour Timezone: %s\nEpoch: %d", t.UTC, t.YourTimezone, t.Epoch)
}

func Now(format string) Time {
	format = parseFormat(format)

	t := time.Now()
	return Time{UTC: t.UTC().Format(format), YourTimezone: t.Format(format), Epoch: t.Unix()}
}

func ConvertTimeFromEpoch(epoch int64, format string) Time {
	format = parseFormat(format)
	t := time.Unix(epoch, 0)
	return Time{UTC: t.UTC().Format(format), YourTimezone: t.Format(format), Epoch: t.Unix()}
}

func ConvertTimeFromFormat(datetime string, fromFormat string, toFormat string) (Time, error) {
	toFormat = parseFormat(toFormat)
	fromFormat = parseFormat(fromFormat)

	t, err := time.Parse(fromFormat, datetime)
	if err != nil {
		return Time{}, err
	}
	return Time{UTC: t.UTC().Format(toFormat), YourTimezone: t.Format(toFormat), Epoch: t.Unix()}, nil
}
