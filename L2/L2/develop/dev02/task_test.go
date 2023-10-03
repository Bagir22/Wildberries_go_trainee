package main

import "testing"

type testData struct {
	input    string
	expected string
}

func TestRLEDecode_Case1(t *testing.T) {
	input := "a4bc2d5e"
	expected := "aaaabccddddde"
	v, _ := RLEDecode(input)
	if expected != v {
		t.Errorf("expected %q, got %q", expected, v)
	}
}

func TestRLEDecode_Case2(t *testing.T) {
	input := "abcd"
	expected := "abcd"
	v, _ := RLEDecode(input)
	if expected != v {
		t.Errorf("expected %q, got %q", expected, v)
	}
}

func TestRLEDecode_Case3(t *testing.T) {
	input := "45"
	expected := ""
	v, _ := RLEDecode(input)
	if expected != v {
		t.Errorf("expected %q, got %q", expected, v)
	}
}

func TestRLEDecode_Case4(t *testing.T) {
	input := ""
	expected := ""
	v, _ := RLEDecode(input)
	if expected != v {
		t.Errorf("expected %q, got %q", expected, v)
	}
}

func TestRLEDecode_Case5(t *testing.T) {
	input := "qwe\\4\\5"
	expected := "qwe45"
	v, _ := RLEDecode(input)
	if expected != v {
		t.Errorf("expected %q, got %q", expected, v)
	}
}

func TestRLEDecode_Case6(t *testing.T) {
	input := "qwe\\45"
	expected := "qwe44444"
	v, _ := RLEDecode(input)
	if expected != v {
		t.Errorf("expected %q, got %q", expected, v)
	}
}

func TestRLEDecode_Case7(t *testing.T) {
	input := "qwe\\\\5"
	expected := "qwe\\\\\\\\\\"
	v, _ := RLEDecode(input)
	if expected != v {
		t.Errorf("expected %q, got %q", expected, v)
	}
}






