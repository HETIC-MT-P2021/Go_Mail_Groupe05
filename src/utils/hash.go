package utils

import (
	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// HashPassword will hash a given string into a unreadable byte array
func HashPassword(password string, saltString string) []byte {
	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		keyLength:   32,
	}

	salt := []byte(saltString)

	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hashedPass := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return hashedPass
}

// CheckPass will check if a user's saved password and a given hashed password are the same
func CheckPass(enteredPassword string, savedPassword []byte, saltString string) bool {
	hashedEnteredPassword := HashPassword(enteredPassword, saltString)

	var isTheSame = true

	if len(hashedEnteredPassword) != len(savedPassword) {
		isTheSame = false
	}

	for i := 0; i < len(savedPassword); i++ {
		if hashedEnteredPassword[i] != savedPassword[i] {
			isTheSame = false
		}
	}

	return isTheSame
}
