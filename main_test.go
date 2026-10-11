package main

import (
	"testing"
)

type TestNormalizeEvent struct {
	input string
	want  string
}

func TestNormalizeEventType(t *testing.T) {
	testCases := []TestNormalizeEvent{
		{input: "push", want: "PushEvent"},
		{input: "PUSH", want: "PushEvent"},
		{input: "create", want: "CreateEvent"},
		{input: "CREATE", want: "CreateEvent"},
		{input: "Invalid", want: "Invalid"},
	}
	for _, test := range testCases {
		result := normalizeEventType(test.input)
		if result != test.want {
			t.Errorf("normalizeEventType(%q): %q; want %q", test.input, result, test.want)
		}
	}
}

// func TestNormalizeEventTypeCreate(t *testing.T) {
// 	got := normalizeEventType("create")
// 	want := "CreateEvent"

// 	if got != want {
// 		t.Errorf("normalizeEventType(create) = %q; want %q", got, want)
// 	}
// }
// func TestNormalizeEventTypePUSH(t *testing.T) {
// 	got := normalizeEventType("PUSH")
// 	want := "PushEvent"

// 	if got != want {
// 		t.Errorf("normalizeEventType(PUSH) = %q; want %q", got, want)
// 	}
// }

// func TestNormalizeEventTypeInvalid(t *testing.T) {
// 	got := normalizeEventType("Invalid")
// 	want := "Invalid"

// 	if got != want {
// 		t.Errorf("normalizeEventType(Invalid) = %q; want %q", got, want)
// 	}
// }
