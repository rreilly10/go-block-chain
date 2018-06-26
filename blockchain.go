package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/cnf/structhash"
)

// Maintains a chain of blocks, as well as
// the list of current transactions
type blockchain struct {
	chain               []block
	currentTransactions []transaction
}

// Create a new Block in the Blockchain
// :param proof: <int> The proof given by the Proof of Work algorithm
// :param previous_hash: (Optional) <str> Hash of previous Block
// :return: <block> New Block
func (bc *blockchain) newBlock(proof int, previousHash string) block {
	b := block{
		index:        len(bc.chain) + 1,
		timestamp:    time.Now().UnixNano() / int64(time.Millisecond),
		transactions: bc.currentTransactions,
		proof:        proof,
		previousHash: previousHash,
	}

	bc.chain = append(bc.chain, b)
	bc.currentTransactions = []transaction{}
	return b
}

// Creates a new transaction to go into the next mined Block
// :param sender: <str> Address of the Sender
// :param recipient: <str> Address of the Recipient
// :param amount: <int> Amount
// :return: <int> The index of the Block that will hold this transaction
func (bc *blockchain) newTransaction(sender string, recipient string, amount int) int {
	bc.currentTransactions = append(
		bc.currentTransactions,
		transaction{
			sender:    sender,
			recipient: recipient,
			amount:    amount,
		},
	)
	return bc.lastBlock().index + 1
}

//Returns the last block in the current chain
func (bc *blockchain) lastBlock() block {
	if len(bc.chain) != 0 {
		return bc.chain[len(bc.chain)-1]
	}
	return block{}
}

// Simple Proof of Work Algorithm:
// - Find a number p' such that hash(pp') contains leading 4 zeroes, where p is the previous p'
// - p is the previous proof, and p' is the new proof
// :param last_proof: <int>
//:return: <int>
func proofOfWork(lastProof int) int {
	proof := 0

	for validProof(lastProof, proof) == false {
		proof++
	}
	return proof
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

// Performaes work in order to forge a new block
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
