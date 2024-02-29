package utils

import "strings"

func Isvalid(cep string) (bool, error) {
	var result strings.Builder

	for _, char := range cep {
		if char != '-' && char != '_' {
			result.WriteRune(char)
		}
	}

	if len(result.String()) > 8 {
		return false, nil
	}

	return true, nil
}

func RemoveFirstAndLast(s []byte) string {
	configs := string(s)

	return configs[1 : len(configs)-1]
}
