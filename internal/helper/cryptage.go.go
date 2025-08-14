package helper

import (
	"bytes"
	"car_project/internal/config"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func Encrypt(data string) (encodedmess string, err error) {
	key := []byte(config.PPE)

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	plaintext, _ := pkcs7Pad([]byte(data), block.BlockSize())
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	bm := cipher.NewCBCEncrypter(block, iv)
	bm.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return fmt.Sprintf("%x", ciphertext), nil

}

func pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, errors.New("invalid blocksize")
	}
	if len(b) == 0 {
		return nil, errors.New("invalid PKCS7 data (empty or not padded)")
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

func Decrypt(encodedmess string) (decodedmess string, err error) {
	key := []byte(config.PPE)

	// Convert the encoded message back to bytes
	ciphertext, err := hex.DecodeString(encodedmess)
	if err != nil {
		return "", err
	}

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV is the first block
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Decrypt the message
	bm := cipher.NewCBCDecrypter(block, iv)
	bm.CryptBlocks(ciphertext, ciphertext)

	// Remove padding
	plaintext, err := pkcs7Unpad(ciphertext, block.BlockSize())
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func pkcs7Unpad(b []byte, blocksize int) ([]byte, error) {
	if len(b) == 0 {
		return nil, errors.New("invalid PKCS7 data (empty or not padded)")
	}
	if len(b)%blocksize != 0 {
		return nil, errors.New("invalid PKCS7 padding (blocksize mismatch)")
	}
	padLen := int(b[len(b)-1])
	if padLen > blocksize || padLen == 0 {
		return nil, errors.New("invalid PKCS7 padding (invalid padding length)")
	}
	for _, v := range b[len(b)-padLen:] {
		if int(v) != padLen {
			return nil, errors.New("invalid PKCS7 padding (invalid padding bytes)")
		}
	}
	return b[:len(b)-padLen], nil
}
