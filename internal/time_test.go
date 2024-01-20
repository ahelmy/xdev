package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	format := "01-02-2006 15:04:05" // Replace with your desired format

	expectedUTC := time.Now().UTC().Format(format)
	expectedYourTimezone := time.Now().Format(format)
	expectedEpoch := time.Now().Unix()

	result := Now(format)

	// Check the UTC time
	if result.UTC != expectedUTC {
		t.Errorf("Expected UTC time %s, but got %s", expectedUTC, result.UTC)
	}

	// Check the time in your timezone
	if result.YourTimezone != expectedYourTimezone {
		t.Errorf("Expected time in your timezone %s, but got %s", expectedYourTimezone, result.YourTimezone)
	}

	// Check the epoch value
	if result.Epoch != expectedEpoch {
		t.Errorf("Expected epoch value %d, but got %d", expectedEpoch, result.Epoch)
	}
}

func TestConvertTimeFromEpoch(t *testing.T) {
	epoch := int64(1625097600) // Replace with your desired epoch value
	format := ""               // Replace with your desired format

	expectedUTC := "01-07-2021 00:00:00"          // Replace with your expected UTC time
	expectedYourTimezone := "01-07-2021 02:00:00" // Replace with your expected time in your timezone
	expectedEpoch := int64(1625097600)            // Replace with your expected epoch value

	result := ConvertTimeFromEpoch(epoch, format)

	// Check the UTC time
	if result.UTC != expectedUTC {
		t.Errorf("Expected UTC time %s, but got %s", expectedUTC, result.UTC)
	}

	// Check the time in your timezone
	if result.YourTimezone != expectedYourTimezone {
		t.Errorf("Expected time in your timezone %s, but got %s", expectedYourTimezone, result.YourTimezone)
	}

	// Check the epoch value
	if result.Epoch != expectedEpoch {
		t.Errorf("Expected epoch value %d, but got %d", expectedEpoch, result.Epoch)
	}
}
func TestConvertTimeFromFormat(t *testing.T) {
	datetime := "2022-01-01 12:00:00"
	fromFormat := "2006-01-02 15:04:05"
	toFormat := "01-02-2006 15:04:05"

	expectedUTC := "01-01-2022 12:00:00"
	expectedYourTimezone := "01-01-2022 12:00:00"
	expectedEpoch := int64(1641038400)

	result, err := ConvertTimeFromFormat(datetime, fromFormat, toFormat)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check the UTC time
	if result.UTC != expectedUTC {
		t.Errorf("Expected UTC time %s, but got %s", expectedUTC, result.UTC)
	}

	// Check the time in your timezone
	if result.YourTimezone != expectedYourTimezone {
		t.Errorf("Expected time in your timezone %s, but got %s", expectedYourTimezone, result.YourTimezone)
	}

	// Check the epoch value
	if result.Epoch != expectedEpoch {
		t.Errorf("Expected epoch value %d, but got %d", expectedEpoch, result.Epoch)
	}

	result, err = ConvertTimeFromFormat("xyz", "123", toFormat)
	if err == nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
func TestParseFormat(t *testing.T) {
	tests := []struct {
		format         string
		expectedFormat string
	}{
		{
			format:         "",
			expectedFormat: DateFormat,
		},
		{
			format:         DateFormat,
			expectedFormat: DateFormat,
		},
		{
			format:         "dd-MM-yyyy HH:mm:ss",
			expectedFormat: "02-01-2006 15:04:05",
		},
		{
			format:         "yyyy/MM/dd",
			expectedFormat: "2006/01/02",
		},
		{
			format:         "HH:mm",
			expectedFormat: "15:04",
		},
		{
			format:         "ss",
			expectedFormat: "05",
		},
		{
			format:         "yy-MM-dd",
			expectedFormat: "06-01-02",
		},
		{
			format:         "Z",
			expectedFormat: "MST",
		},
	}

	for _, test := range tests {
		result := parseFormat(test.format)
		if result != test.expectedFormat {
			t.Errorf("Expected format %s, but got %s", test.expectedFormat, result)
		}
	}
}

func TestTimeString(t *testing.T) {
	format := "01-02-2006 15:04:05" // Replace with your desired format

	expectedUTC := time.Now().UTC().Format(format)
	expectedYourTimezone := time.Now().Format(format)
	expectedEpoch := time.Now().Unix()

	timeObj := Time{
		UTC:          expectedUTC,
		YourTimezone: expectedYourTimezone,
		Epoch:        expectedEpoch,
	}

	expectedString := fmt.Sprintf("UTC: %s\nYour Timezone: %s\nEpoch: %d", expectedUTC, expectedYourTimezone, expectedEpoch)
	result := timeObj.String()

	if result != expectedString {
		t.Errorf("Expected string %s, but got %s", expectedString, result)
	}
}