package gst

import (
	"fmt"
	"regexp"
	"strings"
)

func validatePattern(gstin string) bool {
	re := regexp.MustCompile("^([0][1-9]|[1-2][0-9]|[3][0-5])([a-zA-Z]{5}[0-9]{4}[a-zA-Z]{1}[1-9a-zA-Z]{1}[zZ]{1}[0-9a-zA-Z]{1})+$")
	return re.MatchString(gstin) // Regex validation result GSTIN of 15 digits.
}

func isValidGSTNumber(gstin string) bool {
	gstin = strings.ToUpper(gstin)
	if len(gstin) != 15 {
		return false
	}
	re := validatePattern(gstin)
	if !re {
		return false
	}
	return true
}
