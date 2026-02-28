package builtins

import (
	"fmt"
	"strings"
	"time"

	"github.com/iscarloscoder/fig/environment"
)

// figDateFormat converts a Fig-style format string (YYYY-MM-DD HH:mm:ss)
// to a Go time layout string.
func figDateFormat(format string) string {
	r := strings.NewReplacer(
		"YYYY", "2006",
		"MM", "01",
		"DD", "02",
		"HH", "15",
		"mm", "04",
		"ss", "05",
	)
	return r.Replace(format)
}

// unitToDuration converts a Fig time unit string and quantity to a time.Duration.
func unitToDuration(qty float64, unit string) (time.Duration, error) {
	switch unit {
	case "ms", "millisecond", "milliseconds":
		return time.Duration(qty) * time.Millisecond, nil
	case "s", "second", "seconds":
		return time.Duration(qty) * time.Second, nil
	case "min", "minute", "minutes":
		return time.Duration(qty) * time.Minute, nil
	case "h", "hour", "hours":
		return time.Duration(qty) * time.Hour, nil
	case "day", "days":
		return time.Duration(qty) * 24 * time.Hour, nil
	case "week", "weeks":
		return time.Duration(qty) * 7 * 24 * time.Hour, nil
	default:
		return 0, fmt.Errorf("unknown time unit: %s", unit)
	}
}

func init() {
	register(newModule("date",
		// now() — returns current Unix timestamp in milliseconds
		fn("now", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("now() expects 0 arguments, got %d", len(args))
			}
			ms := float64(time.Now().UnixMilli())
			return environment.NewNumber(ms), nil
		}),

		// format(timestamp, format) — formats a timestamp to a string
		fn("format", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("format() expects 2 arguments, got %d", len(args))
			}
			ts, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("format() first argument must be a number (timestamp)")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("format() second argument must be a string (format)")
			}
			t := time.UnixMilli(int64(ts))
			layout := figDateFormat(args[1].Str)
			return environment.NewString(t.Format(layout)), nil
		}),

		// parse(dateString, format) — parses a date string into a timestamp (ms)
		fn("parse", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("parse() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("parse() first argument must be a string")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("parse() second argument must be a string (format)")
			}
			layout := figDateFormat(args[1].Str)
			t, err := time.ParseInLocation(layout, args[0].Str, time.Local)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("parse() cannot parse %q with format %q: %v", args[0].Str, args[1].Str, err)
			}
			return environment.NewNumber(float64(t.UnixMilli())), nil
		}),

		// add(timestamp, quantity, unit) — adds time to a timestamp
		fn("add", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("add() expects 3 arguments, got %d", len(args))
			}
			ts, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("add() first argument must be a number (timestamp)")
			}
			qty, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("add() second argument must be a number (quantity)")
			}
			if args[2].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("add() third argument must be a string (unit)")
			}
			dur, err := unitToDuration(qty, args[2].Str)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("add(): %v", err)
			}
			t := time.UnixMilli(int64(ts)).Add(dur)
			return environment.NewNumber(float64(t.UnixMilli())), nil
		}),

		// diff(ts1, ts2, unit) — calculates the difference between two timestamps
		fn("diff", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("diff() expects 3 arguments, got %d", len(args))
			}
			ts1, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("diff() first argument must be a number (timestamp)")
			}
			ts2, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("diff() second argument must be a number (timestamp)")
			}
			if args[2].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("diff() third argument must be a string (unit)")
			}
			t1 := time.UnixMilli(int64(ts1))
			t2 := time.UnixMilli(int64(ts2))
			d := t2.Sub(t1)

			var result float64
			switch args[2].Str {
			case "ms", "millisecond", "milliseconds":
				result = float64(d.Milliseconds())
			case "s", "second", "seconds":
				result = d.Seconds()
			case "min", "minute", "minutes":
				result = d.Minutes()
			case "h", "hour", "hours":
				result = d.Hours()
			case "day", "days":
				result = d.Hours() / 24
			case "week", "weeks":
				result = d.Hours() / (24 * 7)
			default:
				return environment.NewNil(), fmt.Errorf("diff() unknown unit: %s", args[2].Str)
			}
			return environment.NewNumber(result), nil
		}),

		// from_timestamp(timestamp) — creates a date object with year, month, day, hour, minute, second
		fn("from_timestamp", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("from_timestamp() expects 1 argument, got %d", len(args))
			}
			ts, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("from_timestamp() argument must be a number (timestamp)")
			}
			t := time.UnixMilli(int64(ts))

			entries := map[string]environment.Value{
				"year":   environment.NewNumber(float64(t.Year())),
				"month":  environment.NewNumber(float64(t.Month())),
				"day":    environment.NewNumber(float64(t.Day())),
				"hour":   environment.NewNumber(float64(t.Hour())),
				"minute": environment.NewNumber(float64(t.Minute())),
				"second": environment.NewNumber(float64(t.Second())),
			}
			keys := []string{"year", "month", "day", "hour", "minute", "second"}
			return environment.NewObject(entries, keys), nil
		}),
	))
}
