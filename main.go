package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func calculateHash(timestamp int64, data string, prevHash string, nonce int) string {
	
	record := strconv.FormatInt(timestamp, 10) + data + prevHash + strconv.Itoa(nonce)

	h := sha256.New()

	h.Write([]byte(record))

	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func main() {
	fmt.Println("--- Blockchain Hash Demonstration ---")

	timestamp := int64(1678886400)
	data := "Send 10 BTC to Bob"
	prevHash := "0000a1b2c3d4e5f60000a1b2c3d4e5f60000a1b2c3d4e5f60000a1b2c3d4e5f6"
	nonce := 12345

	hash1 := calculateHash(timestamp, data, prevHash, nonce)

	fmt.Println("\n--- Original Block ---")
	fmt.Printf("Data:       %s\n", data)
	fmt.Printf("Nonce:      %d\n", nonce)
	fmt.Printf("Hash:       %s\n", hash1)

	data2 := "Send 10 BTC to B0b"

	hash2 := calculateHash(timestamp, data2, prevHash, nonce)

	fmt.Println("\n--- Tampered Block (Tiny Change) ---")
	fmt.Printf("Data:       %s\n", data2)
	fmt.Printf("Nonce:      %d\n", nonce)
	fmt.Printf("Hash:       %s\n", hash2)

	fmt.Println("\n--- Conclusion ---")
	fmt.Printf("Hashes are identical: %t\n", hash1 == hash2)
	fmt.Println("Notice how a tiny 1-character change in the data")
	fmt.Println("results in a completely different hash.")
	fmt.Println("This is what makes the blockchain secure and tamper-evident.")
}