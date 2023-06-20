package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"strings"
)

func SignAddress(signatureKey string) (string, error) {

	signatureKey = strings.TrimPrefix(signatureKey, "0x")

	signatureBytes, err := hex.DecodeString(signatureKey)
	if err != nil {
		return "", err
	}

	x := new(big.Int).SetBytes(signatureBytes[:32])
	y := new(big.Int).SetBytes(signatureBytes[32:])

	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	publicKeyBytes := elliptic.Marshal(publicKey.Curve, publicKey.X, publicKey.Y)
	addressBytes := sha256.Sum256(publicKeyBytes)
	address := hex.EncodeToString(addressBytes[:])
	return address, nil
}
