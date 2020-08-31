package unit

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
)

func TestParseDuration(t *testing.T) {
	t.Parallel()
	t.Run("second", func(t *testing.T) {
		testCases := []struct {
			name             string
			time             string
			expectedHasError bool
			expectedDuration time.Duration
		}{
			{
				name:             "correct format",
				time:             "3s",
				expectedHasError: false,
				expectedDuration: 3 * time.Second,
			},
			{
				name:             "multiple digits",
				time:             "120s",
				expectedHasError: false,
				expectedDuration: 120 * time.Second,
			},
			{
				name:             "empty string",
				time:             "",
				expectedHasError: true,
				expectedDuration: 0,
			},
			{
				name:             "incorrect format",
				time:             "3ss",
				expectedHasError: true,
				expectedDuration: 0,
			},
			{
				name:             "duration with no value",
				time:             "s",
				expectedHasError: true,
				expectedDuration: 0,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				duration, err := ParseDuration(testCase.time)
				if testCase.expectedHasError {
					assert.NotEqual(t, nil, err)
					return
				}
				assert.Equal(t, nil, err)
				assert.Equal(t, testCase.expectedDuration, duration)
			})
		}
	})

	t.Run("minute", func(t *testing.T) {
		testCases := []struct {
			name             string
			time             string
			expectedHasError bool
			expectedDuration time.Duration
		}{
			{
				name:             "correct format",
				time:             "5m",
				expectedHasError: false,
				expectedDuration: 5 * time.Minute,
			},
			{
				name:             "multiple digits",
				time:             "150m",
				expectedHasError: false,
				expectedDuration: 150 * time.Minute,
			},
			{
				name:             "empty string",
				time:             "",
				expectedHasError: true,
				expectedDuration: 0,
			},
			{
				name:             "incorrect format",
				time:             "m5",
				expectedHasError: true,
				expectedDuration: 0,
			},
			{
				name:             "duration with no value",
				time:             "m",
				expectedHasError: true,
				expectedDuration: 0,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				duration, err := ParseDuration(testCase.time)
				if testCase.expectedHasError {
					assert.NotEqual(t, nil, err)
					return
				}
				assert.Equal(t, nil, err)
				assert.Equal(t, testCase.expectedDuration, duration)
			})
		}
	})

	t.Run("hour", func(t *testing.T) {
		testCases := []struct {
			name             string
			time             string
			expectedHasError bool
			expectedDuration time.Duration
		}{
			{
				name:             "correct format",
				time:             "6h",
				expectedHasError: false,
				expectedDuration: 6 * time.Hour,
			},
			{
				name:             "multiple digits",
				time:             "100h",
				expectedHasError: false,
				expectedDuration: 100 * time.Hour,
			},
			{
				name:             "empty string",
				time:             "",
				expectedHasError: true,
				expectedDuration: 0,
			},
			{
				name:             "incorrect format",
				time:             "6sh",
				expectedHasError: true,
				expectedDuration: 0,
			},
			{
				name:             "duration with no value",
				time:             "h",
				expectedHasError: true,
				expectedDuration: 0,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				duration, err := ParseDuration(testCase.time)
				if testCase.expectedHasError {
					assert.NotEqual(t, nil, err)
					return
				}
				assert.Equal(t, nil, err)
				assert.Equal(t, testCase.expectedDuration, duration)
			})
		}
	})

	t.Run("day", func(t *testing.T) {
		testCases := []struct {
			name             string
			time             string
			expectedHasError bool
			expectedDuration time.Duration
		}{
			{
				name:             "correct format",
				time:             "2d",
				expectedHasError: false,
				expectedDuration: 2 * oneDay,
			},
			{
				name:             "multiple digits",
				time:             "201d",
				expectedHasError: false,
				expectedDuration: 201 * oneDay,
			},
			{
				name:             "empty string",
				time:             "",
				expectedHasError: true,
				expectedDuration: 0,
			},
			{
				name:             "incorrect format",
				time:             "d",
				expectedHasError: true,
				expectedDuration: 0,
			},
			{
				name:             "duration with no value",
				time:             "d",
				expectedHasError: true,
				expectedDuration: 0,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				duration, err := ParseDuration(testCase.time)
				if testCase.expectedHasError {
					assert.NotEqual(t, nil, err)
					return
				}
				assert.Equal(t, nil, err)
				assert.Equal(t, testCase.expectedDuration, duration)
			})
		}
	})

	t.Run("week", func(t *testing.T) {
		testCases := []struct {
			name             string
			time             string
			expectedHasError bool
			expectedDuration time.Duration
		}{
			{
				name:             "correct format",
				time:             "1w",
				expectedHasError: false,
				expectedDuration: oneWeak,
			},
			{
				name:             "multiple digits",
				time:             "111w",
				expectedHasError: false,
				expectedDuration: 111 * oneWeak,
			},
			{
				name:             "empty string",
				time:             "",
				expectedHasError: true,
				expectedDuration: 0,
			},
			{
				name:             "incorrect format",
				time:             "w2w",
				expectedHasError: true,
				expectedDuration: 0,
			},
			{
				name:             "duration with no value",
				time:             "w",
				expectedHasError: true,
				expectedDuration: 0,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				duration, err := ParseDuration(testCase.time)
				if testCase.expectedHasError {
					assert.NotEqual(t, nil, err)
					return
				}
				assert.Equal(t, nil, err)
				assert.Equal(t, testCase.expectedDuration, duration)
			})
		}
	})
}
