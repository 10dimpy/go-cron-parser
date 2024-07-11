package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	parser := &CronParserStruct{}

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
	parser := &CronParserStruct{}

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

func TestParseField(t *testing.T) {
	tests := []struct {
		field string
		min   int
		max   int
		want  []string
	}{
		{"*", 0, 59, generateRange(0, 59)},
		{"*/15", 0, 59, generateStepRange(0, 59, 15)},
		{"1,15", 1, 31, []string{"1", "15"}},
		{"1-5", 0, 6, generateRange(1, 5)},
		{"0", 0, 23, []string{"0"}},
	}

	for _, tt := range tests {
		t.Run(tt.field, func(t *testing.T) {
			if got := parseField(tt.field, tt.min, tt.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseField(%v, %v, %v) = %v, want %v", tt.field, tt.min, tt.max, got, tt.want)
			}
		})
	}
}

func TestGenerateRange(t *testing.T) {
	start := 1
	end := 5
	expected := []string{"1", "2", "3", "4", "5"}

	result := generateRange(start, end)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGenerateStepRange(t *testing.T) {
	start := 0
	end := 30
	step := 10
	expected := []string{"0", "10", "20", "30"}

	result := generateStepRange(start, end, step)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
