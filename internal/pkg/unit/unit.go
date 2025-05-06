package unit

import (
	"fmt"
	"regexp"
)

// IDが`^[a-zA-Z0-9]+$`の形式であることを確認する
func IsValidID(id string) (string, error) {
	// 正規表現を使用してIDの形式を確認
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !re.MatchString(id) {
		return "", fmt.Errorf("invalid ID format")
	}

	// IDの長さを確認
	if len(id) > 10 {
		return "", fmt.Errorf("ID length must be 10 characters or less")
	}

	return id, nil
}
