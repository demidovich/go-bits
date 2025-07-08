package base62

import (
	"errors"
	"math"
	"slices"
)

const (
	base       int    = 62
	characters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

var (
	decodeMap map[byte]int
)

func init() {
	decodeMap = make(map[byte]int, len(characters))
	for i, char62 := range characters {
		decodeMap[byte(char62)] = i
	}
}

func EncodeInt(val int) string {
	if val == 0 {
		return ""
	}

	if val < 0 {
		val = -val
	}

	result := make([]byte, 0, 12)
	for val > 0 {
		charnum := val % base
		result = slices.Insert(result, 0, characters[charnum])
		val /= base
	}

	return string(result)
}

func DecodeInt(val string) (int, error) {
	if val == "" {
		return 0, nil
	}

	var result int64
	for i, char := range []byte(val) {
		decimal, ok := decodeMap[char]
		if !ok {
			return 0, errors.New("invalid base62 character: " + string(char))
		}

		power := len(val) - i - 1
		result += int64(decimal) * int64(math.Pow(float64(base), float64(power)))
	}

	return int(result), nil
}
