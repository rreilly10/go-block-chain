package main

// Block struct used to encapsulate a single
// unit of the blockchain
type block struct {
	index        int           // index of this block in reference to its peers
	timestamp    int64         // creation time
	transactions []transaction // transactions interacting with other blocks
	proof        int           //
	previousHash string        // previous struct hash used for ident
}
