package data

import (
	"testing"
	"time"
)

func TestDetailsStringer(t *testing.T) {
	d := LoginDetails{
		UserName:    "Test",
		PassHash:    "hash",
		Salt:        "12345",
		DateCreated: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
	}

	expected := `{"UserName":"Test","PassHash":"hash","Salt":"12345","DateCreated":"2023-01-01T12:00:00Z"}`
	result := d.String()

	if result != expected {
		t.Errorf("stringer not produing correct result, expected %s got %s", expected, result)
	}
}
