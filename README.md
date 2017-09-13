# Ternary

Three-valued logic based on Kleene's strong logic of indeterminacy.

[![Build Status](https://travis-ci.org/mithrandie/ternary.svg?branch=master)](https://travis-ci.org/mithrandie/ternary)
[![GoDoc](https://godoc.org/github.com/mithrandie/ternary?status.svg)](http://godoc.org/github.com/mithrandie/ternary)
[![License: MIT](https://img.shields.io/badge/License-MIT-lightgrey.svg)](https://opensource.org/licenses/MIT)

## Truth values

- FALSE (-1)
- UNKNOWN (0)
- TRUE (1)


## Truth tables

```
  NOT(A) - Negative value of A
  +---+----|
  | A | ¬A |
  +---+----|
  | F |  T |
  | U |  U |
  | T |  F |
  +---+----|

  AND(A, B) - Minimum value of (A, B)
  +-------+-----------|
  |       |     B     |
  +       +---+---+---|
  |       | F | U | T |
  +---+---+---+---+---|
  |   | F | F | F | F |
  | A | U | F | U | U |
  |   | T | F | U | T |
  +---+---+---+---+---|

  OR(A, B) - Maximum value of (A, B)
  +-------+-----------|
  |       |     B     |
  +       +---+---+---|
  |       | F | U | T |
  +---+---+---+---+---|
  |   | F | F | U | T |
  | A | U | U | U | T |
  |   | T | T | T | T |
  +---+---+---+---+---|

  IMP(A, B) - NOT(A) OR B
  +-------+-----------|
  |       |     B     |
  +       +---+---+---|
  |       | F | U | T |
  +---+---+---+---+---|
  |   | F | T | U | F |
  | A | U | T | U | U |
  |   | T | T | T | T |
  +---+---+---+---+---|
```