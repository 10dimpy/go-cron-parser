# Cron Parser

This is a command-line application written in Go that parses a cron string and expands each field to show the times at which it will run. It only considers the standard cron format with five time fields (minute, hour, day of month, month, and day of week) plus a command.

## Usage

To use this application, you need to pass a cron string as a single argument. The cron string should follow the standard format. The application will output the expanded times for each field.

### Example

```bash
./main "*/15 0 1,15 * 1-5 /usr/bin/find"
```
