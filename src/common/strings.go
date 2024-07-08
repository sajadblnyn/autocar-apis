package common

import (
	"math"
	"math/rand"
	"strconv"
	"time"
	"unicode"

	"github.com/sajadblnyn/autocar-apis/config"
)

func CheckPassword(password string) bool {
	cfg := config.GetConfig()
	if len(password) < cfg.Password.MinLength {
		return false
	}

	if cfg.Password.IncludeChars && !HasLetter(password) {
		return false
	}

	if cfg.Password.IncludeDigits && !HasDigits(password) {
		return false
	}

	if cfg.Password.IncludeLowercase && !HasLower(password) {
		return false
	}

	if cfg.Password.IncludeUppercase && !HasUpper(password) {
		return false
	}

	return true
}

func HasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
func HasLower(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasDigits(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func GenerateOtp() string {
	cfg := config.GetConfig()
	rand.Seed(time.Now().UnixNano())
	min := int(math.Pow(10, float64(cfg.Otp.Digits-1)))
	max := int(math.Pow(10, float64(cfg.Otp.Digits))) - 1

	var num = rand.Intn(max-min) + min

	return strconv.Itoa(num)
}
