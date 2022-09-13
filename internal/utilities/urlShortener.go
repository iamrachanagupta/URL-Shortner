package utilities

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
)

//generateHash256Code returns a hash code for any given input string
func generateHash256Code(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

//encodeString returns encoded string of the input byte array
//It uses base64 encoding
func encodeString(bytes []byte) string {
	encoded := base64.StdEncoding.EncodeToString(bytes)
	return encoded
}

//GenerateKey generates a unique key corresponding to the requested url string.
//
//It returns the first 8 letters of the encoding
//8 letters string can generate ~200 trillion(62**8) unique values which is sufficient for application.
func GenerateKey(reqUrl string) string {
	reqUrlHash := generateHash256Code(reqUrl)
	randNumber := new(big.Int).SetBytes(reqUrlHash).Uint64()
	finalString := encodeString([]byte(fmt.Sprintf("%d", randNumber)))
	return finalString[:8]
}
