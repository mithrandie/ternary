package ternary

import (
	"testing"
)

func TestValue_String(t *testing.T) {
	s := FALSE.String()
	if s != "FALSE" {
		t.Errorf("string = %q, want %q for %s.String()", s, "FALSE", FALSE)
	}

	s = UNKNOWN.String()
	if s != "UNKNOWN" {
		t.Errorf("string = %q, want %q for %s.String()", s, "UNKNOWN", UNKNOWN)
	}

	s = TRUE.String()
	if s != "TRUE" {
		t.Errorf("string = %q, want %q for %s.String()", s, "TRUE", TRUE)
	}
}

func TestValue_Int(t *testing.T) {
	i := FALSE.Int()
	if i != -1 {
		t.Errorf("int value = %d, want %d for %s", i, -1, FALSE)
	}

	i = UNKNOWN.Int()
	if i != 0 {
		t.Errorf("int value = %d, want %d for %s", i, 0, UNKNOWN)
	}

	i = TRUE.Int()
	if i != 1 {
		t.Errorf("int value = %d, want %d for %s", i, 1, TRUE)
	}
}

func TestValue_ParseBool(t *testing.T) {
	b := FALSE.ParseBool()
	if b != false {
		t.Errorf("bool value = %t, want %t for %s", b, false, FALSE)
	}

	b = UNKNOWN.ParseBool()
	if b != false {
		t.Errorf("bool value = %t, want %t for %s", b, false, UNKNOWN)
	}

	b = TRUE.ParseBool()
	if b != true {
		t.Errorf("bool value = %t, want %t for %s", b, true, TRUE)
	}
}

var convertFromStringTests = []struct {
	Str    string
	Result Value
	Err    string
}{
	{
		Str:    "false",
		Result: FALSE,
	},
	{
		Str:    "unknown",
		Result: UNKNOWN,
	},
	{
		Str:    "true",
		Result: TRUE,
	},
	{
		Str:    "-1",
		Result: FALSE,
	},
	{
		Str:    "0",
		Result: UNKNOWN,
	},
	{
		Str:    "1",
		Result: TRUE,
	},
	{
		Str: "ParseError",
		Err: "convert from \"ParseError\": invalid value",
	},
}

func TestConvertFromString(t *testing.T) {
	for _, test := range convertFromStringTests {
		v, err := ConvertFromString(test.Str)
		if err != nil {
			if len(test.Err) < 1 {
				t.Errorf("unexpected error: %q", err.Error())
			} else if err.Error() != test.Err {
				t.Errorf("error = %q, want error %q for %s", err.Error(), test.Err, test.Str)
			}
			continue
		}
		if 0 < len(test.Err) {
			t.Errorf("no error, want error %q for %s", test.Err, test.Str)
			continue
		}
		if v != test.Result {
			t.Errorf("ternary = %s, want %s for %q", v, test.Result, test.Str)
		}
	}
}

var convertFromInt64Tests = []struct {
	Int    int64
	Result Value
	Err    string
}{
	{
		Int:    -1,
		Result: FALSE,
	},
	{
		Int:    0,
		Result: UNKNOWN,
	},
	{
		Int:    1,
		Result: TRUE,
	},
	{
		Int: 12345,
		Err: "convert from 12345: invalid value",
	},
}

func TestConvertFromInt64(t *testing.T) {
	for _, test := range convertFromInt64Tests {
		v, err := ConvertFromInt64(test.Int)
		if err != nil {
			if len(test.Err) < 1 {
				t.Errorf("unexpected error: %q", err.Error())
			} else if err.Error() != test.Err {
				t.Errorf("error = %q, want error %q for %d", err.Error(), test.Err, test.Int)
			}
			continue
		}
		if 0 < len(test.Err) {
			t.Errorf("no error, want error %q for %d", test.Err, test.Int)
			continue
		}
		if v != test.Result {
			t.Errorf("ternary = %s, want %s for %d", v, test.Result, test.Int)
		}
	}
}

func TestConvertFromBool(t *testing.T) {
	r := ConvertFromBool(false)
	if r != FALSE {
		t.Errorf("ternary = %s, want %s for %t", r, FALSE, false)
	}

	r = ConvertFromBool(true)
	if r != TRUE {
		t.Errorf("ternary = %s, want %s for %t", r, TRUE, true)
	}
}

var equivalentTests = []struct {
	Value1 Value
	Value2 Value
	Result Value
}{
	{
		Value1: FALSE,
		Value2: FALSE,
		Result: TRUE,
	},
	{
		Value1: FALSE,
		Value2: UNKNOWN,
		Result: FALSE,
	},
}

func TestEquivalent(t *testing.T) {
	for _, test := range equivalentTests {
		v := Equivalent(test.Value1, test.Value2)
		if v != test.Result {
			t.Errorf("ternary = %s, want %s for \"equal(%s, %s)\"", v, test.Result, test.Value1, test.Value2)
		}
	}
}

var notTests = []struct {
	Value  Value
	Result Value
}{
	{
		Value:  FALSE,
		Result: TRUE,
	},
	{
		Value:  TRUE,
		Result: FALSE,
	},
	{
		Value:  UNKNOWN,
		Result: UNKNOWN,
	},
}

func TestNot(t *testing.T) {
	for _, test := range notTests {
		v := Not(test.Value)
		if v != test.Result {
			t.Errorf("ternary = %s, want %s for \"not %s\"", v, test.Result, test.Value)
		}
	}
}

var andTests = []struct {
	Value1 Value
	Value2 Value
	Result Value
}{
	{
		Value1: FALSE,
		Value2: FALSE,
		Result: FALSE,
	},
	{
		Value1: FALSE,
		Value2: UNKNOWN,
		Result: FALSE,
	},
	{
		Value1: FALSE,
		Value2: TRUE,
		Result: FALSE,
	},
	{
		Value1: UNKNOWN,
		Value2: FALSE,
		Result: FALSE,
	},
	{
		Value1: UNKNOWN,
		Value2: UNKNOWN,
		Result: UNKNOWN,
	},
	{
		Value1: UNKNOWN,
		Value2: TRUE,
		Result: UNKNOWN,
	},
	{
		Value1: TRUE,
		Value2: FALSE,
		Result: FALSE,
	},
	{
		Value1: TRUE,
		Value2: UNKNOWN,
		Result: UNKNOWN,
	},
	{
		Value1: TRUE,
		Value2: TRUE,
		Result: TRUE,
	},
}

func TestAnd(t *testing.T) {
	for _, test := range andTests {
		v := And(test.Value1, test.Value2)
		if v != test.Result {
			t.Errorf("ternary = %s, want %s for \"%s and %s\"", v, test.Result, test.Value1, test.Value2)
		}
	}
}

var orTests = []struct {
	Value1 Value
	Value2 Value
	Result Value
}{
	{
		Value1: FALSE,
		Value2: FALSE,
		Result: FALSE,
	},
	{
		Value1: FALSE,
		Value2: UNKNOWN,
		Result: UNKNOWN,
	},
	{
		Value1: FALSE,
		Value2: TRUE,
		Result: TRUE,
	},
	{
		Value1: UNKNOWN,
		Value2: FALSE,
		Result: UNKNOWN,
	},
	{
		Value1: UNKNOWN,
		Value2: UNKNOWN,
		Result: UNKNOWN,
	},
	{
		Value1: UNKNOWN,
		Value2: TRUE,
		Result: TRUE,
	},
	{
		Value1: TRUE,
		Value2: FALSE,
		Result: TRUE,
	},
	{
		Value1: TRUE,
		Value2: UNKNOWN,
		Result: TRUE,
	},
	{
		Value1: TRUE,
		Value2: TRUE,
		Result: TRUE,
	},
}

func TestOr(t *testing.T) {
	for _, test := range orTests {
		v := Or(test.Value1, test.Value2)
		if v != test.Result {
			t.Errorf("ternary = %s, want %s for \"%s or %s\"", v, test.Result, test.Value1, test.Value2)
		}
	}
}

var impTests = []struct {
	Value1 Value
	Value2 Value
	Result Value
}{
	{
		Value1: TRUE,
		Value2: FALSE,
		Result: FALSE,
	},
	{
		Value1: UNKNOWN,
		Value2: UNKNOWN,
		Result: UNKNOWN,
	},
}

func TestImp(t *testing.T) {
	for _, test := range impTests {
		v := Imp(test.Value1, test.Value2)
		if v != test.Result {
			t.Errorf("ternary = %s, want %s for \"%s or %s\"", v, test.Result, test.Value1, test.Value2)
		}
	}
}

var allTests = []struct {
	ValueList []Value
	Result    Value
}{
	{
		ValueList: []Value{TRUE, TRUE, TRUE},
		Result:    TRUE,
	},
	{
		ValueList: []Value{TRUE, UNKNOWN, TRUE},
		Result:    UNKNOWN,
	},
	{
		ValueList: []Value{TRUE, UNKNOWN, FALSE},
		Result:    FALSE,
	},
	{
		ValueList: []Value{},
		Result:    TRUE,
	},
}

func TestAll(t *testing.T) {
	for _, test := range allTests {
		v := All(test.ValueList)
		if v != test.Result {
			t.Errorf("ternary = %s, want %s for all \"%s\"", v, test.Result, test.ValueList)
		}
	}
}

var anyTests = []struct {
	ValueList []Value
	Result    Value
}{
	{
		ValueList: []Value{TRUE, UNKNOWN, FALSE},
		Result:    TRUE,
	},
	{
		ValueList: []Value{FALSE, UNKNOWN, FALSE},
		Result:    UNKNOWN,
	},
	{
		ValueList: []Value{FALSE, FALSE, FALSE},
		Result:    FALSE,
	},
	{
		ValueList: []Value{},
		Result:    FALSE,
	},
}

func TestAny(t *testing.T) {
	for _, test := range anyTests {
		v := Any(test.ValueList)
		if v != test.Result {
			t.Errorf("ternary = %s, want %s for any \"%s\"", v, test.Result, test.ValueList)
		}
	}
}
