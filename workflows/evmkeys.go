package workflows

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func Evmkeys() []Item {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	return []Item{
		{
			Title: address,
			Arg:   fmt.Sprintf("address: %s\nprivate: %x", address, crypto.FromECDSA(privateKey)),
			Icon:  Icon{"default", "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/FileVaultIcon.icns"},
		},
	}
}
