package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/cnf/structhash"
	"github.com/kr/pretty"
	"log"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// Creates a SHA-256 hash of a Block
// :param block: <dict> Block
// :return: <str>
func hash(b block) string {
	hash := structhash.Md5(b, 1)
	fmt.Println(b, hex.EncodeToString(hash))
	return hex.EncodeToString(hash)
}

// Validates the Proof: Does hash(last_proof, proof) contain 4 leading zeroes?
// :param last_proof: <int> Previous Proof
// :param proof: <int> Current Proof
// :return: <bool> True if correct, False if not.
func validProof(lastProof int, proof int) bool {
	guess := string(lastProof) + string(proof)
	guessHash := sha256.Sum256([]byte(guess))
	h := hex.EncodeToString(guessHash[:])
	return h[len(h)-4:] == "0000"
}

func mine(bc *blockchain) response {
	lastBlock := bc.lastBlock()
	lastProof := lastBlock.proof
	proof := proofOfWork(lastProof)

	bc.newTransaction("0", "node", 10)

	previousHash := hash(lastBlock)
	block := bc.newBlock(proof, previousHash)

	return response{
		message:      "New Block Forged",
		index:        block.index,
		transactions: block.transactions,
		proof:        block.proof,
		previousHash: block.previousHash,
	}

}

func main() {
	bc := blockchain{
		chain:               []block{},
		currentTransactions: []transaction{},
	}

	bc.newBlock(1, "100") // 0 block
	mine(&bc)
	bc.newTransaction("0", "new", 1)
	mine(&bc)
	// fmt.Println(len(bc.chain))
	pretty.Print(bc)

}
