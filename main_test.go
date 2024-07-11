package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	parser := &StandardCronParser{}

	tests := []struct {
		cronString string
		expected   *CronFields
	}{
		{
			"*/15 0 1,15 * 1-5 /usr/bin/find",
			&CronFields{
				Minute:     []string{"0", "15", "30", "45"},
				Hour:       []string{"0"},
				DayOfMonth: []string{"1", "15"},
				Month:      generateRange(1, 12),
				DayOfWeek:  []string{"1", "2", "3", "4", "5"},
				Command:    "/usr/bin/find",
			},
		},
		{
			"0 12 * * 0 /usr/bin/backup",
			&CronFields{
				Minute:     []string{"0"},
				Hour:       []string{"12"},
				DayOfMonth: generateRange(1, 31),
				Month:      generateRange(1, 12),
				DayOfWeek:  []string{"0"},
				Command:    "/usr/bin/backup",
			},
		},
		{
			"5 4 * 1 * /path/to/script",
			&CronFields{
				Minute:     []string{"5"},
				Hour:       []string{"4"},
				DayOfMonth: generateRange(1, 31),
				Month:      []string{"1"},
				DayOfWeek:  generateRange(0, 6),
				Command:    "/path/to/script",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.cronString, func(t *testing.T) {
			parsedFields, err := parser.Parse(test.cronString)
			if err != nil {
				t.Fatalf("Parse returned error: %v", err)
			}
			if !reflect.DeepEqual(parsedFields, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, parsedFields)
			}
		})
	}
}

func TestParseInvalid(t *testing.T) {
	parser := &StandardCronParser{}

	invalidCronStrings := []string{
		"*/15 0 1,15 *",
		"invalid cron string",
		"* * * * * * *",
	}

	for _, cronString := range invalidCronStrings {
		t.Run(cronString, func(t *testing.T) {
			_, err := parser.Parse(cronString)
			if err == nil {
				t.Fatalf("Expected error for invalid cron string: %s", cronString)
			}
		})
	}
}
