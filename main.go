package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CronParser interface for parsing cron strings
type CronParser interface {
	Parse(cronString string) (*CronFields, error)
}

// CronFields struct to hold parsed cron fields
type CronFields struct {
	Minute     []string
	Hour       []string
	DayOfMonth []string
	Month      []string
	DayOfWeek  []string
	Command    string
}

// StandardCronParser struct to implement CronParser
type StandardCronParser struct{}

// Parse method to parse the cron string
func (p *StandardCronParser) Parse(cronString string) (*CronFields, error) {
	fields := strings.Fields(cronString)
	if len(fields) != 6 {
		return nil, fmt.Errorf("invalid cron string")
	}

	parsedFields := &CronFields{
		Minute:     parseField(fields[0], 0, 59),
		Hour:       parseField(fields[1], 0, 23),
		DayOfMonth: parseField(fields[2], 1, 31),
		Month:      parseField(fields[3], 1, 12),
		DayOfWeek:  parseField(fields[4], 0, 6),
		Command:    fields[5],
	}

	return parsedFields, nil
}

// Helper function to parse individual cron fields
func parseField(field string, min, max int) []string {
	if field == "*" {
		return generateRange(min, max)
	}
	if strings.Contains(field, "*/") {
		step, _ := strconv.Atoi(strings.TrimPrefix(field, "*/"))
		return generateStepRange(min, max, step)
	}
	if strings.Contains(field, ",") {
		return strings.Split(field, ",")
	}
	if strings.Contains(field, "-") {
		rangeParts := strings.Split(field, "-")
		start, _ := strconv.Atoi(rangeParts[0])
		end, _ := strconv.Atoi(rangeParts[1])
		return generateRange(start, end)
	}
	return []string{field}
}

// Helper function to generate a range of numbers
func generateRange(start, end int) []string {
	var result []string
	for i := start; i <= end; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result
}

// Helper function to generate a step range of numbers
func generateStepRange(start, end, step int) []string {
	var result []string
	for i := start; i <= end; i += step {
		result = append(result, strconv.Itoa(i))
	}
	return result
}

// Main function
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main \"<cron_string>\"")
		os.Exit(1)
	}

	cronString := os.Args[1]
	parser := &StandardCronParser{}
	parsedFields, err := parser.Parse(cronString)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	printParsedFields(parsedFields)
}

// Helper function to print parsed fields in table format
func printParsedFields(fields *CronFields) {
	fmt.Printf("%-14s%s\n", "minute", strings.Join(fields.Minute, " "))
	fmt.Printf("%-14s%s\n", "hour", strings.Join(fields.Hour, " "))
	fmt.Printf("%-14s%s\n", "day of month", strings.Join(fields.DayOfMonth, " "))
	fmt.Printf("%-14s%s\n", "month", strings.Join(fields.Month, " "))
	fmt.Printf("%-14s%s\n", "day of week", strings.Join(fields.DayOfWeek, " "))
	fmt.Printf("%-14s%s\n", "command", fields.Command)
}
