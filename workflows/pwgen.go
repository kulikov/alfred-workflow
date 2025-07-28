package workflows

import (
	"math/rand"
	"time"
)

func Pwgen(length int) []Item {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	output := make([]Item, 10)

	for i := 0; i < 10; i++ {
		password := make([]byte, length)
		for j := 0; j < length; j++ {
			password[j] = chars[rand.Intn(len(chars))]
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
