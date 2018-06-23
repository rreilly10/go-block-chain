package main

type response struct {
	message      string
	index        int
	transactions []transaction
	proof        int
	previousHash string
}
