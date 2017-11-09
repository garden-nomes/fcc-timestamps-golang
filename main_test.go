package main

import (
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	tables := []struct {
		input       string
		expected    time.Time
		shouldError bool
	}{
		{"123456789", time.Unix(123456789, 0), false},
		{"January 2, 2006", time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC), false},
		{"", time.Time{}, true},
	}

	for _, table := range tables {
		output, err := convert(table.input)
		if output != table.expected {
			t.Errorf("convert was incorrect with input \"%v\", got: \"%v\", want: \"%v\"",
				table.input, output, table.expected)
		}

		if err == nil && table.shouldError {
			t.Errorf("expected convert to fail for input \"%v\", got: \"%v\"",
				table.input, output)
		}

		if err != nil && !table.shouldError {
			t.Errorf("convert returned error for input \"%v\": \"%v\"",
				table.input, err)
		}
	}
}

func TestOutput(t *testing.T) {
	e := emptyResponse()
	o := output{nil, nil}
	if e != o {
		t.Errorf("emptyResponse failed, got: \"%v\", want: \"%v\"", e, o)
	}

	d := time.Now()
	o = makeResponse(d)
	if *o.Unix != d.Unix() || *o.Date != d.Format(dateFmt) {
		t.Errorf("emptyResponse failed for date \"%v\", got: \"%v\"", d, o)
	}
}
