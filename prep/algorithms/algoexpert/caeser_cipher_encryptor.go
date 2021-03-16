package main

// Time; O(n) ; Space: O()
func CaesarCipherEncryptor(str string, key int) string {
	encrypted := []byte("")

	for _, letter := range str {
		newLetter := ((int(letter)-int('a'))+key)%26 + int('a')
		encrypted = append(encrypted, byte(newLetter))
	}

	return string(encrypted)
}
