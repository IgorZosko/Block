package blockchain

import "math/big"

const Difficult = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficult))

	pow := &ProofOfWork{b, target}

	return pow
}
