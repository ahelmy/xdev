package internal

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"testing"
)

func TestHashString(t *testing.T) {
	testCases := []struct {
		str       string
		algorithm string
		expected  string
	}{
		{
			str:       "Hello, World!",
			algorithm: "md5",
			expected:  fmt.Sprintf("%x", md5.Sum([]byte("Hello, World!"))),
		},
		{
			str:       "Hello, World!",
			algorithm: "sha256",
			expected:  fmt.Sprintf("%x", sha256.Sum256([]byte("Hello, World!"))),
		},
		{
			str:       "Hello, World!",
			algorithm: "sha512",
			expected:  fmt.Sprintf("%x", sha512.Sum512([]byte("Hello, World!"))),
		},

		{
			str:       "Hello, World!",
			algorithm: "invalid",
			expected:  "invalid algorithm",
		},
	}

	for _, tc := range testCases {
		result, err := HashString(tc.str, tc.algorithm)
		if err != nil {
			if err.Error() != tc.expected {
				t.Errorf("Expected error: %s, but got: %s", tc.expected, err.Error())
			}
		} else {
			if result != tc.expected {
				t.Errorf("Expected result: %s, but got: %s", tc.expected, result)
			}
		}
	}
}

func TestSalt(t *testing.T) {
	testCases := []struct {
		str     string
		length  int
		isError bool
	}{
		{
			str:     "Hello, World!",
			length:  60,
			isError: false,
		},
		{
			str:     "$2a$10$2dIj/VAy0Zhgy6eaNYjfAubIbyP5z2V7e8Qzhyr/xeo56GtQ22kOG$2a$10$2dIj/VAy0Zhgy6eaNYjfAubIbyP5z2V7e8Qzhyr/xeo56GtQ22kOG$2a$10$2dIj/VAy0Zhgy6eaNYjfAubIbyP5z2V7e8Qzhyr/xeo56GtQ22kOG$2a$10$2dIj/VAy0Zhgy6eaNYjfAubIbyP5z2V7e8Qzhyr/xeo56GtQ22kOG$2a$10$2dIj/VAy0Zhgy6eaNYjfAubIbyP5z2V7e8Qzhyr/xeo56GtQ22kOG$2a$10$2dIj/VAy0Zhgy6eaNYjfAubIbyP5z2V7e8Qzhyr/xeo56GtQ22kOG$2a$10$2dIj/VAy0Zhgy6eaNYjfAubIbyP5z2V7e8Qzhyr/xeo56GtQ22kOG$2a$10$2dIj/VAy0Zhgy6eaNYjfAubIbyP5z2V7e8Qzhyr/xeo56GtQ22kOG",
			isError: true,
		},
	}

	for _, tc := range testCases {
		result, err := HashString(tc.str, "salt")
		if err != nil && !tc.isError {
			t.Errorf("Unexpected error: %v", err)
		} else {
			if len(result) != tc.length {
				t.Errorf("Expected result: %v, but got: %v", tc.length, len(result))
			}
		}
	}
}
