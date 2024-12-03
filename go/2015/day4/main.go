package main

import (
	"aoc/util"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	secretKey, _ := util.ReadOneLineFile()
	var number int
	var hash string

	for {
		// Concatenate the secret key and the number
		input := secretKey + strconv.Itoa(number)

		// Compute the MD5 hash
		md5Sum := md5.Sum([]byte(input))

		// Convert the hash to a hexadecimal string
		hash = hex.EncodeToString(md5Sum[:])

		// Check if the hash starts with five zeroes
		if hash[:6] == "000000" {
			break
		}

		// Increment the number
		number++
	}

	fmt.Printf("The lowest number that produces a hash starting with five zeroes is: %d\n", number)
	fmt.Printf("The resulting hash is: %s\n", hash)
}

// func main() {
// 	// input, _ := util.ReadOneLineFile()
// 	fmt.Printf("%x", md5.Sum([]byte("abcdef"+"609043")))
// }

// // md5Encode will take a string and encode it as md5.
// func md5Encode(input string) string {
// 	// Create a new hash & write input string
// 	hash := md5.New()
// 	_, _ = hash.Write([]byte(input))

// 	// Get the resulting encoded byte slice
// 	md5 := hash.Sum(nil)

// 	// Convert the encoded byte slice to a string
// 	return fmt.Sprintf("%x", md5)
// }
