package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	"crypto/sha256"
)

type Block struct {
	nonce int
	previousHash string
	timestamp int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp 		 %d\n", b.timestamp)
	fmt.Printf("nonce  			 %d\n", b.nonce)
	fmt.Printf("previousHash  	 %s\n", b.previousHash)
	fmt.Printf("transactions  	 %s\n", b.transactions)
}

func init() {
	log.SetPrefix("Blockchain: ")
}

type Blockchain struct {
	transactionPool []string
	chain []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}


func main() {
	//log.Println("test")
	//fmt.Println("test2")
	//b := NewBlock(0, "init hash")
	//b.Print()
	blockchain := NewBlockchain()
	blockchain.Print()
	blockchain.CreateBlock(5, "hash 1")
	blockchain.CreateBlock(10, "hash 2")
	blockchain.Print()
	//fmt.Println(blockchain)
}
