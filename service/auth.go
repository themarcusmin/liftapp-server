// Package service contains common functions used by
// the whole application
package service

import (
	"encoding/hex"

	"github.com/pilinux/crypt"
	"golang.org/x/crypto/blake2b"

	"liftapp/config"
	"liftapp/database"
	"liftapp/model"
)

// GetUserByEmail fetches auth info by email or hash of the email
func GetUserByEmail(email string, decryptEmail bool) (*model.Auth, error) {
	db := database.GetDB()
	var err error

	var auth model.Auth

	// when email is saved in plaintext
	if err = db.Where("email = ? ", email).First(&auth).Error; err == nil {
		return &auth, nil
	}

	// encryption at rest
	if config.IsCipher() {
		// hash of the email in hexadecimal string format
		emailHash, err := CalcHash(
			[]byte(email),
			config.GetConfig().Security.Blake2bSec,
		)
		if err != nil {
			return nil, err
		}

		// email must be unique
		if err = db.Where("email_hash = ?", hex.EncodeToString(emailHash)).First(&auth).Error; err == nil {
			if decryptEmail {
				auth.Email, err = DecryptEmail(auth.EmailNonce, auth.EmailCipher)
				if err != nil {
					return nil, err
				}
			}

			return &auth, nil
		}
	}

	return nil, err
}

// CalcHash generates a fixed-sized BLAKE2b-256 hash of the given text
func CalcHash(plaintext, keyOptional []byte) ([]byte, error) {
	blake2b256Hash, err := blake2b.New256(keyOptional)
	if err != nil {
		return nil, err
	}

	_, err = blake2b256Hash.Write(plaintext)
	if err != nil {
		return nil, err
	}

	blake2b256Sum := blake2b256Hash.Sum(nil)

	return blake2b256Sum, nil
}

// DecryptEmail returns the plaintext email from the given cipher and nonce
func DecryptEmail(emailNonce, emailCipher string) (email string, err error) {
	nonce, err := hex.DecodeString(emailNonce)
	if err != nil {
		return
	}
	cipherEmail, err := hex.DecodeString(emailCipher)
	if err != nil {
		return
	}

	email, err = crypt.DecryptChacha20poly1305(
		config.GetConfig().Security.CipherKey,
		nonce,
		cipherEmail,
	)
	return
}
