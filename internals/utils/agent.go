package utils

import "strings"

func CheckAgent(agent string) bool {
	if strings.Contains(agent, "Chrome") {
		return true
	} else {
		return false
	}
}
