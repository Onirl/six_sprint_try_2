package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertFileData(baseStr string) string {
	s := strings.ReplaceAll(baseStr, ".", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, " ", "")
	if len(s) == 0 {
		return morse.ToText(baseStr)
	} else {
		return morse.ToMorse(baseStr)
	}
}
