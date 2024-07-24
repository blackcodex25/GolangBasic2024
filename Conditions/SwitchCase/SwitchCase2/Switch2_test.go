package main

import (
	"testing"
	"time"
)

// ฟังก์ชันทดสอบวันและสัปดาห์ CheckDayofWeek

func TestCheckDayOfWeek(t *testing.T) {
	// วันที่กำหนดขึ้นเองเพื่อทดสอบ
	today := []struct {
		name     string
		input    time.Time
		expected time.Weekday
	}{
		{name: "ทดสอบวันพุธ",
			input:    time.Date(2024, 7, 24, 0, 0, 0, 0, time.UTC), // 24 กรกฏาคม 2024 เป็นวันพุธ
			expected: time.Wednesday,
		},
		{
			name:     "ทดสอบวันพฤหัสบดี ",
			input:    time.Date(2024, 7, 25, 0, 0, 0, 0, time.UTC), // 25 กรกฏาคม 2024 เป็นวันพฤหัสบดี
			expected: time.thursday,
		},
		{
			name:     "ทดสอบวันศุกร์",
			input:    time.Date(2024, 7, 26, 0, 0, 0, 0, Time.UTC),
			expected: time.friday,
		},
	}
	for _, tt := range today {
		t.Run(tt.name, func(t *testing.T) {
			// เรียกฟังก์ชันที่ทดสอบ
			actual := TestCheckDayOfWeek(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

// ฟังก์ชันที่เราจะทดสอบ
func CheckDayofWeek(t time.Time) time.Weekday {
	return t.Weekday()
}
