package tests

import (
	"strings"
	"testing"
	"time"
)

func useDate(code string) string {
	return "use \"date\"\n" + code
}

func TestDateNow(t *testing.T) {
	out, err := runFig(t, useDate(`print(date.now())`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if strings.Contains(out, "e+") || strings.Contains(out, "E+") {
		t.Fatalf("date.now() returned scientific notation: %s", out)
	}
	if len(out) < 13 {
		t.Fatalf("date.now() returned unexpected value: %s", out)
	}
}

func TestDateNowNoSciNotation(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts = date.now()
print(ts)
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if strings.Contains(out, "e+") {
		t.Fatalf("timestamp should not be in scientific notation: %s", out)
	}
}

func TestDateFormat(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts = date.parse("2024-06-15 10:30:00", "YYYY-MM-DD HH:mm:ss")
print(date.format(ts, "YYYY-MM-DD"))
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "2024-06-15" {
		t.Fatalf("expected '2024-06-15', got '%s'", out)
	}
}

func TestDateFormatTime(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts = date.parse("2024-06-15 10:30:45", "YYYY-MM-DD HH:mm:ss")
print(date.format(ts, "HH:mm:ss"))
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "10:30:45" {
		t.Fatalf("expected '10:30:45', got '%s'", out)
	}
}

func TestDateParse(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts = date.parse("2024-01-01", "YYYY-MM-DD")
print(ts)
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if len(out) < 13 {
		t.Fatalf("expected a 13-digit timestamp, got '%s'", out)
	}
	if strings.Contains(out, "e+") {
		t.Fatalf("timestamp should not be in scientific notation: %s", out)
	}
}

func TestDateParseRoundtrip(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts = date.parse("2024-12-25 08:00:00", "YYYY-MM-DD HH:mm:ss")
print(date.format(ts, "YYYY-MM-DD HH:mm:ss"))
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "2024-12-25 08:00:00" {
		t.Fatalf("expected '2024-12-25 08:00:00', got '%s'", out)
	}
}

func TestDateAddDay(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts = date.parse("2024-01-01", "YYYY-MM-DD")
let amanha = date.add(ts, 1, "day")
print(date.format(amanha, "YYYY-MM-DD"))
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "2024-01-02" {
		t.Fatalf("expected '2024-01-02', got '%s'", out)
	}
}

func TestDateAddHour(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts = date.parse("2024-01-01 10:00:00", "YYYY-MM-DD HH:mm:ss")
let after = date.add(ts, 3, "hour")
print(date.format(after, "HH:mm:ss"))
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "13:00:00" {
		t.Fatalf("expected '13:00:00', got '%s'", out)
	}
}

func TestDateAddMinute(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts = date.parse("2024-01-01 10:00:00", "YYYY-MM-DD HH:mm:ss")
let after = date.add(ts, 45, "minute")
print(date.format(after, "HH:mm:ss"))
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "10:45:00" {
		t.Fatalf("expected '10:45:00', got '%s'", out)
	}
}

func TestDateAddWeek(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts = date.parse("2024-01-01", "YYYY-MM-DD")
let after = date.add(ts, 2, "week")
print(date.format(after, "YYYY-MM-DD"))
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "2024-01-15" {
		t.Fatalf("expected '2024-01-15', got '%s'", out)
	}
}

func TestDateAddSeconds(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts = date.parse("2024-01-01 10:00:00", "YYYY-MM-DD HH:mm:ss")
let after = date.add(ts, 30, "second")
print(date.format(after, "HH:mm:ss"))
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "10:00:30" {
		t.Fatalf("expected '10:00:30', got '%s'", out)
	}
}

func TestDateDiffDays(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts1 = date.parse("2024-01-01", "YYYY-MM-DD")
let ts2 = date.parse("2024-01-11", "YYYY-MM-DD")
print(date.diff(ts1, ts2, "day"))
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "10" {
		t.Fatalf("expected '10', got '%s'", out)
	}
}

func TestDateDiffHours(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts1 = date.parse("2024-01-01 00:00:00", "YYYY-MM-DD HH:mm:ss")
let ts2 = date.parse("2024-01-01 05:30:00", "YYYY-MM-DD HH:mm:ss")
print(date.diff(ts1, ts2, "hour"))
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "5.5" {
		t.Fatalf("expected '5.5', got '%s'", out)
	}
}

func TestDateDiffNegative(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts1 = date.parse("2024-01-11", "YYYY-MM-DD")
let ts2 = date.parse("2024-01-01", "YYYY-MM-DD")
print(date.diff(ts1, ts2, "day"))
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "-10" {
		t.Fatalf("expected '-10', got '%s'", out)
	}
}

func TestDateDiffMilliseconds(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts1 = date.parse("2024-01-01 10:00:00", "YYYY-MM-DD HH:mm:ss")
let ts2 = date.parse("2024-01-01 10:00:01", "YYYY-MM-DD HH:mm:ss")
print(date.diff(ts1, ts2, "ms"))
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "1000" {
		t.Fatalf("expected '1000', got '%s'", out)
	}
}

func TestDateFromTimestamp(t *testing.T) {
	out, err := runFig(t, useDate(`
let ts = date.parse("2024-06-15 10:30:45", "YYYY-MM-DD HH:mm:ss")
let dt = date.from_timestamp(ts)
print(dt.year)
print(dt.month)
print(dt.day)
print(dt.hour)
print(dt.minute)
print(dt.second)
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	expected := []string{"2024", "6", "15", "10", "30", "45"}
	if len(lines) != 6 {
		t.Fatalf("expected 6 lines, got %d: %q", len(lines), out)
	}
	for i, exp := range expected {
		if strings.TrimSpace(lines[i]) != exp {
			t.Fatalf("line %d: expected '%s', got '%s'", i, exp, strings.TrimSpace(lines[i]))
		}
	}
}

func TestDateFormatNowLive(t *testing.T) {
	out, err := runFig(t, useDate(`
let data = date.format(date.now(), "YYYY-MM-DD")
print(data)
`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := time.Now().Format("2006-01-02")
	if out != expected {
		t.Fatalf("expected '%s', got '%s'", expected, out)
	}
}

func TestDateErrorHandling(t *testing.T) {
	_, err := runFig(t, useDate(`date.format("abc", "YYYY")`))
	if err == nil {
		t.Fatal("expected error for non-number timestamp")
	}

	_, err = runFig(t, useDate(`date.parse("not-a-date", "YYYY-MM-DD")`))
	if err == nil {
		t.Fatal("expected error for invalid date string")
	}

	_, err = runFig(t, useDate(`date.add(date.now(), 1, "fortnight")`))
	if err == nil {
		t.Fatal("expected error for unknown time unit")
	}

	_, err = runFig(t, useDate(`date.diff(date.now(), date.now(), "fortnight")`))
	if err == nil {
		t.Fatal("expected error for unknown time unit")
	}
}

func TestSystemNowNoSciNotation(t *testing.T) {
	out, err := runFig(t, "use \"system\"\nprint(system.now())")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if strings.Contains(out, "e+") {
		t.Fatalf("system.now() should not use scientific notation: %s", out)
	}
}
