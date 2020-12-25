package main

import "log"

func main() {

	cardPublicKey := 18356117
	//cardPublicKey := 5764801

	doorPublicKey := 5909654
	//doorPublicKey := 17807724

	cardLoopSize := 0
	initialNumber := 7
	for val := 1; val != cardPublicKey; cardLoopSize++ {
		val = val * initialNumber % 20201227
	}

	encryptionKey := 1
	for i := 0; i < cardLoopSize; i++ {
		encryptionKey = encryptionKey * doorPublicKey % 20201227
	}

	log.Printf("Encryption Key: %v", encryptionKey)
}
