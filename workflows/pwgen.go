package workflows

import (
	"math/rand"
	"time"
)

func Pwgen(length int) []Item {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsLower := "abcdefghijklmnopqrstuvwxyz0123456789"
	charsUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	output := make([]Item, 12)

	for i := 0; i < 12; i++ {
		l := length
		c := chars

		if i < 2 {
			// standard long
		} else if i < 4 {
			c = charsLower
		} else if i < 6 {
			c = charsUpper
		} else if i < 8 {
			l = length / 2
		} else {
			l = length / 2
			c = charsLower
		}

		if l < 1 {
			l = 1
		}

		password := make([]byte, l)
		for j := 0; j < l; j++ {
			password[j] = c[rand.Intn(len(c))]
		}

		passwordStr := string(password)

		output[i] = Item{
			Title: passwordStr,
			Arg:   passwordStr,
			Icon:  Icon{"default", "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/FileVaultIcon.icns"},
		}
	}

	return output
}
