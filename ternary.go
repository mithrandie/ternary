// Package ternary is a Go library to calculate three-valued logic.
//
// This package is based on Kleene's strong logic of indeterminacy.
// Ternary has three truth values, TRUE, FALSE and UNKNOWN.
//
// Numeric representation of truth values
/*
  FALSE:   -1
  UNKNOWN:  0
  TRUE:     1
*/
//
// Truth tables
/*
  NOT(A) - Logical negation
  +---+----+
  | A | ¬A |
  |---+----|
  | F |  T |
  | U |  U |
  | T |  F |
  +---+----+

  AND(A, B) - Logical conjunction. Minimum value of (A, B)
  +--------+-----------|
  |        |     B     |
  | A ∧ B  |---+---+---|
  |        | F | U | T |
  |----+---+---+---+---|
  |    | F | F | F | F |
  | A  | U | F | U | U |
  |    | T | F | U | T |
  +----+---+---+---+---+

  OR(A, B) - Logical disjunction. Maximum value of (A, B)
  +--------+-----------+
  |        |     B     |
  | A ∨ B  |---+---+---|
  |        | F | U | T |
  |----+---+---+---+---|
  |    | F | F | U | T |
  | A  | U | U | U | T |
  |    | T | T | T | T |
  +----+---+---+---+---+

  IMP(A, B) - Logical implication. OR(NOT(A), B)
  +--------+-----------+
  |        |     B     |
  | A → B  |---+---+---|
  |        | F | U | T |
  |----+---+---+---+---|
  |    | F | T | T | T |
  | A  | U | U | U | T |
  |    | T | F | U | T |
  +----+---+---+---+---+

  EQV(A, B) - Logical biconditional. OR(AND(A, B), AND(NOT(A), NOT(B)))
  +--------+-----------+
  |        |     B     |
  | A ↔ B  |---+---+---|
  |        | F | U | T |
  |----+---+---+---+---|
  |    | F | T | U | F |
  | A  | U | U | U | U |
  |    | T | F | U | T |
  +----+---+---+---+---+

*/
package ternary

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Truth values
type Value int8

const (
	FALSE Value = iota - 1
	UNKNOWN
	TRUE
)

var literals = map[Value]string{
	FALSE:   "FALSE",
	UNKNOWN: "UNKNOWN",
	TRUE:    "TRUE",
}

// Returns string representation of the value.
func (value Value) String() string {
	return literals[value]
}

// Returns integer representation of the value.
func (value Value) Int() int64 {
	return reflect.ValueOf(value).Int()
}

// Returns true if the value is TRUE, otherwise returns false.
func (value Value) ParseBool() bool {
	if value != TRUE {
		return false
	}
	return true
}

// Converts s to a ternary value.
// If s is any of "false", "FALSE" and "-1", then it is converted to FALSE.
// If s is any of "unknown", "UNKNOWN" and "0", then it is converted to UNKNOWN.
// If s is any of "true", "TRUE" and "1", then it is converted to TRUE.
// Otherwise returns an error.
func ConvertFromString(s string) (Value, error) {
	switch strings.ToUpper(s) {
	case "FALSE", "-1":
		return FALSE, nil
	case "TRUE", "1":
		return TRUE, nil
	case "UNKNOWN", "0":
		return UNKNOWN, nil
	}
	return UNKNOWN, errors.New(fmt.Sprintf("convert from %q: invalid value", s))
}

// Converts i to a ternary value.
// Returns FALSE if i is -1, returns UNKNOWN if i is 0, returns TRUE if i is 1.
// Otherwise returns an error.
func ConvertFromInt64(i int64) (Value, error) {
	switch i {
	case -1:
		return FALSE, nil
	case 0:
		return UNKNOWN, nil
	case 1:
		return TRUE, nil
	}
	return UNKNOWN, errors.New(fmt.Sprintf("convert from %d: invalid value", i))
}

// Converts b to a ternary value.
// Returns FALSE if i is false, returns TRUE if i is true.
func ConvertFromBool(b bool) Value {
	if b {
		return TRUE
	}
	return FALSE
}

// Check if two values are the same value, not logical equality.
func Equal(a Value, b Value) Value {
	if a == b {
		return TRUE
	}
	return FALSE
}

// Returns the result of logical negation on a.
func Not(a Value) Value {
	switch a {
	case FALSE:
		return TRUE
	case TRUE:
		return FALSE
	}
	return UNKNOWN
}

// Returns the result of logical conjunction on two values.
func And(a Value, b Value) Value {
	switch {
	case a == FALSE || b == FALSE:
		return FALSE
	case a == UNKNOWN || b == UNKNOWN:
		return UNKNOWN
	}
	return TRUE
}

// Returns the result of logical disjunction on two values.
func Or(a Value, b Value) Value {
	switch {
	case a == TRUE || b == TRUE:
		return TRUE
	case a == UNKNOWN || b == UNKNOWN:
		return UNKNOWN
	}
	return FALSE
}

// Returns the result of logical implication that is represented as "a implies b".
func Imp(a Value, b Value) Value {
	return Or(Not(a), b)
}

// Returns the result of logical biconditional on two values.
func Eqv(a Value, b Value) Value {
	if a == UNKNOWN || b == UNKNOWN {
		return UNKNOWN
	}
	return ConvertFromBool(a == b)
}

// Returns the result of logical conjunction on all values.
func All(values []Value) Value {
	t := TRUE
	if 0 < len(values) {
		t = values[0]
	}
	for i := 1; i < len(values); i++ {
		t = And(t, values[i])
		if t == FALSE {
			return FALSE
		}
	}
	return t
}

// Returns the result of logical disjunction on all values.
func Any(values []Value) Value {
	t := FALSE
	if 0 < len(values) {
		t = values[0]
	}
	for i := 1; i < len(values); i++ {
		t = Or(t, values[i])
		if t == TRUE {
			return TRUE
		}
	}
	return t
}
