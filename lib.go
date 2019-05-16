package main

import (
	"errors"
	"strconv"
)

// ==============================================================
// Input Sanitation - dumb input checking, look for empty strings
// ==============================================================
func sanitize_arguments(strs []string) error {
	for i, val := range strs {
		if len(val) <= 0 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be a non-empty string")
		}
		if len(val) > 200 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be <= 200 characters")
		}
	}
	return nil
}
